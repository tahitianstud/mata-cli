package date

import (
	"github.com/araddon/dateparse"
	"time"
)

// StrToDate converts a string to a date instance
func StrToDate(dateStr string) (time.Time, error) {

	timeDate, err := dateparse.ParseAny(dateStr)

	if err != nil {
		return time.Time{}, err
	}

	return timeDate, nil

}

func DateToStr(time time.Time, format ...string) string {

	if len(format) > 0 {
		formatStr := format[0]

		return time.Format(formatStr)
	}

	return time.String()

}