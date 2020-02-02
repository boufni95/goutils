package lnd

import (
	"io/ioutil"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/macaroon.v2"
)

//NewLndClient - return a new LND client, call Connect after creation
func NewLndClient() Client {
	lndC := lndClient{}
	return &lndC

}

func (Lnd *lndClient) Connect(
	nodeURL string,
	macaroonPath string,
	certPath string,
) error {
	tlsCreds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		return err
	}

	macaroonBytes, err := ioutil.ReadFile(macaroonPath)
	if err != nil {
		return err
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macaroonBytes); err != nil {
		return err
	}
	//#endregion

	//#region dial grpc
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(macaroons.NewMacaroonCredential(mac)),
	}

	conn, err := grpc.Dial(nodeURL, opts...)
	if err != nil {
		return err
	}

	Lnd.client = lnrpc.NewLightningClient(conn)
	//spew.Dump(l.lndClient)
	//#endregion
	return nil
}
