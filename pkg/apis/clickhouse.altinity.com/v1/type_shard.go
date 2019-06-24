// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

func (shard *ChiShard) InheritTemplates(cluster *ChiCluster) {
	(&shard.Templates).MergeFrom(&cluster.Templates)
}

func (shard *ChiShard) GetServiceTemplate() (*ChiServiceTemplate, bool) {
	name := shard.Templates.ShardServiceTemplate
	template, ok := shard.Chi.GetServiceTemplate(name)
	return template, ok
}

func (shard *ChiShard) WalkReplicas(
	f func(replica *ChiReplica) error,
) []error {
	res := make([]error, 0)

	for replicaIndex := range shard.Replicas {
		replica := &shard.Replicas[replicaIndex]
		res = append(res, f(replica))
	}

	return res
}