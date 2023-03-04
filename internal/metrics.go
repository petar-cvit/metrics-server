package internal

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	dataFile       = "data/metrics_from_special_app.txt"
	metricsKey     = "metrics"
	secondsInCache = 10
)

// Metrics uses sync map for consistent concurrent calls
type Metrics struct {
	storage sync.Map
	expires int64
}

// NewMetricsStorage constructor
func NewMetricsStorage() *Metrics {
	m := &Metrics{
		storage: sync.Map{},
		expires: 0,
	}

	m.refreshCache()

	return m
}

// GetMetrics returns metrics read from file and returns as string
func (m *Metrics) GetMetrics() string {
	// cache miss
	if m.staleMetrics() {
		m.refreshCache()
	}

	metricsEntry, exists := m.storage.Load(metricsKey)
	if !exists {
		return "no metrics"
	}

	return metricsEntry.(string)
}

func (m *Metrics) staleMetrics() bool {
	return time.Now().Unix() > m.expires
}

// refreshCache reads file with metrics and stores it in memory
func (m *Metrics) refreshCache() {
	fmt.Println("refreshing cache")

	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("error reading from file", dataFile, err)
	}

	m.storage.Store(metricsKey, string(bytes))
	m.expires = time.Now().Unix() + secondsInCache
}
