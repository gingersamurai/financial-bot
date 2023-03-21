package messages

import (
	"fmt"
	"strings"
	"time"

	"github.com/gingersamurai/financial-bot/internal/model/storage"
)

func parseGetReportMessage(rawData string, dateStart *time.Time) error {
	rawdataSlice := strings.Split(rawData, " ")
	if len(rawdataSlice) == 2 {
		duration := rawdataSlice[1]
		switch duration {
		case "day":
			*dateStart = time.Now().Add(-24 * time.Hour)
		case "week":
			*dateStart = time.Now().Add(-24 * 31 * time.Hour)
		case "year":
			*dateStart = time.Now().Add(-24 * 365 * time.Hour)
		default:
			return fmt.Errorf("wrong duration")
		}
		return nil
	} else {
		return fmt.Errorf("need 1 argument, got :%d", len(rawdataSlice)-1)
	}
}

func getSum(data map[string][]storage.Spending) (map[string]int, error) {
	result := make(map[string]int)
	for k, v := range data {
		for _, elem := range v {
			result[k] += elem.Count
		}
	}
	return result, nil
}

func getReport(msg Message, dest storage.Storage) (string, error) {
	var dateStart time.Time
	err := parseGetReportMessage(msg.Text, &dateStart)
	if err != nil {
		return "", err
	}

	var data map[string][]storage.Spending
	data, err = dest.Find(dateStart, time.Now())
	if err != nil {
		return "", err
	}

	result := "отчет:\n"
	sums, err := getSum(data)
	if err != nil {
		return "", err
	}

	for k, v := range sums {
		result += fmt.Sprintf("    группа: %v    общие расходы: %v\n", k, v)
	}
	return result, nil

}
