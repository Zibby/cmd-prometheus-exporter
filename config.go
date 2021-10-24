package main

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const ShellToUse = "bash"

type Config struct {
	LogLevel      string   `yaml:"log_level"`
	ListenAddress string   `yaml:"listen_address"`
	Metrics       []Metric `yaml:"metrics"`
	Gauge         *prometheus.GaugeVec
}

var config Config

func (c Config) processMetrics() {
	for _, metric := range c.Metrics {
		log.Debug("processing: ", metric.Name)
		metric.updateValue()
	}
}
