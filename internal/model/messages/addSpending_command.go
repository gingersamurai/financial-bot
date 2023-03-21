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

type Storage interface {
	Insert(spending Spending) error
	Find(startDate time.Time, finishDate time.Time) (map[string][]Spending, error)
}

type MemoryStorage struct {
	storage map[string][]Spending
}

func NewMemoryStorage() MemoryStorage {
	result := MemoryStorage{}
	result.storage = make(map[string][]Spending)
	return result
}

func (m *MemoryStorage) Insert(spending Spending) error {
	m.storage[spending.group] = append(m.storage[spending.group], spending)
	return nil
}

func (m *MemoryStorage) Find(startDate time.Time, finishDate time.Time) (map[string][]Spending, error) {
	result := make(map[string][]Spending)
	for k, v := range m.storage {
		for _, elem := range v {
			if elem.date.After(startDate) && elem.date.Before(finishDate) {
				result[k] = append(result[k], elem)
			}
		}
	}
	return result, nil
}

var myStorage MemoryStorage

func parseDateFromMessage(dest *time.Time, src string) error {
	layout := "01/02/2006" // month/day/year
	var err error

	*dest, err = time.Parse(layout, src)

	return err

}

func parseAddSpendingMessage(rawData string, dest *Spending) error {
	rawDataSlice := strings.Split(rawData, " ")
	var err error
	if len(rawDataSlice) == 3 || len(rawDataSlice) == 4 {

		// validate count
		dest.count, err = strconv.Atoi(string(rawDataSlice[1]))
		if err != nil {
			return err
		}

		// validate group
		dest.group = string(rawDataSlice[2])

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

func addSpending(msg Message, dest Storage) error {
	var spending Spending
	err := parseAddSpendingMessage(msg.Text, &spending)
	if err != nil {
		return err
	}
	dest.Insert(spending)
	return nil
}
