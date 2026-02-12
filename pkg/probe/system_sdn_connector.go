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

type SystemSDNConnectorResults struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Updating   bool   `json:"updating"`
	LastUpdate int    `json:"last_update"`
}

type SystemSDNConnector struct {
	Results []SystemSDNConnectorResults `json:"results"`
	VDOM    string                      `json:"vdom"`
}

func probeSystemSDNConnector(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		SDNConnectorsStatusDisabled = prometheus.NewDesc(
			"fortigate_system_sdn_connector_disabled",
			"Status of SDN connectors disabled",
			[]string{"vdom", "name", "type"}, nil,
		)
		SDNConnectorsStatusDown = prometheus.NewDesc(
			"fortigate_system_sdn_connector_down",
			"Status of SDN connectors down",
			[]string{"vdom", "name", "type"}, nil,
		)
		SDNConnectorsStatusUnknown = prometheus.NewDesc(
			"fortigate_system_sdn_connector_unknown",
			"Status of SDN connectors unknown",
			[]string{"vdom", "name", "type"}, nil,
		)
		SDNConnectorsStatusUp = prometheus.NewDesc(
			"fortigate_system_sdn_connector_up",
			"Status of SDN connectors up",
			[]string{"vdom", "name", "type"}, nil,
		)
		SDNConnectorsStatusUpdating = prometheus.NewDesc(
			"fortigate_system_sdn_connector_updating",
			"Status of SDN connectors updating",
			[]string{"vdom", "name", "type"}, nil,
		)
		SDNConnectorsLastUpdate = prometheus.NewDesc(
			"fortigate_system_sdn_connector_last_update_seconds",
			"Last update time for SDN connectors (in seconds from epoch)",
			[]string{"vdom", "name", "type"}, nil,
		)
	)

	var res []SystemSDNConnector
	if err := c.Get("api/v2/monitor/system/sdn-connector/status", "vdom=*", &res); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, r := range res {
		for _, sdnConn := range r.Results {
			switch sdnConn.Status {
			case "Disabled":
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDisabled, prometheus.GaugeValue, float64(1), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUnknown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUp, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUpdating, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
			case "Down":
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDisabled, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDown, prometheus.GaugeValue, float64(1), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUnknown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUp, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUpdating, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
			case "Unknown":
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDisabled, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUnknown, prometheus.GaugeValue, float64(1), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUp, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUpdating, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
			case "Up":
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDisabled, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUnknown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUp, prometheus.GaugeValue, float64(1), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUpdating, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
			case "Updating":
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDisabled, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusDown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUnknown, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUp, prometheus.GaugeValue, float64(0), r.VDOM, sdnConn.Name, sdnConn.Type))
				m = append(m, prometheus.MustNewConstMetric(SDNConnectorsStatusUpdating, prometheus.GaugeValue, float64(1), r.VDOM, sdnConn.Name, sdnConn.Type))
			}
			m = append(m, prometheus.MustNewConstMetric(SDNConnectorsLastUpdate, prometheus.GaugeValue, float64(sdnConn.LastUpdate), r.VDOM, sdnConn.Name, sdnConn.Type))
		}
	}

	return m, true
}
