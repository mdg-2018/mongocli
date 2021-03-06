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

package cli

import (
	"github.com/mongodb/mongocli/internal/description"
	"github.com/spf13/cobra"
)

func AtlasClustersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "clusters",
		Aliases:    []string{"cluster"},
		SuggestFor: []string{"replicasets"},
		Short:      description.Clusters,
		Long:       description.ClustersLong,
	}
	cmd.AddCommand(AtlasClustersCreateBuilder())
	cmd.AddCommand(AtlasClustersListBuilder())
	cmd.AddCommand(AtlasClustersDescribeBuilder())
	cmd.AddCommand(AtlasClustersDeleteBuilder())
	cmd.AddCommand(AtlasClustersUpdateBuilder())

	return cmd
}
