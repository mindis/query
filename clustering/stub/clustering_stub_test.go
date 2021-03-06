//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package clustering_stub

import (
	"testing"

	"github.com/couchbaselabs/query/accounting/stub"
)

func TestStub(t *testing.T) {
	cs, _ := NewConfigurationStore()

	c, _ := cs.ClusterByName("cluster_id")
	if c != nil {
		t.Fatalf("Expected nil cluster")
	}

	cnames, _ := cs.ClusterNames()
	if len(cnames) != 1 {
		t.Fatalf("Expected length of cluster names to be one")
	}

	c, _ = cs.ClusterByName(cnames[0])
	if c == nil {
		t.Fatalf("Expected to retrieve cluster using name from ClusterNames()")
	}

	if c.ConfigurationStoreId() != cs.Id() {
		t.Fatalf("Cluster does not have expected configuration store ID")
	}

	qnames, _ := c.QueryNodeNames()
	if len(qnames) != 1 {
		t.Fatalf("Expected length of Query Node names to be one")
	}

	q, _ := c.QueryNodeByName(qnames[0])
	if q == nil {
		t.Fatalf("Expected to retrieve Query Node using name from QueryNodeNames()")
	}

	if q.Cluster().Name() != c.Name() {
		t.Fatalf("Unexpected cluster name in Query Node: %v", q.Cluster().Name())
	}

	as := q.Cluster().AccountingStore()

	if as.Id() != c.AccountingStore().Id() {
		t.Fatalf("Unexpected Accounting store id in Query Node: %v", as.Id())
	}

	mr := as.MetricRegistry()

	mr.Register("metric1", accounting_stub.GaugeStub{})

	g := mr.Get("metric1")
	if g != nil {
		t.Fatalf("MetricsRegsitryStub should not have any state")
	}

	gauges := mr.Gauges()

	for k, v := range gauges {
		t.Fatalf("Gauges map should be empty, found values: %v, %v", k, v)
	}
}
