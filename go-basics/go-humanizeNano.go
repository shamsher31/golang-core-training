package main

import "fmt"
import "time"
import "strconv"

func main() {
	now := time.Now().UnixNano() / int64(time.Nanosecond)
	mili := time.Millisecond
	nano := time.Nanosecond
	unix := time.Now().UnixNano()
	fmt.Println(now, mili, nano, humanizeNano(int64(unix)))
}

func humanizeNano(n int64) string {
	var suffix string

	switch {
	case n > 1e9:
		n /= 1e9
		suffix = "s"
	case n > 1e6:
		n /= 1e6
		suffix = "ms"
	case n > 1e3:
		n /= 1e3
		suffix = "us"
	default:
		suffix = "ns"
	}

	return strconv.Itoa(int(n)) + suffix
}
