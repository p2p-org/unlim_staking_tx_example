package main

import (
	"context"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	directstakingv3 "github.com/p2p-org/unlim_staking_tx_example/contract"
	"log"
	"math/big"
	"strings"
)

const (
	ZeroAddress                = "0x0000000000000000000000000000000000000000"
	wcEmptyBytes               = "0000000000000000000000"
	WCTypePartiallyWithdrawing = "0x01"
	WCTypeAutoCompounding      = "0x02"
	// Change to address corresponding to your private key you use in this code
	clientAddrStr = "0x1e5988c1A678EA49f17C72244Abd8c7D41945427"
)

var clientAddress = common.HexToAddress(clientAddrStr)

func main() {
	rpc, err := ethclient.Dial("http://141.253.108.252:8545")
	if err != nil {
		log.Fatal(err)
	}
	contract, err := directstakingv3.NewDirectStakingDepositorV3Transactor(common.HexToAddress("0xd88eA323b51513cdD4bD581c872903dC8c7Ef8fe"), rpc)
	if err != nil {
		log.Fatal(err)
	}
	pk, err := crypto.HexToECDSA("a577c6b89c9cfea82ebedc485e388571c9002cc8aab41c9e9616775561e293d2") // Insert your private key here in hex format without 0x prefix
	if err != nil {
		log.Fatal(err)
	}
	signer := func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(transaction, types.LatestSignerForChainID(big.NewInt(7078815900)), pk)
	}
	nonce, err := rpc.PendingNonceAt(context.Background(), clientAddress)
	if err != nil {
		log.Fatal(err)
	}
	var withdrawalCreds [32]byte
	copy(withdrawalCreds[:], common.FromHex(WithdrawalCredentialsFromAddress(clientAddrStr, WCTypePartiallyWithdrawing)))
	txOpts := &bind.TransactOpts{
		From:      clientAddress,
		Value:     WeiFromETHString("32"),
		Nonce:     big.NewInt(int64(nonce)),
		Signer:    signer,
		GasFeeCap: GweiToWei(100),
		GasTipCap: big.NewInt(1_000_000_000),
		GasLimit:  30000000,
	}
	tx, err := contract.AddEth(
		txOpts,
		withdrawalCreds,
		GweiToWei(uint64(EthToGwei(32))),
		common.HexToAddress("0x6c882B6bffe33e9D1abBE0afAa6b070F4308EB66"),
		directstakingv3.FeeRecipient{
			BasisPoints: big.NewInt(9500),
			Recipient:   clientAddress,
		}, directstakingv3.FeeRecipient{
			BasisPoints: big.NewInt(0),
			Recipient:   common.HexToAddress(ZeroAddress),
		},
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tx sent: %s", tx.Hash().Hex())
}

func WithdrawalCredentialsFromAddress(address string, wcType string) string {
	address = strings.TrimPrefix(address, "0x")
	return strings.Join([]string{wcType, wcEmptyBytes, address}, "")
}

func EthToGwei(eth float64) spec.Gwei {
	return spec.Gwei(eth * float64(1_000_000_000))
}

func GweiToEth(gwei spec.Gwei) float64 {
	return float64(gwei) / 1_000_000_000
}

func GweiToWei(gwei uint64) *big.Int {
	return big.NewInt(0).Mul(big.NewInt(int64(gwei)), big.NewInt(1_000_000_000))
}

func WeiToGwei(wei *big.Int) uint64 {
	return wei.Div(wei, big.NewInt(1_000_000_000)).Uint64()
}

func WeiFromETHString(eth string) *big.Int {
	amount, _ := new(big.Int).SetString(eth, 10)
	return new(big.Int).Mul(amount, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}
