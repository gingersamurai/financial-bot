package storage

import (
	"fmt"
	"time"
)

type MemoryStorage struct {
	storage map[string][]Spending
}

func NewMemoryStorage() MemoryStorage {
	result := MemoryStorage{}
	result.storage = make(map[string][]Spending)
	return result
}

func (m *MemoryStorage) Insert(spending Spending) error {
	m.storage[spending.Group] = append(m.storage[spending.Group], spending)
	return nil
}

func (m *MemoryStorage) Find(startDate time.Time, finishDate time.Time) (map[string][]Spending, error) {
	result := make(map[string][]Spending)
	for k, v := range m.storage {
		for _, elem := range v {
			if elem.Date.After(startDate) && elem.Date.Before(finishDate) {
				result[k] = append(result[k], elem)
			}
		}
	}
	return result, nil
}

func (m *MemoryStorage) String() string {
	result := "memoryStorage:\n"
	for k, v := range m.storage {
		result += fmt.Sprintf("[%v]\n", k)
		for _, elem := range v {
			result += fmt.Sprintf("    %v %v\n", elem.Date.Format("01/02/2006"), elem.Count)
		}
	}
	return result
}
