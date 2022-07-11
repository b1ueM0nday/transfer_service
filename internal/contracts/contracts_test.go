package contracts

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"testing"
)

var (
	blockchain *backends.SimulatedBackend
)

func TestPrepare_EmptyConfig(t *testing.T) {
	badConfig := Config{
		IP:             "",
		HttpPort:       "",
		WsPort:         "",
		AddressPath:    "",
		PrivateKeyPath: "",
	}
	c := NewClient(context.Background())
	err := c.Prepare(&badConfig)
	if err == nil {
		t.Errorf("bad config didn't trigger an error")
	}
}
