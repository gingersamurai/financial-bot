package messages

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gingersamurai/financial-bot/internal/model/storage"
)

func parseDateFromMessage(dest *time.Time, src string) error {
	layout := "01/02/2006" // month/day/year
	var err error

	*dest, err = time.Parse(layout, src)

	return err

}

func parseAddSpendingMessage(rawData string, dest *storage.Spending) error {
	rawDataSlice := strings.Split(rawData, " ")
	var err error
	if len(rawDataSlice) == 3 || len(rawDataSlice) == 4 {

		// validate count
		dest.Count, err = strconv.Atoi(string(rawDataSlice[1]))
		if err != nil {
			return err
		}

		// validate group
		dest.Group = string(rawDataSlice[2])

		// validate case with date
		if len(rawDataSlice) == 4 {
			err = parseDateFromMessage(&dest.Date, rawDataSlice[3])
			if err != nil {
				return err
			}
		} else {
			dest.Date = time.Now()
		}

	} else {
		return fmt.Errorf("need 3 or 2 arguments, got %d", len(rawDataSlice))
	}
	return nil
}

func addSpending(msg Message, dest storage.Storage) error {
	var spending storage.Spending
	err := parseAddSpendingMessage(msg.Text, &spending)
	if err != nil {
		return err
	}
	return dest.Insert(spending)
}
