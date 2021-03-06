// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search_test

import (
	"fmt"

	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/search"
)

// This example demonstrates searching a list.
func ExampleStringInSlice() {
	a := []string{"a", "b", "c"}
	x := "a"

	if search.StringInSlice(a, x) {
		fmt.Printf("found %s in %v\n", x, a)
	} else {
		fmt.Printf("%s not found in %v\n", x, a)
	}
	// Output:
	// found a in [a b c]
}

// This example demonstrates searching a list of processes by name.
func ExampleProcesses() {
	a := fixtures.AutomationConfig().Processes
	x := "myReplicaSet_2"
	i, found := search.Processes(a, func(p *om.Process) bool { return p.Name == x })
	if i < len(a) && found {
		fmt.Printf("found %v at index %d\n", x, i)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet_2 at index 1
}

// This example demonstrates searching a list of replica sets by ID.
func ExampleReplicaSets() {
	a := fixtures.AutomationConfig().ReplicaSets
	x := "myReplicaSet"
	i, found := search.ReplicaSets(a, func(r *om.ReplicaSet) bool { return r.ID == x })
	if i < len(a) && found {
		fmt.Printf("found %v at index %d\n", x, i)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet at index 0
}

// This example demonstrates searching a list of members by host.
func ExampleMembers() {
	a := fixtures.AutomationConfig().ReplicaSets[0].Members
	x := "myReplicaSet_2"
	i, found := search.Members(a, func(m om.Member) bool { return m.Host == x })
	if i < len(a) && found {
		fmt.Printf("found %v at index %d\n", x, i)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet_2 at index 1
}

// This example demonstrates searching a list of db users by username.
func ExampleMongoDBUsers() {
	a := fixtures.AutomationConfigWithMongoDBUsers().Auth.Users
	x := "test"
	i, found := search.MongoDBUsers(a, func(m *om.MongoDBUser) bool { return m.Username == x })
	if i < len(a) && found {
		fmt.Printf("found %v at index %d\n", x, i)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found test at index 0
}

// This example demonstrates searching a cluster in an automation config.
func ExampleClusterExists() {
	a := fixtures.AutomationConfig()
	x := "myReplicaSet"
	found := search.ClusterExists(a, x)
	if found {
		fmt.Printf("found %v\n", x)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet
}
