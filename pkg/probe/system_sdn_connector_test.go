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
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestSystemSDNConnector(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/sdn-connector/status", "testdata/system-sdn-connector.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemSDNConnector, c, r) {
		t.Errorf("probeSystemSDNConnector() returned non-success")
	}

	em := `
	# HELP fortigate_system_sdn_connector_disabled Status of SDN connectors disabled
	# TYPE fortigate_system_sdn_connector_disabled gauge
	fortigate_system_sdn_connector_disabled{name="AWS Infra",type="aws",vdom="root"} 0
	fortigate_system_sdn_connector_disabled{name="GCP Infra",type="gcp",vdom="google"} 0
	# HELP fortigate_system_sdn_connector_down Status of SDN connectors down
	# TYPE fortigate_system_sdn_connector_down gauge
	fortigate_system_sdn_connector_down{name="AWS Infra",type="aws",vdom="root"} 0
	fortigate_system_sdn_connector_down{name="GCP Infra",type="gcp",vdom="google"} 1
	# HELP fortigate_system_sdn_connector_last_update_seconds Last update time for SDN connectors (in seconds from epoch)
	# TYPE fortigate_system_sdn_connector_last_update_seconds gauge
	fortigate_system_sdn_connector_last_update_seconds{name="AWS Infra",type="aws",vdom="root"} 1.680708575e+09
	fortigate_system_sdn_connector_last_update_seconds{name="GCP Infra",type="gcp",vdom="google"} 1.680708001e+09
	# HELP fortigate_system_sdn_connector_unknown Status of SDN connectors unknown
	# TYPE fortigate_system_sdn_connector_unknown gauge
	fortigate_system_sdn_connector_unknown{name="AWS Infra",type="aws",vdom="root"} 0
	fortigate_system_sdn_connector_unknown{name="GCP Infra",type="gcp",vdom="google"} 0
	# HELP fortigate_system_sdn_connector_up Status of SDN connectors up
	# TYPE fortigate_system_sdn_connector_up gauge
	fortigate_system_sdn_connector_up{name="AWS Infra",type="aws",vdom="root"} 1
	fortigate_system_sdn_connector_up{name="GCP Infra",type="gcp",vdom="google"} 0
	# HELP fortigate_system_sdn_connector_updating Status of SDN connectors updating
	# TYPE fortigate_system_sdn_connector_updating gauge
	fortigate_system_sdn_connector_updating{name="AWS Infra",type="aws",vdom="root"} 0
	fortigate_system_sdn_connector_updating{name="GCP Infra",type="gcp",vdom="google"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
