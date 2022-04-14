package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

const URL string = "http://api.aladhan.com/v1/timingsByAddress"

type prayer struct {
	azan      string
	timezone  string
	timestamp string
}

var (
	ErrRequestFailed = errors.New("response code is not 200")
)

func getPrayerTime(address, iso8601, ts string, method int) (prayer, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return prayer{}, err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("address", address)
	q.Add("date_or_timestamp", ts)
	q.Add("method", strconv.Itoa(method))
	q.Add("iso8601", iso8601)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return prayer{}, err
	}

	if resp.StatusCode != 200 {
		return prayer{}, ErrRequestFailed
	}

	var result Result

	if err = loadJSON(resp, &result); err != nil {
		return prayer{}, err
	}

	// x, _ := json.MarshalIndent(&result, "", "   ")
	// fmt.Printf("%s\n", x)

	return prayer{
		azan:      result.Data.Timings.Maghrib,
		timezone:  result.Data.Meta.Timezone,
		timestamp: result.Data.Date.Timestamp}, nil
}

func loadJSON[T any](resp *http.Response, strct *T) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &strct); err != nil {
		return err
	}

	return nil
}

func getNow(tz string) (time.Time, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.Time{}, err
	}

	now := time.Now().In(loc)

	return now, nil
}

func delta(now, then time.Time) time.Duration {
	diff := then.Sub(now)
	return diff

}

func Reminder(city, timestamp string) (time.Duration, error) {
	prayer, err := getPrayerTime(city, "true", timestamp, 7)
	if err != nil {
		return time.Nanosecond, err
	}

	now, err := getNow(prayer.timezone)
	if err != nil {
		return time.Nanosecond, err
	}
	azan, err := time.Parse(time.RFC3339, prayer.azan)
	if err != nil {
		return time.Nanosecond, err
	}
	rem := delta(now, azan)

	return rem, nil
}
