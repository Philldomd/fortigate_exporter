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

func TestBGPNeighborsIPv4Old(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/router/bgp/neighbors", "testdata/router-bgp-neighbors-v4old.jsonnet")
	r := prometheus.NewPedanticRegistry()
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 4,
	}
	if !testProbeWithMetadata(probeBGPNeighborsIPv4, c, meta, r) {
		t.Errorf("probeBGPNeighborsIPv4() returned non-success")
	}

	em := `
	# HELP fortigate_bgp_neighbor_ipv4_info Configured bgp neighbor over ipv4, return state as value (1 - Idle, 2 - Connect, 3 - Active, 4 - Open sent, 5 - Open confirm, 6 - Established)
	# TYPE fortigate_bgp_neighbor_ipv4_info gauge
	fortigate_bgp_neighbor_ipv4_info{admin_status="true",local_ip="10.0.0.0",neighbor_ip="10.0.0.1",remote_as="1337",state="Established",vdom="root"} 6
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestBGPNeighborsIPv6Old(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/router/bgp/neighbors6", "testdata/router-bgp-neighbors-v6old.jsonnet")
	r := prometheus.NewPedanticRegistry()
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 4,
	}
	if !testProbeWithMetadata(probeBGPNeighborsIPv6, c, meta, r) {
		t.Errorf("probeBGPNeighborsIPv6() returned non-success")
	}

	em := `
	# HELP fortigate_bgp_neighbor_ipv6_info Configured bgp neighbor over ipv6, return state as value (1 - Idle, 2 - Connect, 3 - Active, 4 - Open sent, 5 - Open confirm, 6 - Established)
	# TYPE fortigate_bgp_neighbor_ipv6_info gauge
	fortigate_bgp_neighbor_ipv6_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::2",remote_as="1337",state="Established",vdom="root"} 6
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestBGPNeighborsIPv4(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/router/bgp/neighbors", "testdata/router-bgp-neighbors-v4.jsonnet")
	r := prometheus.NewPedanticRegistry()
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 6,
	}
	if !testProbeWithMetadata(probeBGPNeighborsIPv4, c, meta, r) {
		t.Errorf("probeBGPNeighborsIPv4() returned non-success")
	}

	em := `
	# HELP fortigate_bgp_neighbor_ipv4_active_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv4_active_info gauge
	fortigate_bgp_neighbor_ipv4_active_info{admin_status="true",local_ip="10.0.0.0",neighbor_ip="10.0.0.1",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_active_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 1
	fortigate_bgp_neighbor_ipv4_active_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv4_connected_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv4_connected_info gauge
	fortigate_bgp_neighbor_ipv4_connected_info{admin_status="true",local_ip="10.0.0.0",neighbor_ip="10.0.0.1",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_connected_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_connected_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv4_established_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv4_established_info gauge
	fortigate_bgp_neighbor_ipv4_established_info{admin_status="true",local_ip="10.0.0.0",neighbor_ip="10.0.0.1",remote_as="1337",vdom="root"} 1
	fortigate_bgp_neighbor_ipv4_established_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_established_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv4_idle_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv4_idle_info gauge
	fortigate_bgp_neighbor_ipv4_idle_info{admin_status="true",local_ip="10.0.0.0",neighbor_ip="10.0.0.1",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_idle_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_idle_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv4_open_confirm_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv4_open_confirm_info gauge
	fortigate_bgp_neighbor_ipv4_open_confirm_info{admin_status="true",local_ip="10.0.0.0",neighbor_ip="10.0.0.1",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_open_confirm_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_open_confirm_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv4_open_sent_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv4_open_sent_info gauge
	fortigate_bgp_neighbor_ipv4_open_sent_info{admin_status="true",local_ip="10.0.0.0",neighbor_ip="10.0.0.1",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_open_sent_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv4_open_sent_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestBGPNeighborsIPv6(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/router/bgp/neighbors6", "testdata/router-bgp-neighbors-v6.jsonnet")
	r := prometheus.NewPedanticRegistry()
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 6,
	}
	if !testProbeWithMetadata(probeBGPNeighborsIPv6, c, meta, r) {
		t.Errorf("probeBGPNeighborsIPv6() returned non-success")
	}

	em := `
	# HELP fortigate_bgp_neighbor_ipv6_active_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv6_active_info gauge
	fortigate_bgp_neighbor_ipv6_active_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::2",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_active_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_active_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv6_connected_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv6_connected_info gauge
	fortigate_bgp_neighbor_ipv6_connected_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::2",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_connected_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_connected_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv6_established_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv6_established_info gauge
	fortigate_bgp_neighbor_ipv6_established_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::2",remote_as="1337",vdom="root"} 1
	fortigate_bgp_neighbor_ipv6_established_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_established_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv6_idle_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv6_idle_info gauge
	fortigate_bgp_neighbor_ipv6_idle_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::2",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_idle_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_idle_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv6_open_confirm_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv6_open_confirm_info gauge
	fortigate_bgp_neighbor_ipv6_open_confirm_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::2",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_open_confirm_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 1
	fortigate_bgp_neighbor_ipv6_open_confirm_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 0
	# HELP fortigate_bgp_neighbor_ipv6_open_sent_info Configured bgp neighbor over ipv4 state
	# TYPE fortigate_bgp_neighbor_ipv6_open_sent_info gauge
	fortigate_bgp_neighbor_ipv6_open_sent_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::2",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_open_sent_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::5",remote_as="1337",vdom="root"} 0
	fortigate_bgp_neighbor_ipv6_open_sent_info{admin_status="true",local_ip="fd00::1",neighbor_ip="fd00::7",remote_as="1337",vdom="root"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
