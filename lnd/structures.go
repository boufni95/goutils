package lnd

import (
	"context"

	"github.com/lightningnetwork/lnd/lnrpc"
)

//Client - The client interface with all implemented requests and subs ~WIP
type Client interface {
	Connect(
		nodeURL string,
		macaroonPath string,
		certPath string,
	) error
	//GetInfo - info of the current node
	GetInfo(
		ctx context.Context,
	) (*GetInfoResponse, error)
	//ListPeers - Peers of the current node
	ListPeers(
		ctx context.Context,
	) (*ListPeersResponse, error)
	//VerifyMessage - verify a message from a member in the network
	VerifyMessage(
		ctx context.Context,
		message string,
		sign string,
	) (*VerifyMessageResponse, error)
	//AddInvoice - Add invoice on current node
	AddInvoice(
		ctx context.Context,
		memo string,
		valueSats int64,
		fallbackAddr string,
	) (*AddInvoiceResponse, error)
	//GetNodeInfo - Get info on a remote node
	GetNodeInfo(
		ctx context.Context,
		pubkey string,
		includeChannels bool,
	) (*GetNodeInfoResponse, error)
	SubscribeGraphAsync(
		ctx context.Context,
		Listener ChannelGraphListener,
	) error
	SubscribeInvoicesAsync(
		ctx context.Context,
		Listener InvoicesListener,
	) error
	SubscribeOpenChannel(
		ctx context.Context,
		pubkey string,
		capacity int64,
		Listener OpenChannelListener,
	) error
	/*ListChannels(
		ctx context.Context,
	) (*ListChannelsResponse, error)*/
}

type lndClient struct {
	client lnrpc.LightningClient
}

//GetInfoResponse - Info about current node
type GetInfoResponse struct {
	Res *lnrpc.GetInfoResponse
}

//ListPeersResponse - Info about current node's peers
type ListPeersResponse struct {
	Res *lnrpc.ListPeersResponse
}

//VerifyMessageResponse - Message owner if exists
type VerifyMessageResponse struct {
	Res *lnrpc.VerifyMessageResponse
}

//AddInvoiceResponse -
type AddInvoiceResponse struct {
	Res *lnrpc.AddInvoiceResponse
}

//GetNodeInfoResponse - Info about remote node
type GetNodeInfoResponse struct {
	Res *lnrpc.NodeInfo
}

//ListChannelsResponse - List open channels
type ListChannelsResponse struct {
	Res *lnrpc.ListChannelsResponse
}

//ChannelGraphListener - Callback to handle channel graph update or error
type ChannelGraphListener func(update *ChannelGraphUpdate, err error)

//ChannelGraphUpdate - Content of the channel graph update
type ChannelGraphUpdate struct {
	Update *lnrpc.GraphTopologyUpdate
}

//InvoicesListener - Callback to handle invoice update
type InvoicesListener func(update *InvoicesUpdate, err error)

//InvoicesUpdate - Content of invoice update
type InvoicesUpdate struct {
	Update *lnrpc.Invoice
}

//OpenChannelListener - Callback to handle open channel update
type OpenChannelListener func(update *OpenChannelUpdate, err error)

//OpenChannelUpdate - Content of open channel update
type OpenChannelUpdate struct {
	Update *lnrpc.OpenStatusUpdate
}
