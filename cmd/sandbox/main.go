package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Spending struct {
	count int
	group string
	date  time.Time
}

func parseDateFromMessage(dest *time.Time, src string) error {
	return nil
}

func parseMessage(rawData string, dest *Spending) error {
	rawDataSlice := strings.Split(rawData, " ")
	var err error
	if len(rawDataSlice) == 3 || len(rawDataSlice) == 4 {

		dest.count, err = strconv.Atoi(string(rawData[1]))
		if err != nil {
			return err
		}

		dest.group = string(rawData[2])

		if len(rawDataSlice) == 4 {
			err = parseDateFromMessage(&dest.date, rawDataSlice[3])
			if err != nil {
				return err
			}
		}

	} else {
		return fmt.Errorf("need 3 or 2 arguments, got %d", len(rawDataSlice))
	}
	return nil
}

func main() {
	err := parseMessage()
}
