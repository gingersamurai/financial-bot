package messages

import (
	"testing"
	"time"

	"github.com/gingersamurai/financial-bot/internal/model/storage"
	"github.com/stretchr/testify/assert"
)

func TestParseDateFromMessage(t *testing.T) {
	t.Run("simpleTest", func(t *testing.T) {
		var result time.Time
		err := parseDateFromMessage(&result, "09/15/2003")

		assert.NoError(t, err)

		assert.Equal(t, int(result.Month()), 9)
		assert.Equal(t, int(result.Day()), 15)
		assert.Equal(t, int(result.Year()), 2003)
	})

	t.Run("wrongFormat", func(t *testing.T) {
		var result time.Time
		err := parseDateFromMessage(&result, "13/15/2003")

		assert.Error(t, err)
	})
}

func TestParseAddSpendingMessage(t *testing.T) {
	t.Run("simpleTest", func(t *testing.T) {
		var result storage.Spending
		err := parseAddSpendingMessage(
			"/addSpending 500 testGroup 09/15/2003",
			&result,
		)

		assert.NoError(t, err)

		parsedDate, _ := time.Parse("01/02/2006", "09/15/2003")
		assert.Equal(
			t,
			result,
			storage.Spending{
				Count: 500,
				Group: "testGroup",
				Date:  parsedDate,
			},
		)
	})

	t.Run("tooMuchArgs", func(t *testing.T) {
		var result storage.Spending
		err := parseAddSpendingMessage(
			"/addSpending 100 rub pivGroup 01/02/2003",
			&result,
		)

		assert.Error(t, err)
	})

	t.Run("notEnoughArgs", func(t *testing.T) {
		var result storage.Spending
		err := parseAddSpendingMessage(
			"/addSpending 0",
			&result,
		)

		assert.Error(t, err)
	})
}
