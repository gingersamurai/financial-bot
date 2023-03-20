package messages

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDateFromMessage(t *testing.T) {
	t.Run("simpleTest", func(t *testing.T) {
		var result time.Time
		err := parseDateFromMessage(&result, "09/15/2003")

		assert.Equal(t, int(result.Month()), 9)
		assert.Equal(t, int(result.Day()), 15)
		assert.Equal(t, int(result.Year()), 2003)

		assert.NoError(t, err)
	})

	t.Run("wrongFormat", func(t *testing.T) {
		var result time.Time
		err := parseDateFromMessage(&result, "13/15/2003")

		assert.Error(t, err)
	})
}
