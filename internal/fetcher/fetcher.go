package fetcher

import (
//	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/lcyvin/castchive/pkg/v1/episode"
)

type FetcherBackend string

const (
  BACKEND_YTDLP    FetcherBackend = "yt-dlp"
  BACKEND_INTERNAL FetcherBackend = "internal"
)

type FetcherOptions struct {
  UseTLS          bool            `json:"useTLS"`
  Retry           bool            `json:"retry"`
  RetryLimit      int             `json:"retryLimit"`
  RetryWait       string          `json:"retryWait"`
  Backend         FetcherBackend  `json:"backend"`
  ForceDownload   bool            `json:"forceDownload"`
  TimeoutSeconds  int64           `json:"timeoutSeconds"`
}

type FetcherOption func(*FetcherOptions)

func WithYTDLP() (FetcherOption) {
  return func(fo *FetcherOptions) {
    fo.Backend = BACKEND_YTDLP
  }
}

func NewFetcherOptions(opts...FetcherOption) (*FetcherOptions) {
  fo := &FetcherOptions{
    Backend: BACKEND_INTERNAL,
    Retry: true,
    RetryLimit: 3,
    RetryWait: "0s",
    UseTLS: true,
    ForceDownload: false,
    TimeoutSeconds: 0,
  }

  for _,opt := range opts {
    opt(fo)
  }

  return fo
}

type Fetcher struct {
  Options *FetcherOptions
}

func (f *Fetcher) httpGetFile(uri string, w io.Writer) (int64, error) {
  client := http.Client{}
  if f.Options.TimeoutSeconds > 0 {
    client.Timeout = time.Duration(f.Options.TimeoutSeconds) * time.Second
  }

  resp, err := client.Get(uri)
  if err != nil {
    return 0, err
  }
  defer resp.Body.Close()

  bytecount, err := io.Copy(w, resp.Body)
  if err != nil {
    return bytecount, err
  }

  return bytecount, nil
}

func (f *Fetcher) streamFetcher(e *episode.Episode) (int64, error) {
  return f.httpGetFile(e.Source, os.Stdout)
}

func (f *Fetcher) writeFetcher(e *episode.Episode) (int64, error) {
  fh, err := os.Create(e.StoragePath)
  if err != nil {
    return 0, err
  }

  return f.httpGetFile(e.Source, fh)
}

func (f *Fetcher) FetchEpisode(e *episode.Episode) (int64, error) {
  var bytecount int64
  if e.Source == "" {
    return 0, errors.New("Unable to fetch without source")
  }

  if e.OutputStream {
    bytecount, err := f.streamFetcher(e)
    if err != nil {
      return bytecount, err
    }
  }

  return bytecount, nil
}
