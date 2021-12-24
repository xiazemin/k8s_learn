package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	TemperatureHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "beijing_temperature",
		Help:    "The temperature of the beijing",
		Buckets: prometheus.LinearBuckets(0, 10, 3),
	})
)

func InsertTemperature() {
	var temperature = [10]float64{1, 4, 5, 10, 14, 15, 20, 25, 11, 30}
	for i := 0; i < len(temperature); i++ {
		TemperatureHistogram.Observe(temperature[i])
		fmt.Printf("insert number: %f \n", temperature[i])
	}
}

func init() {
	prometheus.MustRegister(TemperatureHistogram)
}
