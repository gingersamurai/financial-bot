package main

import (
	"fmt"
	"time"
)

type Spending struct {
	count int
	group string
	date  time.Time
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

func (m MemoryStorage) Find(startDate time.Time, finishDate time.Time) (map[string][]Spending, error) {
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

func main() {
	s := NewMemoryStorage()
	s.Insert(Spending{
		group: "abc",
		count: 1203,
		date:  time.Now(),
	})
	startDate, _ := time.Parse("01/02/2006", "03/19/2023")
	finishDate, _ := time.Parse("01/02/2006", "03/21/2024")
	fmt.Println(s.storage)
	res, err := s.Find(startDate, finishDate)
	if err != nil {
		fmt.Println(res)
	} else {
		fmt.Println()
	}
}
