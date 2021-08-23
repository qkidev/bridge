package main

import (
	"background/token"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/util/gconv"
	"math"
)

func bridgeSafe() {
	pairs, _ := g.Model("pair").All()
	for _, pair := range pairs {
		isMain := pair.GMap().Get("isMain")
		if gconv.Int(isMain) == 1 {
			decimal := gconv.Int(pair.GMap().Get("decimal"))
			isNative := pair.GMap().Get("isNative")
			fromChain, _ := g.Model("chain").One("chainId", pair.GMap().Get("fromChain"))
			fromProvider, _ := ethclient.Dial(gconv.String(fromChain.GMap().Get("url")))
			var fromBalance uint64
			if gconv.Int(isNative) == 0 {
				fromToken, _ := token.NewToken(common.HexToAddress(gconv.String(pair.GMap().Get("fromToken"))), fromProvider)
				fromBalanceBig, _ := fromToken.BalanceOf(nil, common.HexToAddress(gconv.String(fromChain.GMap().Get("bridge"))))
				fromBalance = fromBalanceBig.Uint64()
			} else {
				fromBalanceBig, _ := fromProvider.BalanceAt(context.Background(), common.HexToAddress(gconv.String(fromChain.GMap().Get("bridge"))), nil)
				fromBalance = fromBalanceBig.Uint64()
			}
			toChain, _ := g.Model("chain").One("chainId", pair.GMap().Get("toChain"))
			toProvider, _ := ethclient.Dial(gconv.String(toChain.GMap().Get("url")))
			toToken, _ := token.NewToken(common.HexToAddress(gconv.String(pair.GMap().Get("toToken"))), toProvider)
			toBalance, _ := toToken.TotalSupply(nil)
			errNumBig := fromBalance - toBalance.Uint64()
			g.Model("amount_error").Save(g.Map{
				"mainChain":   fromChain.GMap().Get("name"),
				"mainBalance": formatBalance(float64(fromBalance), decimal),
				"toBalance":   formatBalance(float64(toBalance.Uint64()), decimal),
				"toChain":     toChain.GMap().Get("name"),
				"tokenName":   pair.GMap().Get("name"),
				"errNum":      formatBalance(gconv.Float64(errNumBig), decimal),
			})
			fmt.Println(fromChain.GMap().Get("name"), toChain.GMap().Get("name"), pair.GMap().Get("name"), formatBalance(gconv.Float64(errNumBig), decimal))
		}
	}
}

func formatBalance(num float64, decimal int) string {
	return gconv.String(num / gconv.Float64(math.Pow10(decimal)))
}

func setPairStop() {
	errors, _ := g.Model("amount_error").All()
	for _, record := range errors {
		limit := gconv.Float64(record.GMap().Get("errLimit"))
		if limit != 0 {
			errNum := gconv.Float64(record.GMap().Get("errNum"))
			if math.Abs(errNum) > limit {
				g.Model("pair").Update(g.Map{
					"isStop": 1,
				}, "name", record.GMap().Get("tokenName"))
			}
		}
	}
}

func main() {
	bridgeSafe()
	gcron.Add("0 * * * * *", func() {
		bridgeSafe()
		setPairStop()
	}, "bridge safe")
	select {}
}
