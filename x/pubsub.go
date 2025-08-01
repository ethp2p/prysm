package x

import pubsub "github.com/libp2p/go-libp2p-pubsub"

var InterceptPubsubOpts func([]pubsub.Option) []pubsub.Option = Identity
