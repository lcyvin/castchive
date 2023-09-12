package podcast

import (
  "encoding/json"
  "github.com/lcyvin/castchive/pkg/v1/episode"
)

type NewPodcastOption func(*Podcast)

func WithNameOverride(name string) NewPodcastOption {
  return func(p *Podcast) {
    p.Name = name
  }
}

func WithRSSUrl(url string) NewPodcastOption {
  return func(p *Podcast) {
    p.RSSUrl = url
  }
}

func WithName(name string) NewPodcastOption {
  return func(p *Podcast) {
    p.Name = name
  }
}

func WithDescription(desc string) NewPodcastOption {
  return func(p *Podcast) {
    p.Description = desc
  }
}

type Podcast struct {
  Name          string  `json:"name"`
  NameOverride  string  `json:"nameOverride"`
  Description   string  `json:"description"`
  Provider      string  `json:"provider"`
  EpisodeCount  int     `json:"episodeCount"`
  RSSUrl        string  `json:"rss"`
}

func NewFromUrl(url string, opts...NewPodcastOption) (*Podcast, error) {
  p := &Podcast{
    RSSUrl: url,
  }

  for _, opt := range opts {
    opt(p)
  }

  return p, nil
}

func New(opts...NewPodcastOption) (*Podcast, error) {
  p := &Podcast{}

  for _, opt := range opts {
    opt(p)
  }

  return p, nil
}

func (p *Podcast) Marshal() ([]byte, error) {
  out,err := json.Marshal(p)
  if err != nil {
    return nil, err
  }

  return out, nil
}

func Unmarshal(doc []byte) (*Podcast, error) {
  p := &Podcast{}

  err := json.Unmarshal(doc, p)
  if err != nil {
    return nil, err
  }

  return p, nil
}
