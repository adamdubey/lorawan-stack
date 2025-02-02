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

package mac

import (
	"context"

	pbtypes "github.com/gogo/protobuf/types"

	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	EvtReceiveDeviceTimeRequest = defineReceiveMACRequestEvent(
		"device_time", "device time",
	)()
	EvtEnqueueDeviceTimeAnswer = defineEnqueueMACAnswerEvent(
		"device_time", "device time",
		events.WithData(&ttnpb.MACCommand_DeviceTimeAns{}),
	)()
)

func HandleDeviceTimeReq(ctx context.Context, dev *ttnpb.EndDevice, msg *ttnpb.UplinkMessage) (events.Builders, error) {
	ans := &ttnpb.MACCommand_DeviceTimeAns{
		Time: messageTimestamp(msg),
	}
	dev.MacState.QueuedResponses = append(dev.MacState.QueuedResponses, ans.MACCommand())
	return events.Builders{
		EvtReceiveDeviceTimeRequest,
		EvtEnqueueDeviceTimeAnswer.With(events.WithData(ans)),
	}, nil
}

func messageTimestamp(msg *ttnpb.UplinkMessage) *pbtypes.Timestamp {
	var ts *pbtypes.Timestamp
	for _, md := range msg.RxMetadata {
		if t := md.GpsTime; t != nil {
			return t
		}
		if ts == nil && md.ReceivedAt != nil {
			ts = md.ReceivedAt
		}
	}
	if ts != nil {
		return ts
	}
	return msg.ReceivedAt
}
