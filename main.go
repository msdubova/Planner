package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	plans := PlanResource{
		s: NewStorage(),
	}

	mux.HandleFunc("GET /plans", plans.GetAllPlans)
	mux.HandleFunc("POST /plans", plans.CreatePlan)

	fmt.Println("Слухаєм :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Невдала спроба створити та прослухати 8080", err)
	}
}

type PlanResource struct {
	s *Storage
}

func (p *PlanResource) GetAllPlans(w http.ResponseWriter, r *http.Request) {
	plans := p.s.GetAllPlans()

	err := json.NewEncoder(w).Encode(plans)
	if err != nil {
		fmt.Println("ПОмилка кодування в JSON", err)
		return
	}
}

func (p *PlanResource) CreatePlan(w http.ResponseWriter, r *http.Request) {
	var plan Plan

	err := json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		fmt.Println("ПОмилка декодування", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	plan.ID = p.s.CreatePlan(plan)

	err = json.NewEncoder(w).Encode(plan)
	if err != nil {
		fmt.Println("ПОмилка кодування в JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
