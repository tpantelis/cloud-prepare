/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"k8s.io/utils/ptr"
)

func ec2Filter(name, value string) types.Filter {
	return types.Filter{
		Name:   ptr.To(name),
		Values: []string{value},
	}
}

func ec2Tag(key, value string) types.Tag {
	return types.Tag{
		Key:   ptr.To(key),
		Value: ptr.To(value),
	}
}

func ec2FilterByTag(tag types.Tag) types.Filter {
	return ec2Filter("tag:"+*tag.Key, *tag.Value)
}

func hasTag(tags []types.Tag, desired types.Tag) bool {
	for _, tag := range tags {
		if *tag.Key == *desired.Key {
			return true
		}
	}

	return false
}

func extractName(tags []types.Tag) string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}

	return ""
}

func (ac *awsCloud) withAWSInfo(str string) string {
	r := strings.NewReplacer("{infraID}", ac.infraID, "{region}", ac.region)
	return r.Replace(str)
}

func (ac *awsCloud) filterByName(name string) types.Filter {
	return ec2Filter("tag:Name", ac.withAWSInfo(name))
}

func (ac *awsCloud) filterByCurrentCluster() []types.Filter {
	return []types.Filter{
		ec2Filter(ac.withAWSInfo("tag:kubernetes.io/cluster/{infraID}"), "owned"),
		ec2Filter(ac.withAWSInfo("tag:sigs.k8s.io/cluster-api-provider-aws/cluster/{infraID}"), "owned"),
	}
}
