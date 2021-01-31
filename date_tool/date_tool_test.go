package date_tool

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

const (
    dateTimeStr string = "2021-01-31 15:15:15"
    dateStr     string = "2021-01-31"
    timeStr     string = "15:15:15"
)

var W8 = time.FixedZone("GMT-8", -8*60*60)

func TestNowInE8(t *testing.T) {
    nowInE8 := NowInE8()
    location := nowInE8.Location()
    assert.Equal(t, "GMT+8", location.String())
}

func TestNowDateTime(t *testing.T) {
    assert.NotEmpty(t, NowDateTime())
    assert.NotEmpty(t, NowDate())
    assert.NotEmpty(t, NowTime())
}

func TestDifferentLocalTime(t *testing.T) {
    now := time.Now()

    e8 := now.In(E8)
    e8Str := DateTime(e8)

    w8 := now.In(W8)
    w8str := DateTime(w8)

    assert.NotEqual(t, e8Str, w8str)
}

func TestParseDateTime(t *testing.T) {
    dateTime := ParseDateTime(dateTimeStr)
    formatDateTime := DateTime(dateTime)
    assert.Equal(t, dateTimeStr, formatDateTime)
}

func TestParseDate(t *testing.T) {
    dateTime := ParseDate(dateStr)
    formatDate := Date(dateTime)
    assert.Equal(t, dateStr, formatDate)
}

func TestParseTime(t *testing.T) {
    dateTime := ParseTime(timeStr)
    formatTime := Time(dateTime)
    assert.Equal(t, timeStr, formatTime)
}
