package main

import (
  "log"

  "github.com/lcyvin/castchive/pkg/v1/episode"
  "github.com/lcyvin/castchive/internal/fetcher"
)

func main() {
  epFile := "https://chtbl.com/track/5899E/podtrac.com/pts/redirect.mp3/traffic.omny.fm/d/clips/e73c998e-6e60-432f-8610-ae210140c5b1/e5f91208-cc7e-4726-a312-ae280140ad11/2e145345-7114-4d70-9e26-b06f0020b8d4/audio.mp3?utm_source=Podcast&amp;in_playlist=d64f756d-6d5e-4fae-b24f-ae280140ad36"
  pe, err := episode.New("Behind the Bastards Test", "", episode.WithStream(), episode.WithSource(epFile))
  if err != nil {
    log.Fatal(err)
  }

  f := fetcher.Fetcher{
    Options: fetcher.NewFetcherOptions(),
  }

  _,err = f.FetchEpisode(pe)
  if err != nil {
    log.Fatal(err)
  }
}
