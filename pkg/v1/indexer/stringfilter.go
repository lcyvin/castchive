package indexer

import (
	"errors"
	"strings"

	"github.com/lcyvin/castchive/pkg/v1/episode"
)

type StringFilterKind string

const (
  STRING_FILTER_FULL            StringFilterKind = "full"
  STRING_FILTER_SUBSTR          StringFilterKind = "substr"
  STRING_FILTER_SUBSTR_NTH      StringFilterKind = "substrNth"
  STRING_FILTER_SUBSTR_IDX      StringFilterKind = "substrIdx"
  STRING_FILTER_SUBSTR_WORD_IDX StringFilterKind = "substrWordIdx"
)

type FilterOpts struct {
  SubstringIdxStart     int           `json:"substrIdxStart"`
  SubstringIdxEnd       int           `json:"substrIdxEnd"`
  SubstringWordIdx      []int         `json:"substrWordIdx"`
  SubstringNthWord      int           `json:"substrNthWord"`
  IdxExact              bool          `json:"substrExactMatch"`
  FilterTarget          FilterTarget  `json:"filterTarget"`
  substrWordIdxStart    int
  substrWordIdxEnd      int
}

type StringFilter struct{
  String        string            `json:"matchStr"`
  Kind          StringFilterKind  `json:"filterKind"`
  Opts          *FilterOpts       `json:"filterOpts"`
}

func NewStringFilter(s string, kind StringFilterKind, opts *FilterOpts) (*StringFilter, error) {
  if s == "" {
    return nil, errors.New("cannot instantiate empty string match filter")
  }

  sf := &StringFilter{String: s, Kind: kind}
  if opts != nil {
    sf.Opts = opts
  }

  return sf, nil
}

func (sf *StringFilter) Filter(e *episode.Episode) (bool, error) {
  var res bool = false
  var err error
  var testStr string 
  
  if sf.Opts != nil && sf.Opts.FilterTarget == FILTER_TARGET_META {
    testStr, err = e.StringMetadata()
    if err != nil {
      return false, err
    }
  }
  
  if sf.Opts != nil && sf.Opts.SubstringWordIdx != nil {
    if len(sf.Opts.SubstringWordIdx) > 1 {
      sf.Opts.substrWordIdxStart = sf.Opts.SubstringWordIdx[0]
      sf.Opts.substrWordIdxEnd = sf.Opts.SubstringWordIdx[1]
    }
  }

  if sf.Opts != nil {
    switch sf.Kind {
    case STRING_FILTER_SUBSTR_IDX:
      testStr = testStr[sf.Opts.SubstringIdxStart:sf.Opts.SubstringIdxEnd]
    case STRING_FILTER_SUBSTR_WORD_IDX:
      if len(sf.Opts.SubstringWordIdx) > 1 {
        testStr = strings.Join(wordSlice(testStr)[sf.Opts.substrWordIdxStart:sf.Opts.substrWordIdxEnd], " ")
      } else {
        testStr = strings.Join(wordSlice(testStr)[sf.Opts.SubstringWordIdx[0]:], " ")
      }
    }
  }

  switch sf.Kind {
  case STRING_FILTER_FULL:
    res = testStr == sf.String
  case STRING_FILTER_SUBSTR:
    res = strings.Contains(sf.String, testStr)
  case STRING_FILTER_SUBSTR_IDX:
    res = substrMatch(testStr, sf.String, sf.Opts.IdxExact)
  case STRING_FILTER_SUBSTR_NTH:
    res = wordSlice(testStr)[sf.Opts.SubstringNthWord] == sf.String
  case STRING_FILTER_SUBSTR_WORD_IDX:
    res = substrMatch(testStr, sf.String, sf.Opts.IdxExact) 
  }

  return res, nil 
}

func wordSlice(s string) ([]string) {
  return strings.Fields(s)
}

func substrMatch(test, substr string, useExact bool) (bool) {
  if useExact {
    return test == substr
  }

  return strings.Contains(test, substr)
}
