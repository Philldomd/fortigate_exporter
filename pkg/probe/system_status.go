// Copyright The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package probe

import (
	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeSystemStatus(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	mVersion := prometheus.NewDesc(
		"fortigate_version_info",
		"System version and build information",
		[]string{"serial", "version", "build", "name", "number", "module", "hostname"}, nil,
	)
	mLogDiskAvailable := prometheus.NewDesc(
		"fortigate_system_status_log_disk_available",
		"System log disk availability status",
		[]string{"serial", "version", "build", "name", "number", "module", "hostname"}, nil,
	)
	mLogDiskNeedFormat := prometheus.NewDesc(
		"fortigate_system_status_log_disk_need_format",
		"System log disk availability status",
		[]string{"serial", "version", "build", "name", "number", "module", "hostname"}, nil,
	)
	mLogDiskNotAvailable := prometheus.NewDesc(
		"fortigate_system_status_log_disk_not_available",
		"System log disk availability status",
		[]string{"serial", "version", "build", "name", "number", "module", "hostname"}, nil,
	)

	type systemResult struct {
		Name          string `json:"model_name"`
		Number        string `json:"model_number"`
		Model         string `json:"model"`
		Hostname      string `json:"hostname"`
		LogDiskStatus string `json:"log_disk_status"`
	}

	type systemStatus struct {
		Status  string
		Serial  string
		Version string
		Build   int64
		Results systemResult
	}
	var st systemStatus

	if err := c.Get("api/v2/monitor/system/status", "", &st); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	switch st.Results.LogDiskStatus {
	case "available":
		m = append(m, prometheus.MustNewConstMetric(mLogDiskAvailable, prometheus.GaugeValue, 1.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
		m = append(m, prometheus.MustNewConstMetric(mLogDiskNeedFormat, prometheus.GaugeValue, 0.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
		m = append(m, prometheus.MustNewConstMetric(mLogDiskNotAvailable, prometheus.GaugeValue, 0.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
	case "need_format":
		m = append(m, prometheus.MustNewConstMetric(mLogDiskAvailable, prometheus.GaugeValue, 0.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
		m = append(m, prometheus.MustNewConstMetric(mLogDiskNeedFormat, prometheus.GaugeValue, 1.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
		m = append(m, prometheus.MustNewConstMetric(mLogDiskNotAvailable, prometheus.GaugeValue, 0.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
	case "not_available":
		m = append(m, prometheus.MustNewConstMetric(mLogDiskAvailable, prometheus.GaugeValue, 0.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
		m = append(m, prometheus.MustNewConstMetric(mLogDiskNeedFormat, prometheus.GaugeValue, 0.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
		m = append(m, prometheus.MustNewConstMetric(mLogDiskNotAvailable, prometheus.GaugeValue, 1.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
	}
	m = append(m, prometheus.MustNewConstMetric(mVersion, prometheus.GaugeValue, 1.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
	return m, true
}
