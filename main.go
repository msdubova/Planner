package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	plans := PlanResource{}
	mux.HandleFunc("GET /plans", plans.GetAll)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Невдала спроба створити та прослухати 8080", err)
	}
}

type PlanResource struct {
}

func (p *PlanResource) GetAll(w http.ResponseWriter, r http.Request) {

}
