package indexer

import (
	"errors"
	"strings"
	"time"
  "fmt"

	"github.com/lcyvin/castchive/pkg/v1/episode"
)

type StringDate string

func (s StringDate) String() (string) {
  return string(s)
}

func (s StringDate) Time() (time.Time) {
  t,_ := time.Parse(time.RFC3339, s.String())
  return t
}

type DateFilter struct {
  DateBefore  StringDate  `json:"dateBefore"`
  DateAfter   StringDate  `json:"dateAfter"`
  Day         string      `json:"day"`
}

func NewDateFilter(dateBefore, dateAfter, day string) (*DateFilter, error) {
  df := &DateFilter{
    DateBefore: StringDate(dateBefore),
    DateAfter: StringDate(dateAfter),
  }

  if !isValidDay(day) {
    return nil, errors.New(fmt.Sprintf("Not a valid day: %s", day))
  }

  df.Day = day

  return df, nil
}

func (df *DateFilter) Filter(e *episode.Episode) (bool, error) {
  var res bool = true

  if df.Day != "" {
    res = testDayMatch(df.Day, e.PublishDate) && res
  }
 
  if df.DateAfter != "" {
    res = testDateAfter(e.PublishDate, df.DateAfter.Time()) && res
  }

  if df.DateBefore != "" {
    res = testDateBefore(e.PublishDate, df.DateBefore.Time()) && res
  }  

  return res, nil
}

func testDateBefore(date, testDate time.Time) (bool) {
  res := false

  if testDate.Compare(date) <= 0 {
    res = true
  }

  return res
}

func testDateAfter(date, testDate time.Time) (bool) {
  res := false

  if testDate.Compare(date) >= 0 {
    res = true
  }

  return res
}

func testDayMatch(day string, testDate time.Time) (bool) {
  return strings.ToLower(day) == strings.ToLower(testDate.Weekday().String())
}

func isValidDay(day string) (bool) {
  var isValid bool = false
  for i := 0; i < 7; i++ {
    if isValid {
      break
    }

    wd := time.Weekday(i).String()
    if strings.ToLower(wd) == strings.ToLower(day) {
      isValid = true
    }
  }

  return isValid
}
