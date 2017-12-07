package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	neighbors = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "neighbors",
		Help: "Current number of neighbors",
	})
)

func initPrometheusClient() {
	prometheus.MustRegister(neighbors)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)
}

func setNeighborMetric(num int) {
	neighbors.Set(float64(num))
}
