package lnd

import (
	"context"
	"encoding/hex"

	"github.com/lightningnetwork/lnd/lnrpc"
)

func (Lnd *lndClient) SubscribeGraphAsync(ctx context.Context, Listener ChannelGraphListener) error {
	sub, err := Lnd.client.SubscribeChannelGraph(ctx, &lnrpc.GraphTopologySubscription{})
	if err != nil {
		return err
	}
	go func(sub lnrpc.Lightning_SubscribeChannelGraphClient, Listener ChannelGraphListener) {
		for {
			update, err := sub.Recv()
			if err != nil {
				Listener(nil, err)
				return
			}
			up := ChannelGraphUpdate{
				Update: update,
			}
			Listener(&up, nil)
		}
	}(sub, Listener)
	return nil
}

func (Lnd *lndClient) SubscribeInvoicesAsync(ctx context.Context, Listener InvoicesListener) error {
	sub, err := Lnd.client.SubscribeInvoices(ctx, &lnrpc.InvoiceSubscription{})
	if err != nil {
		return err
	}
	go func(sub lnrpc.Lightning_SubscribeInvoicesClient, Listener InvoicesListener) {
		for {
			update, err := sub.Recv()
			if err != nil {
				Listener(nil, err)
				return
			}
			up := InvoicesUpdate{
				Update: update,
			}
			Listener(&up, err)
		}
	}(sub, Listener)
	return nil
}

func (Lnd *lndClient) SubscribeOpenChannel(ctx context.Context, pubkey string, capacity int64, Listener OpenChannelListener) error {
	b, err := hex.DecodeString(pubkey)
	if err != nil {
		return err
	}
	params := lnrpc.OpenChannelRequest{
		NodePubkey:         b,
		LocalFundingAmount: capacity,
	}
	sub, err := Lnd.client.OpenChannel(ctx, &params)
	if err != nil {
		return err
	}
	go func(sub lnrpc.Lightning_OpenChannelClient, Listener OpenChannelListener) {
		for {
			update, err := sub.Recv()
			if err != nil {
				Listener(nil, err)
				return
			}
			up := OpenChannelUpdate{
				Update: update,
			}
			Listener(&up, nil)

		}
	}(sub, Listener)
	return nil
}
