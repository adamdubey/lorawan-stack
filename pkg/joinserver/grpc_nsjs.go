// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type nsJsServer struct {
	JS *JoinServer
}

// HandleJoin is called by the Network Server to join a device.
func (srv nsJsServer) HandleJoin(ctx context.Context, req *ttnpb.JoinRequest) (res *ttnpb.JoinResponse, err error) {
	return srv.JS.HandleJoin(ctx, req, ClusterAuthorizer(ctx))
}

// GetNwkSKeys returns the NwkSKeys associated with session keys identified by the supplied request.
func (srv nsJsServer) GetNwkSKeys(ctx context.Context, req *ttnpb.SessionKeyRequest) (*ttnpb.NwkSKeysResponse, error) {
	return srv.JS.GetNwkSKeys(ctx, req, ClusterAuthorizer(ctx))
}
