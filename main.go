package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	plans := PlanResource{}

	mux.HandleFunc("GET /plans", plans.GetAll)
	fmt.Println("Слухаєм :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Невдала спроба створити та прослухати 8080", err)
	}
}

type PlanResource struct {
	s Storage
}

func (p *PlanResource) GetAll(w http.ResponseWriter, r *http.Request) {
	plans := p.s.GetAll()

	err := json.NewEncoder(w).Encode(plans)
	if err != nil {
		fmt.Println("ПОмилка кодування в JSON", err)
		return
	}
}
