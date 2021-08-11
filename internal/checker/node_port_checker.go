package checker

import (
	"errors"
	"log"
	"time"

	goTezos "github.com/goat-systems/go-tezos/v4/rpc"
)

const (
	publicRpc = "http://mainnet-tezos.giganode.io"
	localRpc  = "http://localhost:8732"
)

type nodeUrlChecker struct {
	localRpcCli *goTezos.Client
	// publicRpcCli *goTezos.Client
}

func (c *nodeUrlChecker) AssertRunning() (err error) {
	_, head, err := c.localRpcCli.Block(&goTezos.BlockIDHead{})
	if err != nil {
		return err
	}
	log.Printf("Current head timestamp %v", head.Header.Timestamp)
	if head.Header.Timestamp.Sub(time.Now()).Minutes() > 10.0 {
		return errors.New("node is unsync")
	}
	return nil
}

func NewNodePortChecker() (c Checker, err error) {
	c = &nodeUrlChecker{}
	//c.(*nodeUrlChecker).publicRpcCli, err = goTezos.New(publicRpc)
	//if err != nil {
	//	return nil, err
	//}
	c.(*nodeUrlChecker).localRpcCli, err = goTezos.New(localRpc)
	if err != nil {
		return nil, err
	}
	return c, nil
}
