package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Metric struct {
	Name  string `yaml:"name"`
	Cmd   string `yaml:"cmd"`
	Value float64
}

func (m Metric) Shellout() (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", m.Cmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	out := strings.TrimRight(stdout.String(), "\r\n")
	log.Debug(m.Name, " output: ", out)
	return err, out, stderr.String()
}

func (m Metric) updateGauge() {
	log.Debug("setting ", m.Name, "u to ", m.Value)
	config.Gauge.WithLabelValues(m.Name, m.Cmd).Set(m.Value)
}

func (m Metric) updateValue() {
	log.Debug("Processing: ", m.Name)
	err, out, errout := m.Shellout()
	if err != nil {
		log.Error(out)
		log.Error(m)
	}
	if errout != "" {
		log.Error(errout)
	}
	s, err := strconv.ParseFloat(out, 64)
	if err != nil {
		log.Error("Could not get a float64 for ", m.Name)
	}
	log.Info(m.Name, " new value: ", s)
	m.Value = s
	m.updateGauge()
}
