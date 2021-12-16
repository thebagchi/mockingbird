package common

import (
	"time"
)

func CalculateTimeForOffset(duration time.Duration) time.Time {
	return time.Now().UTC().Add(duration)
}

func CalculateDurationFromSecond(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}

func FormatTime(value string) time.Duration {
	t, err := time.Parse("15:04:05", value)
	if nil != err {
		return time.Duration(0)
	} else {
		return time.Duration(t.Hour()*60*60 + t.Minute()*60 + t.Second())
	}
}

func FormatDate(value string) time.Time {
	t, err := time.Parse("2006-01-02", value)
	if nil == err {
		return t
	} else {
		return time.Date(01, 01, 01, 0, 0, 0, 0, time.UTC)
	}
}

func FormatDateTime(timestamp int64) string {
	return LocalTimeOf(timestamp).Format("2006-01-02 15:04:05")
}

func DefaultTimeString(timestamp int64) string {
	return LocalTimeOf(timestamp).Format("15:04:05")
}

func LocalTimeString() string {
	return LocalTimeUTC().Format("2006-01-02 15:04:05")
}

func LocalTimeUTC() time.Time {
	return time.Now().UTC()
}

func UTCTimeOffset(offset int64) time.Time {
	t := time.Now().UTC()
	return t.Add(time.Duration(offset) * time.Millisecond)
}

func UTCTimeOf(timestamp int64) time.Time {
	local := time.Now().Local()
	t := time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, time.UTC)
	return t.Add(time.Duration(timestamp) * time.Millisecond)
}

func LocalTimeOffset(offset int64) time.Time {
	t := time.Now().UTC()
	return t.Add(time.Duration(offset) * time.Millisecond)
}

func LocalTimeOf(timestamp int64) time.Time {
	local := time.Now().Local()
	t := time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, local.Location())
	return t.Add(time.Duration(timestamp) * time.Millisecond)
}

func OffsetTimeString(timestamp, offset int64) string {
	t := time.Unix(timestamp/MillisPerSecond, 0)
	t.Add(time.Duration(offset) * time.Millisecond)
	return t.Format("15:04 PM")
}

func ResetDate(t *time.Time) time.Time {
	return time.Date(1970, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
}

func EpochDate() time.Time {
	return time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
}

func InBetweenOffset(start, end, offset int64) bool {
	currentTime := LocalTimeOffset(offset)
	startTime := LocalTimeOf(start)
	endTime := LocalTimeOf(end)

	currentTime = ResetDate(&currentTime)
	startTime = ResetDate(&startTime)
	endTime = ResetDate(&endTime)

	if startTime.Before(endTime) && !startTime.Equal(endTime) {
		if (startTime.Before(currentTime) || startTime.Equal(currentTime)) &&
			(endTime.After(currentTime) || endTime.Equal(currentTime)) {
			return true
		}
	} else {
		if currentTime.Before(endTime) || currentTime.Equal(endTime) {
			currentTime = currentTime.Add(time.Duration(24) * time.Hour)
		}
		endTime = endTime.Add(time.Duration(24) * time.Hour)
		if (startTime.Before(currentTime) || startTime.Equal(currentTime)) &&
			(endTime.After(currentTime) || endTime.Equal(currentTime)) {
			return true
		}
	}
	return false
}

func InBetween(start, end, now int64) bool {
	currentTime := LocalTimeOf(now)
	startTime := LocalTimeOf(start)
	endTime := LocalTimeOf(end)
	if startTime.Before(endTime) && !startTime.Equal(endTime) {
		if (startTime.Before(currentTime) || startTime.Equal(currentTime)) &&
			(endTime.After(currentTime) || endTime.Equal(currentTime)) {
			return true
		}
	} else {
		if currentTime.Before(endTime) || currentTime.Equal(endTime) {
			currentTime = currentTime.Add(time.Duration(24) * time.Hour)
		}
		endTime = endTime.Add(time.Duration(24) * time.Hour)
		if (startTime.Before(currentTime) || startTime.Equal(currentTime)) &&
			(endTime.After(currentTime) || endTime.Equal(currentTime)) {
			return true
		}
	}
	return false
}

const (
	DaysPerCycle     = 146097
	Days0000To1970   = (DaysPerCycle * 5) - (30*365 + 7)
	HoursPerDay      = 24
	MinutesPerHour   = 60
	MinutesPerDay    = MinutesPerHour * HoursPerDay
	SecondsPerMinute = 60
	SecondsPerHour   = SecondsPerMinute * MinutesPerHour
	SecondsPerDay    = SecondsPerHour * HoursPerDay
	MillisPerSecond  = 1000
	MillisPerMinute  = MillisPerSecond * SecondsPerMinute
	MillisPerHour    = MillisPerSecond * SecondsPerHour
	MillisPerDay     = SecondsPerDay * MillisPerSecond
	MicrosPerMilli   = 1000
	MicrosPerDay     = MillisPerDay * MicrosPerMilli
	NanosPerMicro    = 1000
	NanosPerMilli    = 1000 * 1000
	NanosPerSecond   = 1000 * 1000 * 1000
	NanosPerMinute   = NanosPerSecond * SecondsPerMinute
	NanosPerHour     = NanosPerMinute * MinutesPerHour
	NanosPerDay      = NanosPerHour * HoursPerDay
	SecondsPerWeek   = 7 * SecondsPerDay
	DaysPer400Years  = 365*400 + 97
	DaysPer100Years  = 365*100 + 24
	DaysPer4Years    = 365*4 + 1
)

func NumberOfDays(date string) int64 {
	value, err := time.Parse("2006-01-02", date)
	if nil == err {
		days := (value.Unix() / SecondsPerDay) + Days0000To1970
		return days
	}
	return -1
}

func DaysOfNumber(days int64) string {
	value := days - Days0000To1970
	date := time.Unix(value*SecondsPerDay, 0)
	return date.Format("2006-01-02")
}

func TimeToDays(date time.Time) int64 {
	return (date.Unix() / SecondsPerDay) + Days0000To1970
}

func DaysToTime(days int64) time.Time {
	return time.Unix((days-Days0000To1970)*SecondsPerDay, 0)
}

func LocalDay(offset int64) int32 {
	t := LocalTimeOffset(offset)
	return int32(t.Weekday())
}

func OffsetIndia() int64 {
	return (5 * MillisPerHour) + (30 * MillisPerMinute)
}

func CurrentMilliSeconds() int64 {
	return time.Now().Unix() * 1000
}

func UnixTimestampToTime(timestamp int64) time.Time {
	var (
		seconds = timestamp / MillisPerSecond
		nanos   = (timestamp - seconds*MillisPerSecond) / (NanosPerMilli)
	)
	return time.Unix(seconds, nanos).In(time.UTC)
}

func TimeToUnixTimestamp(t time.Time) int64 {
	return t.Unix() + int64(t.Nanosecond()/NanosPerMilli)
}
