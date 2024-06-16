package main

import (
	"fmt"
	"sort"
	"sync"
)

type Storage struct {
	m        sync.Mutex
	lastID   int
	allPlans map[int]Plan
}

func NewStorage() *Storage {
	return &Storage{
		allPlans: make(map[int]Plan),
	}
}

func (s *Storage) GetAllPlans() []Plan {
	s.m.Lock()
	defer s.m.Unlock()
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
	s.m.Lock()
	defer s.m.Unlock()
	fmt.Println("Створюємо новий план! Намагаємося")
	s.lastID++
	p.ID = s.lastID
	s.allPlans[p.ID] = p
	fmt.Println("Вуху, план створено! Останній id", s.lastID)
	return p.ID
}
func (s *Storage) GetPlanById(id int) (Plan, bool) {
	s.m.Lock()
	defer s.m.Unlock()
	fmt.Println("Перевіряємо чи існує план з таким ID")

	p, ok := s.allPlans[id]

	return p, ok
}
func (s *Storage) DeletePlanById(id int) bool {
	s.m.Lock()
	defer s.m.Unlock()
	fmt.Println("Видаляємо план")
	_, ok := s.allPlans[id]

	if !ok {
		return false
	}

	delete(s.allPlans, id)
	return true
}
