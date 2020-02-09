package lnd

import (
	"context"

	"github.com/lightningnetwork/lnd/lnrpc"
)

func (Lnd *lndClient) GetInfo(ctx context.Context) (*GetInfoResponse, error) {
	getInfoReq := lnrpc.GetInfoRequest{}
	res, err := Lnd.client.GetInfo(ctx, &getInfoReq)
	if err != nil {
		return nil, err
	}
	in := GetInfoResponse{
		Res: res,
	}
	return &in, nil
}

func (Lnd *lndClient) ListPeers(ctx context.Context) (*ListPeersResponse, error) {
	peersReq := &lnrpc.ListPeersRequest{}
	peersRes, err := Lnd.client.ListPeers(ctx, peersReq)
	if err != nil {
		return nil, err
	}
	pe := ListPeersResponse{
		Res: peersRes,
	}
	return &pe, nil
}

func (Lnd *lndClient) VerifyMessage(ctx context.Context, message string, sign string) (*VerifyMessageResponse, error) {
	verReq := &lnrpc.VerifyMessageRequest{
		Msg:       []byte(message),
		Signature: sign,
	}
	verifyMessageResp, err := Lnd.client.VerifyMessage(ctx, verReq)
	if err != nil {
		return nil, err
	}
	ve := VerifyMessageResponse{
		Res: verifyMessageResp,
	}
	return &ve, nil
}

func (Lnd *lndClient) AddInvoice(ctx context.Context, memo string, valueSats int64, fallbackAddr string) (*AddInvoiceResponse, error) {
	invoice := &lnrpc.Invoice{
		Memo:         memo,
		Value:        valueSats,
		FallbackAddr: fallbackAddr,
	}
	addInvoiceResp, err := Lnd.client.AddInvoice(ctx, invoice)
	if err != nil {
		return nil, err
	}
	in := AddInvoiceResponse{
		Res: addInvoiceResp,
	}
	return &in, nil
}

func (Lnd *lndClient) GetNodeInfo(ctx context.Context, pubkey string, includeChannels bool) (*GetNodeInfoResponse, error) {
	getReq := &lnrpc.NodeInfoRequest{
		PubKey:          pubkey,
		IncludeChannels: includeChannels,
	}
	getNodeInfoResp, err := Lnd.client.GetNodeInfo(ctx, getReq)

	if err != nil {
		return nil, err
	}
	in := GetNodeInfoResponse{
		Res: getNodeInfoResp,
	}
	return &in, nil
}

/*
func (Lnd *lndClient) ListChannels(ctx context.Context) (*ListChannelsResponse, error) {
	chans, err := Lnd.client.ListChannels(ctx, &lnrpc.ListChannelsRequest{})
	if err != nil {
		return nil, err
	}
	li := ListChannelsResponse{
		Res: chans,
	}
	return &li, nil
}

*/
