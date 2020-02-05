//Package gigasecond calculates when someone would be a gigasecond old. Includes testing.
package gigasecond

// import path for the time package from the standard library
import "time"

//a gigasecond in nanoseconds
const gigasecond = "1000000000s"

// AddGigasecond Adds a Gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	gs, _ := time.ParseDuration(gigasecond)
	t2 := t.Add(gs)
	return t2
}
