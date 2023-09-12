package indexer

import (
  "github.com/lcyvin/castchive/pkg/v1/episode"
)

type FilterTarget string

func (ft FilterTarget) String() string {
  return string(ft)
}

const (
  FILTER_TARGET_TITLE FilterTarget = "title"
  FILTER_TARGET_META FilterTarget  = "metadata"
)

type IndexerFilter interface {
  Filter(episode *episode.Episode) (bool, error)
}

type Indexer struct {
  Episodes      []*episode.Episode        `json:"episodes"`
  AutoIndex     map[int]*episode.Episode  `json:"autoIndex"`
  RuledIndex    map[int]*episode.Episode  `json:"ruledIndex"`
  ExcludeRules  []*IndexerFilter          `json:"excludeRules"`
  IncludeRules  []*IndexerFilter          `json:"includeRules"`
}
