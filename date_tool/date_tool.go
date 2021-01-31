package date_tool

import "time"

var E8 = time.FixedZone("GMT+8", +8*60*60)

//current time in EAST 8 time zone
func NowInE8() time.Time {
    return time.Now().In(E8)
}

//format current time in EAST 8 time zone to yyyy-MM-dd HH:mm:ss
func NowDateTime() string {
    return DateTime(time.Now().In(E8))
}

//format current time in EAST 8 time zone to yyyy-MM-dd
func NowDate() string {
    return Date(time.Now().In(E8))
}

//format current time in EAST 8 time zone to HH:mm:ss
func NowTime() string {
    return Time(time.Now().In(E8))
}

//format time to yyyy-MM-dd HH:mm:ss
func DateTime(time time.Time) string {
    return time.Format("2006-01-02 15:04:05")
}

//format time to yyyy-MM-dd
func Date(time time.Time) string {
    return time.Format("2006-01-02")
}

//format time to HH:mm:ss
func Time(time time.Time) string {
    return time.Format("15:04:05")
}

//format time to yyyy-MM-dd HH:mm:ss, return "date",formatTime
func NowDateTimeWithLabel() (string, string) {
    return "date", DateTime(time.Now().In(E8))
}

//parse time from yyyy-MM-dd HH:mm:ss format time string, eg: "2006-01-02 15:04:05"
func ParseDateTime(timeStr string) time.Time {
    result, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, E8)
    if nil != err {
        panic(err)
    }
    return result
}

//parse time from yyyy-MM-dd format time string, eg: "2006-01-02"
func ParseDate(str string) time.Time {
    result, err := time.ParseInLocation("2006-01-02", str, E8)
    if nil != err {
        panic(err)
    }
    return result
}

//parse time from HH:mm:ss format time string, eg: "15:04:05"
func ParseTime(str string) time.Time {
    result, err := time.ParseInLocation("15:04:05", str, E8)

    if nil != err {
        panic(err)
    }
    return result
}
