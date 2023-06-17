package uptime

import (
	"fmt"
	"net/http"
	"time"
)

// Uptime returns the number of seconds since the last successful HTTP GET request to the given URL.
func Uptime(url string) (int, error) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Request to %s failed: %s", url, response.Status)
	}
	return int(time.Since(start).Milliseconds()), nil
}
