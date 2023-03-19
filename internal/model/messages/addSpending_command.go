package messages

import "time"

type Spending struct {
	count int
	group string
	date  time.Time
}

// func parseMessage(rawData string) (Spending, error) {

// }

// func addSpending(msg Message) error {

// }
