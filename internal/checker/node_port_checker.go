package checker

import (
	"errors"
	"log"
	"time"

	goTezos "github.com/goat-systems/go-tezos/v4/rpc"
)

const (
	localRpc    = "http://localhost:8732"
	nodeChecker = "NodeChecker"
)

type nodeUrlChecker struct {
	localRpcCli *goTezos.Client
}

func (c *nodeUrlChecker) AssertRunning() (err error) {
	_, head, err := c.localRpcCli.Block(&goTezos.BlockIDHead{})
	if err != nil {
		return err
	}
	log.Printf("Current head timestamp %v", head.Header.Timestamp)
	if time.Now().Sub(head.Header.Timestamp).Minutes() > 10.0 {
		log.Println("node is unsync")
		return errors.New("node is unsync")
	}
	return nil
}

func (c *nodeUrlChecker) GetTitle() string {
	return nodeChecker
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
