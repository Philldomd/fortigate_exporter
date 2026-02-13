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
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeSystemSandboxConnection(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	connectionStatusDisable := prometheus.NewDesc(
		"fortigate_sandbox_connection_status_disabled",
		"Sandbox connection status",
		[]string{"sandbox_type"}, nil,
	)
	connectionStatusUreachable := prometheus.NewDesc(
		"fortigate_sandbox_connection_status_unreachable",
		"Sandbox connection status",
		[]string{"sandbox_type"}, nil,
	)
	connectionStatusReachable := prometheus.NewDesc(
		"fortigate_sandbox_connection_status_reachable",
		"Sandbox connection status",
		[]string{"sandbox_type"}, nil,
	)
	connectionStatusUnauthorized := prometheus.NewDesc(
		"fortigate_sandbox_connection_status_unauthorized",
		"Sandbox connection status",
		[]string{"sandbox_type"}, nil,
	)
	connectionStatusIncompatible := prometheus.NewDesc(
		"fortigate_sandbox_connection_status_incompatible",
		"Sandbox connection status",
		[]string{"sandbox_type"}, nil,
	)
	connectionStatusUnverified := prometheus.NewDesc(
		"fortigate_sandbox_connection_status_unverified",
		"Sandbox connection status",
		[]string{"sandbox_type"}, nil,
	)

	type SystemSandboxConnection struct {
		Status string `json:"status"`
		Type   string `json:"type"`
	}

	type SystemSandboxConnectionResult struct {
		Results []SystemSandboxConnection `json:"results"`
	}
	var res SystemSandboxConnectionResult
	if err := c.Get("api/v2/monitor/system/sandbox/connection", "", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, r := range res.Results {
		switch r.Status {
		case "unreachable":
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUreachable, prometheus.GaugeValue, 1, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusReachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusDisable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnauthorized, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusIncompatible, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnverified, prometheus.GaugeValue, 0, r.Type))
		case "reachable":
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUreachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusReachable, prometheus.GaugeValue, 1, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusDisable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnauthorized, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusIncompatible, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnverified, prometheus.GaugeValue, 0, r.Type))
		case "disabled":
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUreachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusReachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusDisable, prometheus.GaugeValue, 1, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnauthorized, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusIncompatible, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnverified, prometheus.GaugeValue, 0, r.Type))
		case "unauthorized":
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUreachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusReachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusDisable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnauthorized, prometheus.GaugeValue, 1, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusIncompatible, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnverified, prometheus.GaugeValue, 0, r.Type))
		case "incompatible":
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUreachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusReachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusDisable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnauthorized, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusIncompatible, prometheus.GaugeValue, 1, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnverified, prometheus.GaugeValue, 0, r.Type))
		case "unverified":
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUreachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusReachable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusDisable, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnauthorized, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusIncompatible, prometheus.GaugeValue, 0, r.Type))
			m = append(m, prometheus.MustNewConstMetric(connectionStatusUnverified, prometheus.GaugeValue, 1, r.Type))
		}
	}
	return m, true
}
