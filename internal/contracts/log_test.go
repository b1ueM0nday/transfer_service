package contracts

import (
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

func TestRun_EmptyAddress(t *testing.T) {
	c, _ := CreateFakeContract()
	l := NewLogger(nil, make(chan *types.Transaction))

	err := l.Run("", c.contractAddr)
	if err == nil {
		t.Fatalf("run with empty ws address")
	}
}
