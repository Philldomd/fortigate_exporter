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

func TestSystemSandboxConnection(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/sandbox/connection", "testdata/system-sandbox-connection.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemSandboxConnection, c, r) {
		t.Errorf("probeSystemSandboxConnection() returned non-success")
	}

	em := `
	# HELP fortigate_sandbox_connection_status_disabled Sandbox connection status
	# TYPE fortigate_sandbox_connection_status_disabled gauge
	fortigate_sandbox_connection_status_disabled{sandbox_type="appliance"} 0
	# HELP fortigate_sandbox_connection_status_incompatible Sandbox connection status
	# TYPE fortigate_sandbox_connection_status_incompatible gauge
	fortigate_sandbox_connection_status_incompatible{sandbox_type="appliance"} 0
	# HELP fortigate_sandbox_connection_status_reachable Sandbox connection status
	# TYPE fortigate_sandbox_connection_status_reachable gauge
	fortigate_sandbox_connection_status_reachable{sandbox_type="appliance"} 1
	# HELP fortigate_sandbox_connection_status_unauthorized Sandbox connection status
	# TYPE fortigate_sandbox_connection_status_unauthorized gauge
	fortigate_sandbox_connection_status_unauthorized{sandbox_type="appliance"} 0
	# HELP fortigate_sandbox_connection_status_unreachable Sandbox connection status
	# TYPE fortigate_sandbox_connection_status_unreachable gauge
	fortigate_sandbox_connection_status_unreachable{sandbox_type="appliance"} 0
	# HELP fortigate_sandbox_connection_status_unverified Sandbox connection status
	# TYPE fortigate_sandbox_connection_status_unverified gauge
	fortigate_sandbox_connection_status_unverified{sandbox_type="appliance"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
