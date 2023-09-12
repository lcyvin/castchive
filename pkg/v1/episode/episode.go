package episode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
  "time"
)

// Episode describes the data and information contained within
// a given podcast episode. It is used to structure data for the database
// as well as hold configuration information for fetcher/streamer/writer
type Episode struct {
  // Media source URL for the podcasts's file
  Source          string        `json:"source"`
  // Media format (mp3, m4a, etc.)
  Format          string        `json:"format"`
  // Episode title (can be set automatically from XML or overwritten)
  // manually or with custom naming rules
  Title           string        `json:"title"`
  // Publishing date for the particular episode
  PublishDate     time.Time     `json:"publishDate"`
  // Autonumbering of the episode's idx within the larger collection
  // of episodes in the database/collection
  EpisodeNumber   int           `json:"episodeNumber"`
  // Alternate mode of addressing the specific episode (eg 12A, 125.5, etc.)
  // this field is not used for sorting operations, it is decorative and only
  // used in querying
  Episode         string        `json:"episode"`
  // XML metadata related to the episode
  Metadata        xml.CharData  `json:"metadata"`
  // States whether or not the podcast file is in a video format
  Video           bool          `json:"video"`
  // Location of the stored file if it is stored locally
  StoragePath     string        `json:"storagePath"`
  // The "channel" this episode belongs to, eg podcast title
  PodcastChannel  string        `json:"podcastChannel"`
  // Tells fetcher/streamer that this episode should not be downloaded
  // or streamed automatically, only logged in the catalogue of episodes
  NoFetch         bool          `json:"noFetch"`
  // internal use, tells fetcher/streamer/etc. to stream file data
  // to a bytebuffer for use with stdout or other data-streaming options
  OutputStream    bool
}

func (ep *Episode) Marshal() ([]byte, error) {
  out, err := json.Marshal(ep)
  if err != nil {
    return nil, err
  }

  return out, nil
}

func (ep *Episode) StringMetadata() (string, error) {
  if ep.Metadata == nil {
    return "", fmt.Errorf("No metadata present")
  }

  return string(ep.Metadata), nil
}

type NewEpisodeOption func(*Episode)

func WithStream() (NewEpisodeOption) {
  return func(ep *Episode) {
    ep.OutputStream = true
  }
}

func WithNoFetch() (NewEpisodeOption) {
  return func(ep *Episode) {
    ep.NoFetch = true
  }
}

func WithPublishDate(d time.Time) NewEpisodeOption {
  return func(ep *Episode) {
    ep.PublishDate = d
  }
}

func WithTitleOverride(title string) (NewEpisodeOption) {
  return func(ep *Episode) {
    ep.Title = title
  }
}

func WithEpisodeNumberOverride(epnum int) (NewEpisodeOption) {
  return func(ep *Episode) {
    ep.EpisodeNumber = epnum
  }
}

func WithMetadata(raw xml.CharData) (NewEpisodeOption) {
  return func(ep *Episode) {
    ep.Metadata = raw
  }
}

func WithPodcastChannel(c string) NewEpisodeOption {
  return func(ep *Episode) {
    ep.PodcastChannel = c
  }
}

func WithSource(s string) NewEpisodeOption {
  return func(ep *Episode) {
    ep.Source = s
  }
}

func New(title string, storagePath string, opts...NewEpisodeOption) (*Episode, error) {
  ep := &Episode{
    StoragePath: storagePath,
    Title: title,
    NoFetch: false,
    OutputStream: false,
  }

  for _,opt := range opts {
    opt(ep)
  }

  return ep, nil
} 
