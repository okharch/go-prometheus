package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_cpu_usage",
		Help: "CPU usage of the application",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_memory_usage",
		Help: "Memory usage of the application",
	})
)

func init() {
	// Register custom metrics
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
}

func monitorResources() {
	for {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		memoryUsage.Set(float64(memStats.Alloc) / 1024 / 1024) // Convert bytes to MB
		// Simulate CPU usage (this can be replaced with real CPU usage metrics)
		cpuUsage.Set(float64(time.Now().Unix() % 100))
		time.Sleep(time.Second)
	}
}

func allocateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	size := r.FormValue("size")
	duration := r.FormValue("duration")

	var mem []byte
	if size != "" {
		allocSize := 1024 * 1024 * 1024 // Default 1MB
		fmt.Sscanf(size, "%d", &allocSize)
		mem = make([]byte, allocSize) // Allocate memory
	}

	if duration != "" {
		delay := 5
		fmt.Sscanf(duration, "%d", &delay)
		time.Sleep(time.Duration(delay) * time.Second) // Simulate CPU usage
	}

	fmt.Fprintf(w, "Allocated %d bytes of memory and simulated CPU usage for %s seconds", len(mem), duration)
}

func main() {
	// Start monitoring resources
	go monitorResources()

	// Expose metrics and handlers
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/allocate", allocateHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
