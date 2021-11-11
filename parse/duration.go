package parse

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(durationString string) (*time.Time, error) {
	// format is # (second, minute, hour, month, year)
	durationArray := strings.Split(durationString, " ")
	if len(durationArray) < 2 {
		return nil, fmt.Errorf("invalid duration specified")
	}

	amount, err := strconv.Atoi(durationArray[0])
	if err != nil {
		return nil, err
	}

	var expiration = time.Now()

	switch strings.ToLower(durationArray[1]) {
	case "year":
		fallthrough
	case "years":
		expiration = expiration.AddDate(amount, 0, 0)
	case "month":
		fallthrough
	case "months":
		expiration = expiration.AddDate(0, amount, 0)
	case "day":
		fallthrough
	case "days":
		expiration = expiration.AddDate(0, 0, amount)
	case "hour":
		fallthrough
	case "hours":
		expiration = expiration.Add(time.Duration(amount) * time.Hour)
	case "minute":
		fallthrough
	case "minutes":
		expiration = expiration.Add(time.Duration(amount) * time.Minute)
	case "second":
		fallthrough
	case "seconds":
		expiration = expiration.Add(time.Duration(amount) * time.Second)
	}

	return &expiration, nil
}
