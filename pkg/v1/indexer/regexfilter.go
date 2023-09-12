package indexer

import (
  "regexp"
  "github.com/lcyvin/castchive/pkg/v1/episode"
)

type RegexMatchFilter struct {
  Regex       *regexp.Regexp  `json:"regex"`
  RegexTarget FilterTarget    `json:"regexFilterTarget"`
}

func NewRegexMatchFilter(pattern string, usePosix bool, target FilterTarget) (*RegexMatchFilter, error) {
  var re *regexp.Regexp
  var err error

  if usePosix {
    re, err = regexp.CompilePOSIX(pattern)
  } else {
    re, err = regexp.Compile(pattern)
  }

  if err != nil {
    return nil, err
  }

  return &RegexMatchFilter{Regex: re, RegexTarget: target}, nil
}

func (rmf *RegexMatchFilter) Filter(e *episode.Episode) (bool, error) {
  var testStr string = e.Title
  var err error

  if rmf.RegexTarget == FILTER_TARGET_META {
    testStr, err = e.StringMetadata()
    if err != nil {
      return false, err
    }
  }

  return rmf.Regex.MatchString(testStr), nil
}
