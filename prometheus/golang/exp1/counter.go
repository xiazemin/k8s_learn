package main

import "github.com/prometheus/client_golang/prometheus"

var (
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)
)

func init() {
	prometheus.MustRegister(hdFailures)
}
