# lilac

lilac,go tools.

latest version: `v1.1.7`

require:`github.com/bittygarden/lilac v1.1.7`

import: `github.com/bittygarden/lilac/date_tool`

import: `github.com/bittygarden/lilac/io_tool`

## date_tool

* `date_tool.NowInE8() time.Time` current time in EAST 8 time zone.
* `date_tool.NowDateTime() string` return current time in EAST 8 time zone with yyyy-MM-dd HH:mm:ss format.
* `date_tool.NowDate() string` return current time in EAST 8 time zone with yyyy-MM-dd format.
* `date_tool.NowTime() string` return current time in EAST 8 time zone with HH:mm:ss format.
* `date_tool.DateTime(time time.Time) string` format time to yyyy-MM-dd HH:mm:ss.
* `date_tool.Date(time time.Time) string` format time to yyyy-MM-dd.
* `date_tool.Time(time time.Time) string` format time to HH:mm:ss.
* `date_tool.ParseDateTime(timeStr string) time.Time` parse yyyy-MM-dd HH:mm:ss format time string to time, eg: "2006-01-02 15:04:05".
* `date_tool.ParseDate(timeStr string) time.Time` parse yyyy-MM-dd format time string to time, eg: "2006-01-02".
* `date_tool.ParseTime(timeStr string) time.Time` parse HH:mm:ss format time string to time, eg: "15:04:05".

## io_tool

* `io_tool.FileNExists(filePath string) bool` return true if file exists.
* `io_tool.FileNotExists(filePath string) bool` return true if file not exists.
