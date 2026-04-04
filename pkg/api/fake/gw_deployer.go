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

package fake

import (
	"context"

	"github.com/submariner-io/cloud-prepare/pkg/api"
	"github.com/submariner-io/admiral/pkg/reporter"
)

type GatewayDeployer struct {
	CleanupInvoked             bool
	CapturedGatewayDeployInput *api.GatewayDeployInput
	ReturnError                error
}

func (g *GatewayDeployer) Deploy(_ context.Context, input api.GatewayDeployInput, _ reporter.Interface) error {
	g.CapturedGatewayDeployInput = &input
	return g.ReturnError
}

func (g *GatewayDeployer) Cleanup(_ context.Context, _ reporter.Interface) error {
	g.CleanupInvoked = true
	return g.ReturnError
}
