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

package service

import (
	"context"

	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"gocloud.dev/pubsub"
)

// DownlinkSubscriptions contains the subscriptions for the push and replace queue operations.
type DownlinkSubscriptions struct {
	Push    *pubsub.Subscription
	Replace *pubsub.Subscription
}

// Shutdown shutdowns the active subscriptions.
func (ds *DownlinkSubscriptions) Shutdown(ctx context.Context) error {
	for _, sub := range []*pubsub.Subscription{
		ds.Push,
		ds.Replace,
	} {
		if sub != nil {
			if err := sub.Shutdown(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

// UplinkTopics contains the topics for the uplink messages.
type UplinkTopics struct {
	UplinkMessage  *pubsub.Topic
	JoinAccept     *pubsub.Topic
	DownlinkAck    *pubsub.Topic
	DownlinkNack   *pubsub.Topic
	DownlinkSent   *pubsub.Topic
	DownlinkFailed *pubsub.Topic
	DownlinkQueued *pubsub.Topic
	LocationSolved *pubsub.Topic
}

// Shutdown shutdowns the active topics.
func (ut *UplinkTopics) Shutdown(ctx context.Context) error {
	for _, topic := range []*pubsub.Topic{
		ut.UplinkMessage,
		ut.JoinAccept,
		ut.DownlinkAck,
		ut.DownlinkNack,
		ut.DownlinkSent,
		ut.DownlinkFailed,
		ut.DownlinkQueued,
		ut.LocationSolved,
	} {
		if topic != nil {
			if err := topic.Shutdown(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

// Service represents a PubSub service provided by go-cloud pubsub.
type Service interface {
	// OpenSubscriptions opens the subscriptions for the downlink queue operations of a given ttnpb.ApplicationPubSub.
	OpenSubscriptions(ctx context.Context, pubsub *ttnpb.ApplicationPubSub) (*DownlinkSubscriptions, error)
	// OpenTopics opens the subscriptions for the uplink messages of a given ttnpb.ApplicationPubSub.
	OpenTopics(ctx context.Context, pubsub *ttnpb.ApplicationPubSub) (*UplinkTopics, error)
}
