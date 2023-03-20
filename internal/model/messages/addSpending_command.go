package messages

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
	layout := "01/02/2006" // month/day/year
	var err error

	*dest, err = time.Parse(layout, src)

	return err

}

func parseMessage(rawData string, dest *Spending) error {
	rawDataSlice := strings.Split(rawData, " ")
	var err error
	if len(rawDataSlice) == 3 || len(rawDataSlice) == 4 {

		// validate count
		dest.count, err = strconv.Atoi(string(rawData[1]))
		if err != nil {
			return err
		}

		// validate group
		dest.group = string(rawData[2])

		// validate case with date
		if len(rawDataSlice) == 4 {
			err = parseDateFromMessage(&dest.date, rawDataSlice[3])
			if err != nil {
				return err
			}
		} else {
			dest.date = time.Now()
		}

	} else {
		return fmt.Errorf("need 3 or 2 arguments, got %d", len(rawDataSlice))
	}
	return nil
}

// func addSpending(msg Message) error {
// 	Spending
// 	return parseMessage()
// }
