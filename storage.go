package main

import (
	"fmt"
	"sort"
)

type Storage struct {
	lastID   int
	allPlans map[int]Plan
}

func NewStorage() *Storage {
	return &Storage{
		allPlans: make(map[int]Plan),
	}
}

func (s *Storage) GetAllPlans() []Plan {

	var plans = make([]Plan, 0, len(s.allPlans))

	for _, p := range s.allPlans {
		plans = append(plans, p)
	}

	sort.Slice(plans, func(i, j int) bool {
		return plans[i].ID < plans[j].ID
	})

	return plans
}

func (s *Storage) CreatePlan(p Plan) int {
	fmt.Println("Створюємо новий план! Намагаємося")
	s.lastID++
	p.ID = s.lastID
	s.allPlans[p.ID] = p
	fmt.Println("Вуху, план створено! Останній id", s.lastID)
	return p.ID
}
