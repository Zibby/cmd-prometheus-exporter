package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var reg = prometheus.NewRegistry()

func initLog() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("logger initialised")
}

func setLogLeveL() {
	switch config.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "err":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Info("No valid log_level set in config, defaulting to Info")
		log.Info("Valid options are [debug, info, warn, err]")
	}
}

func loadYAML() {
	yfile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(yfile, &config)
	log.Info("Loaded config: ", config)
	setLogLeveL()
}

func createGauge() {
	config.Gauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cmd_output",
			Help: "Generates gauges from arbitary linux cmds",
		},
		[]string{"name"},
	)
	reg.Register(config.Gauge)
}

func init() {
	initLog()
	loadYAML()
	log.Debug("loaded yaml")
	createGauge()
	log.Debug("created all guages")
	log.Debug("init complete")
}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	config.processMetrics()
	promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/metrics", metricHandler)
	log.Fatal(http.ListenAndServe(config.ListenAddress, r))
}
