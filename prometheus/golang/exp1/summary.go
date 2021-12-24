package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	SalarySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "beijing_salary",
		Help:       "the relationship between salary and population of beijing city",
		Objectives: map[float64]float64{0.5: 0.05, 0.8: 0.001, 0.9: 0.01, 0.95: 0.01},
	})
)

func InsertSummary() {
	var salary = [10]float64{8000, 7000, 8900, 10000, 9800, 17000, 15000, 14000, 11000, 12000}
	for i := 0; i < len(salary); i++ {
		SalarySummary.Observe(salary[i])
		fmt.Printf("Insert number: %f \n", salary[i])
	}
}
