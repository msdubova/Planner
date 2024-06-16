package main

import "sort"

type Storage struct {
	allPlans map[int]Plan
}

func NewStorage() *Storage {
	return &Storage{
		allPlans: make(map[int]Plan),
	}
}

func (s *Storage) GetAll() []Plan {

	var plans = make([]Plan, 0, len(s.allPlans))

	for _, p := range s.allPlans {
		plans = append(plans, p)
	}

	sort.Slice(plans, func(i, j int) bool {
		return plans[i].ID < plans[j].ID
	})

	return plans
}
