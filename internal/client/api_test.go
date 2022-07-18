package client

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func genKeyString() string {
	receiverKey, _ := crypto.GenerateKey()
	publicKeyECDSA, _ := receiverKey.Public().(*ecdsa.PublicKey)

	return crypto.PubkeyToAddress(*publicKeyECDSA).String()
}

func TestClient_Deposit(t *testing.T) {
	tests := []struct {
		name    string
		amount  *big.Int
		wantErr bool
	}{
		{
			name:    "negative amount",
			amount:  big.NewInt(-100),
			wantErr: true,
		},
		{
			name:    "nil amount returns error, error expected",
			amount:  nil,
			wantErr: true,
		},
		{
			name:    "valid amount returns nil, nil expected",
			amount:  big.NewInt(100500),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeCli, _ := CreateFakeContract()
			options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.cli.owner.pk, fakeCli.cli.chainId)

			tx, err := fakeCli.cli.deposit(tt.amount, options)
			fakeCli.blockchain.Commit()
			if (tx == nil || err != nil) != tt.wantErr {
				t.Fatalf("deposit: %s \n expected error = %t", tt.name, tt.wantErr)
			}

		})
	}
}

func TestClient_Withdraw(t *testing.T) {
	tests := []struct {
		name    string
		amount  *big.Int
		wantErr bool
	}{
		{
			name:    "negative amount",
			amount:  big.NewInt(-100),
			wantErr: true,
		},
		{
			name:    "nil amount",
			amount:  nil,
			wantErr: true,
		},
		{
			name:    "valid amount",
			amount:  big.NewInt(0),
			wantErr: false,
		},
		{
			name:    "insufficient funds",
			amount:  big.NewInt(100),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeCli, _ := CreateFakeContract()
			options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.cli.owner.pk, fakeCli.cli.chainId)

			tx, err := fakeCli.cli.withdraw(tt.amount, options)
			fakeCli.blockchain.Commit()
			if (tx == nil || err != nil) != tt.wantErr {
				t.Fatalf("withdraw: %s \n expected error = %t", tt.name, tt.wantErr)
			}

		})
	}
}

func TestClient_Transfer(t *testing.T) {
	tests := []struct {
		name     string
		amount   *big.Int
		receiver string
		wantErr  bool
	}{
		{
			name:     "negative amount",
			amount:   big.NewInt(-100),
			receiver: genKeyString(),
			wantErr:  true,
		},
		{
			name:     "nil amount",
			amount:   nil,
			receiver: genKeyString(),
			wantErr:  true,
		},
		{
			name:     "valid amount",
			amount:   big.NewInt(0),
			receiver: genKeyString(),
			wantErr:  false,
		},
		{
			name:     "insufficient funds",
			amount:   big.NewInt(100),
			receiver: genKeyString(),
			wantErr:  true,
		},

		{
			name:     "empty receiver field",
			amount:   big.NewInt(100),
			receiver: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeCli, _ := CreateFakeContract()
			options, _ := bind.NewKeyedTransactorWithChainID(fakeCli.cli.owner.pk, fakeCli.cli.chainId)

			tx, err := fakeCli.cli.transfer(tt.receiver, tt.amount, options)
			fakeCli.blockchain.Commit()
			if (tx == nil || err != nil) != tt.wantErr {
				t.Fatalf("transfer: %s \n error was expected  = %t", tt.name, tt.wantErr)
			}

		})
	}
}
func TestClient_GetBalance(t *testing.T) {
	tests := []struct {
		name            string
		balanceExpected *big.Int
		hasReceiver     bool
		wantErr         bool
	}{
		{
			name:            "nil receiver",
			balanceExpected: big.NewInt(0),
			hasReceiver:     false,
			wantErr:         false,
		},

		{
			name:            "empty wallet with receiver",
			balanceExpected: big.NewInt(0),
			hasReceiver:     true,
			wantErr:         false,
		},
		{
			name:            "empty wallet w/o receiver",
			balanceExpected: big.NewInt(0),
			hasReceiver:     true,
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeCli, _ := CreateFakeContract()
			var b *big.Int
			var err error

			if tt.hasReceiver {
				receiver := genKeyString()
				b, err = fakeCli.cli.GetBalance(&receiver)
			} else {
				b, err = fakeCli.cli.GetBalance(nil)
			}

			if (b == tt.balanceExpected || err != nil) != tt.wantErr {
				t.Fatalf("balance: %s \n expected error = %t", tt.name, tt.wantErr)
			}

		})
	}
}
