package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
)

func main() {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	httpReqs.WithLabelValues("404", "POST").Add(42)
	// If you have to access the same set of labels very frequently, it
	// might be good to retrieve the metric only once and keep a handle to
	// it. But beware of deletion of that metric, see below!
	m := httpReqs.WithLabelValues("200", "GET")
	for i := 0; i < 1000000; i++ {
		m.Inc()
	}
	// Delete a metric from the vector. If you have previously kept a handle
	// to that metric (as above), future updates via that handle will go
	// unseen (even if you re-create a metric with the same label set
	// later).
	httpReqs.DeleteLabelValues("200", "GET")
	// Same thing with the more verbose Labels syntax.
	httpReqs.Delete(prometheus.Labels{"method": "GET", "code": "200"})

	// 10 operations queued by the goroutine managing incoming requests.
	opsQueued.Add(10)
	// A worker goroutine has picked up a waiting operation.
	opsQueued.Dec()
	// And once more...
	opsQueued.Dec()

	// Simulate some observations.
	for i := 0; i < 1000; i++ {
		temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	}
	// Just for demonstration, let's check the state of the histogram by
	// (ab)using its Write method (which is usually only used by Prometheus
	// internally).
	metric := &dto.Metric{}
	temps.Write(metric)
	fmt.Println(proto.MarshalTextString(metric))

	// Simulate some observations.
	for i := 0; i < 1000; i++ {
		temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	}
	// Just for demonstration, let's check the state of the summary by
	// (ab)using its Write method (which is usually only used by Prometheus
	// internally).
	metricsummary := &dto.Metric{}
	tempsummary.Write(metric)
	fmt.Println(proto.MarshalTextString(metricsummary))

	InsertTemperature()
	InsertSummary()

	http.HandleFunc("/httpInc", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < rand.Int()%100; i++ {
			hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
		}
		hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

		cpuTemp.Set(65.3 + rand.Float64()*100)

		InsertTemperature()
		InsertSummary()

		fmt.Fprintf(w, "httpInc")
		return
	})
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":12345", nil))
}
