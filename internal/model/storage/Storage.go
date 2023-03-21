package storage

import "time"

type Spending struct {
	Count int
	Group string
	Date  time.Time
}

type Storage interface {
	Insert(spending Spending) error
	Find(startDate time.Time, finishDate time.Time) (map[string][]Spending, error)
}
