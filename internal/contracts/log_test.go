package contracts

import (
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

func TestRun(t *testing.T) {
	c, _ := CreateFakeContract()
	l := NewLogger(nil, make(chan *types.Transaction))

	err := l.Run("bad address", c.contractAddr)
	if err == nil {
		t.Fatalf("run with bad ws address")
	}
}
