package util

import "time"
import "testing"

func TestTimetoString(T *testing.T) {
	println(timetoString(time.Now()))
}
