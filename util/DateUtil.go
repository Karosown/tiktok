package util

import "time"

func timetoString(time time.Time) string {
	return time.String()[0:19]
}
