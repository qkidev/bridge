package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/encoding/gjson"
	"recheck-go/store"
	"strconv"

	//"github.com/gogf/gf/encoding/gjson"

	//"github.com/gogf/gf/container/gmap"

	//geth "github.com/ethereum/go-ethereum/mobile"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"math/big"
)

func main() {
	ctx := context.Background()
	client, _ := ethclient.Dial("https://hz.node.quarkblockchain.cn")
	privateKey, _ := crypto.HexToECDSA("")
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	//g.Dump(publicKeyECDSA)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//g.Dump(fromAddress.String())
	nonce, _ := client.PendingNonceAt(ctx, fromAddress)
	g.Dump("nonce: ", nonce)
	gasPrice, _ := client.SuggestGasPrice(ctx)
	g.Dump("gasPrice: ", gasPrice)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(20181205))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(20)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	address := common.HexToAddress("0x5996aEEBCb45180619ab63f0E537f766954D484a")
	instanceBridgeManager, _ := store.NewBridgeManager(address, client)
	signLimit, _ := instanceBridgeManager.BridgeAddress(nil)
	g.Dump(signLimit)
}

func demo() {

	ctx := context.Background()

	chains, _ := g.DB().Model("chain").Where("chainId", 20181205).All()
	for _, v := range chains {
		chain := v.GMap()
		url := gconv.String(chain.Get("url"))
		//bridgeAddress := gconv.String(chain.Get("bridge"))

		//conn,_:= rpc.Dial(url)

		//ctx := geth.NewContext()
		//conn2,_ := geth.NewEthereumClient(url)
		//
		//
		//block2, _ := conn2.GetBlockByNumber(ctx, 7550431)
		//g.Dump(block2.GetHash())

		conn, _ := ethclient.Dial(url)

		//header, _ := conn.HeaderByNumber(context.Background(),big.NewInt(7543998))
		//g.Dump(header)

		block, _ := conn.BlockByNumber(ctx, big.NewInt(7568068))

		for _, tx := range block.Transactions() {
			if tx.To().String() == chain.Get("bridge") {
				txStr, _ := tx.MarshalJSON()
				txJson := gjson.New(txStr)
				input := txJson.GetString("input")
				//g.Dump(input)
				methodId := input[0:10]
				//g.Dump(methodId)
				if methodId == "0xbc157ac1" {
					// deposit
					chainId, _ := strconv.ParseInt(input[10:74], 16, 32)
					toToken := input[74:138]
					value, _ := strconv.ParseInt(input[138:202], 16, 32)
					g.Dump("chainId: ", chainId, "toToken: ", toToken, "\nvalue: ", value)
				}

				if methodId == "0x02367ad2" {
					// deposit native

				}

				//g.Dump(tx)
				hash := tx.Hash().String()
				//g.Dump(hash)
				existData, _ := g.DB().Model("log").Where("depositHash", hash).One()
				if existData != nil {
					withdrawHash := existData.GMap().Get("withdrawHash")
					if withdrawHash == nil {

					}
				} else {
					//receipt, _ := conn.TransactionReceipt(ctx, tx.Hash())
					//g.Dump(receipt)
					//fromChain, _ := conn.ChainID(ctx)
					//g.Dump(fromChain)
					//_, _ = g.DB().Model("log").Insert(g.Map{
					//	"depositHash": hash,
					//	"fromChain":   fromChain,
					//})
				}
				//g.Dump(existData)

			}

			//g.Dump(receipt.Logs[0].Topics[0].String())
			//receiptJson:=gjson.New(receipt.Logs,false)
			//g.Dump(receiptJson.GetArray("logs.0"))
			//g.Dump(tx)
			//sender, _ :=conn.TransactionSender(context.Background(),tx,tx.Hash(), uint(txIndex))
			//g.Dump(sender)
			//g.Dump(tx.ChainId())
			//g.Dump(tx.Hash().String())
			//detail,isPending,_:=conn.TransactionByHash(context.Background(),tx.Hash())
			//g.Dump(isPending)
		}
		//num,_:=conn.BlockNumber(context.Background())
		//g.Dump(num)

		//bridgeToken,_:= abi.NewBridge(common.HexToAddress(bridgeAddress),conn)
		//g.Dump(bridgeToken)
	}
}

// 0xbc157ac1
// 0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000
// 6af164364182b96385c588232c84a576d875573b
// 0000000000000000000000000000000000000000000000000000000000
// 0x895440
