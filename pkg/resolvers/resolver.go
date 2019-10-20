package resolvers
//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	datasources "github.com/proxima-one/binance-chain-subgraph/pkg/datasources"
	models "github.com/proxima-one/binance-chain-subgraph/pkg/models"
	gql "github.com/proxima-one/binance-chain-subgraph/pkg/gql"
	proxima "github.com/proxima-one/proxima-db-client-go"
	"encoding/json"
	_ "fmt"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
type Resolver struct{
	db *proxima.ProximaDB
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func NewResolver(db *proxima.ProximaDB) (gql.Config) {
	r := Resolver{}
	r.db = db
	return gql.Config{
		Resolvers: &r,
	}
}
//TODO cached no expiration
func (r *queryResolver) BlockStats(ctx context.Context, prove *bool) (*models.ProximaBlockStats, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove } else { inputs["prove"] = nil}
	args := GenerateInputs(inputs, BlockStatsDefaultInputs)
	resp, _ := r.db.Get(datasources.Primary, "BlockStats", args)
	result := *resp
	value := models.BlockStats{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaBlockStats{BlockStats: &value, Proof: &proof}, nil
}
//TODO cached no expiration
func (r *queryResolver) Fees(rctx context.Context, prove *bool) (*models.ProximaFees, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	args := GenerateInputs(inputs, FeesDefaultInputs)
	resp, _ := r.db.Get(datasources.Primary, "Fees", args)
	result := *resp
	value := []*models.Fee{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaFees{Fees: value, Proof: &proof}, nil
}
//TODO cached no expiration
func (r *queryResolver) Tokens(rctx context.Context, limit *int, offset *int, prove *bool) (*models.ProximaTokens, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (limit != nil ) { inputs["limit"] = *limit }
	if (offset != nil ) { inputs["offset"] = *offset }
	args := GenerateInputs(inputs, TokensDefaultInputs)
	resp, _ := r.db.Get(datasources.Primary, "Tokens", map[string]interface{}(args))
	result := *resp
	value := []*models.Token{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaTokens{Tokens: value, Proof: &proof}, nil
}

func (r *queryResolver) Account(rctx context.Context, address *string, prove *bool) (*models.ProximaAccount, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (address != nil ) { inputs["address"] = *address }
	args := GenerateInputs(inputs, AccountDefaultInputs)
	result, _ := r.db.Get(datasources.AccountsByAddress, *address, map[string]interface{}(args))
	value := models.Account{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaAccount{Account: &value, Proof: &proof}, nil
}
//fetch from network first
func (r *queryResolver) Orders(rctx context.Context, address *string, symbol *string, start *string, end *string, orderSide *int, open *bool, status *string, total *int, limit *int, offset *int, prove *bool) ([]*models.ProximaOrder, error) {
	// args :=  map[string]interface{"prove":prove,}
	// orders, _ := r.db.Get(datasources.OrdersByOwner, address, (map[string]interface{}(args))
	//
	// //symbol
	// //start
	// //endtime
	// //status
	//
	// for _, order := range orders {
	// 		//filter the orders based on symbol, start and endtime, open, status
	// 		if filter_fn(order, conditions) {
	// 			//append to the correct list
	// 		}
	// }
	// //return the proof
	// return &orderList, nil
	return nil, nil
}
//fetch from network first
func (r *queryResolver) Order(rctx context.Context, orderID *string, prove *bool) (*models.ProximaOrder, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (orderID != nil ) { inputs["orderId"] = *orderID }
	args := GenerateInputs(inputs, OrderDefaultInputs)
	result, _ := r.db.Get(datasources.OrdersByOrderId, *orderID, map[string]interface{}(args))
	// if (result == nil) {
	// 	datasource.updateOrdersByOrderId(orderIds map[string]bool)
	// }
	value := models.Order{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaOrder{Order: &value, Proof: &proof}, nil
}
//cached by block?
func (r *queryResolver) Transactions(rctx context.Context, address *string, txType *string, txAsset *string, txSide *int, blockHeight *string, startTime *string, endTime *string, limit *int, offset *int, prove *bool) ([]*models.ProximaTransaction, error) {
	// args :=  map[string]interface{"prove":prove,}
	// transactions, err := "", nil //need to get the transactions which one...
	// if address != undefined {
	// 	if txSide == 1 {
	// 		transactions, err = r.db.Get(datasources.TransactionsByFromAddr, address, (map[string]interface{}(args))
	// 	} else {
	// 		transactions, err = r.db.Get(datasources.TransactionsByToAddr, address, (map[string]interface{}(args))
	// 	}
	// } else if blockHeight != undefined {
	// 	transactions, err = r.db.Get(datasources.TransactionsByBlockHeight, blockHeight, (map[string]interface{}(args))
	// } else if {
	// 	queryString:= "time" //TODO need to create an accurate query string
	// 	transactions, err = r.db.Query(datasources.TransactionsByTimeStamp, queryString, (map[string]interface{}(args))
	// } else {
	// 	return &transactions, err
	// }
	// //filter should do this...
	// transaction["transaction"]["toAddr"] == address //"fromAddr"
	// transaction["transaction"]["txAsset"] == txAsset
	// transaction["transaction"]["txType"] == txType
	// transaction["transaction"]["timeStamp"] >= startTime
	// transaction["transaction"]["timeStamp"] <= endTime
	// transaction["transaction"]["blockHeight"] == blockHeight
	//
	// for _, transaction := range transactions {
	// 	if  filter_fn(transaction, conditions) {
	// 		//append to another array of transactionsList
	// 	}
	// }
	// //unmarshall to the proper struct???
	// if prove {
	// 	for _, transaction := range transactionsList {
	// 		proximaTx,_ := r.db.Get(datasources.TransactionsByTxHash, transaction["txHash"], (map[string]interface{}(args))
	// 	}
	// } else {
	// 	proximaTransactions := &transactionsList
	// }

	return nil, nil
}
//might have to fetch
func (r *queryResolver) Transaction(rctx context.Context, txHash *string, prove *bool) (*models.ProximaTransaction, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (txHash != nil ) { inputs["txHash"] = *txHash}
	args := GenerateInputs(inputs, TransactionDefaultInputs)

	result, _ := r.db.Get(datasources.TransactionsByTxHash, txHash, map[string]interface{}(args))

	value := models.Transaction{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaTransaction{Transaction: &value, Proof: &proof}, nil
}
//cached no expiration
func (r *queryResolver) Markets(rctx context.Context, limit *int, offset *int, prove *bool) (*models.ProximaMarkets, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (limit != nil ) { inputs["limit"] = *limit }
	if (offset != nil ) { inputs["offset"] = *offset }
	args := GenerateInputs(inputs, MarketsDefaultInputs)

	result, _ := r.db.Get(datasources.Primary, "Markets", map[string]interface{}(args))

	value := []*models.Market{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaMarkets{Markets: value, Proof: &proof}, nil
}
//cached expiration, 2 seconds?
func (r *queryResolver) MarketTicker(rctx context.Context, symbol *string, prove *bool) (*models.ProximaMarketTicker, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (symbol != nil ) { inputs["symbol"] = *symbol }
	args := GenerateInputs(inputs, MarketTickerDefaultInputs)

	result, _ := r.db.Get(datasources.MarketTickersBySymbol, args["symbol"], map[string]interface{}(args))

	value := models.MarketTicker{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaMarketTicker{MarketTicker: &value, Proof: &proof}, nil
}
//cached expiration, 2 seconds?
func (r *queryResolver) MarketTickers(rctx context.Context, limit *int, offset *int, prove *bool) (*models.ProximaMarketTickers, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (limit != nil ) { inputs["limit"] = *limit }
	if (offset != nil ) { inputs["offset"] = *offset }
	args := GenerateInputs(inputs, MarketTickersDefaultInputs)

	result, _ := r.db.Get(datasources.Primary, "MarketTickers", args)

	value := []*models.MarketTicker{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaMarketTickers{MarketTickers: value, Proof: &proof}, nil
}
//cached expiration, 2 seconds?
func (r *queryResolver) MarketDepth(rctx context.Context, symbolPair *string, limit *int, prove *bool) (*models.ProximaMarketDepth, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (symbolPair != nil ) { inputs["symbol_pair"] = *symbolPair }
	if (limit != nil ) { inputs["limit"] = *limit }
	args := GenerateInputs(inputs, MarketDepthDefaultInputs)

	result, _ := r.db.Get(datasources.MarketDepthBySymbol, args["symbol_pair"], map[string]interface{}(args))
	
	value := models.MarketDepth{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaMarketDepth{MarketDepth: &value, Proof: &proof}, nil
}
//cached expiration, 2 seconds? most common intervals
func (r *queryResolver) MarketCandleSticks(rctx context.Context, symbol *string, startTime *string, endTime *string, interval *string, limit *int, prove *bool) (*models.ProximaMarketCandleSticks, error) {
	// proximaMarketCandleSticks := map[string]interface{}
	// proximaMarketCandleSticks["proof"] = map[string]string{"proof": "", "root": "", }
	// requestArgs := map[string]interface{}{
	// 	"symbol" : symbol,
	// 	"interval": interval,
	// }
	// marketCandleSticks, _ := datasources.dataRequest("marketCandleSticks", map[string]interface{}(requestArgs))
	// proximaMarketCandleSticks["market_candlesticks"] = &marketCandleSticks
	// proximaMarketCandleSticks["symbol"] = symbol
	// proximaMarketCandleSticks["interval"]= interval
	// return &proximaMarketCandleSticks, nil
	return nil, nil
}
//cached by block?
func (r *queryResolver) Trades(rctx context.Context, address *string, symbol *string, quoteAssetSymbol *string, blockHeight *string, startTime *string, endTime *string, buyerOrderID *string, sellerOrderID *string, orderSide *int, limit *int, offset *int, prove *bool) ([]*models.ProximaTrade, error) {
	// args :=  map[string]interface{"prove":prove,}
	// trades, err := "", nil
	//
	//
	// //filter the transactions
	// //if address, then get from address
	// 	//filter other vars
	// //if blockHeight
	// //confirm blocks...
	// //if proof then get the transaction again
	//
	// for _, trade := range trades {
	// 	if filter_fn(trade, conditions) {
	// 		//add to new slice
	// 	}
	//
	// }
	//
	// //make new array of proximaTrades ...
	// for _, trade := range tradeList {
	// 	if prove {
	// 		//have to lookup
	// 	} else {
	// 		//do not have to look up again
	// 	}
	// }
	return nil, nil
}
//might have to fetch from extras
func (r *queryResolver) Trade(rctx context.Context, tradeID *string, prove *bool) (*models.ProximaTrade, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (tradeID != nil ) { inputs["tradeId"] = *tradeID }
	args := GenerateInputs(inputs, TradeDefaultInputs)

	result, _ := r.db.Get(datasources.TradesByTradeId, *tradeID, map[string]interface{}(args))
	value := models.Trade{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaTrade{Trade: &value, Proof: &proof}, nil
}
//might have to fetch from the last block
func (r *queryResolver) AtomicSwaps(rctx context.Context, fromAddress *string, toAddress *string, startTime *string, endTime *string, limit *int, offset *int, prove *bool) ([]*models.ProximaAtomicSwap, error) {
	// panic("not implemented")
	// //list of proximaatomicswaps
	// atomicSwaps := "" //need to get the transactions
	// for _, atomicSwap := range atomicSwaps {
	// 	//filter the transactions
	// 	//if address, then get from address
	// 		//filter other vars
	// 	//if blockHeight
	// 	//confirm blocks...
	// 	//if proof then get the transaction again
	// }
	//
	return nil, nil
}
//might have to fetch
func (r *queryResolver) AtomicSwap(rctx context.Context, id *string, prove *bool) (*models.ProximaAtomicSwap, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	if (id != nil ) { inputs["id"] = *id }
	args := GenerateInputs(inputs, AtomicSwapDefaultInputs)
	result, _ := r.db.Get(datasources.AtomicSwapsBySwapId, *id, map[string]interface{}(args))
	value := models.AtomicSwap{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaAtomicSwap{AtomicSwap: &value, Proof: &proof}, nil
}
//cached expiration, 2 seconds?
func (r *queryResolver) Validators(rctx context.Context, prove *bool) (*models.ProximaValidators, error) {
	inputs :=  make(map[string]interface{});
	if (prove != nil ) { inputs["prove"] = *prove }
	args := GenerateInputs(inputs, ValidatorsDefaultInputs)
	result, _ := r.db.Get(datasources.Primary, "Validators", map[string]interface{}(args))
	value := []*models.Validator{}
	json.Unmarshal(result.GetValue(), &value)
	proof := GenerateProof(result.GetProof())
	return &models.ProximaValidators{Validators: value, Proof: &proof}, nil
}
//fetch
func (r *queryResolver) Timelocks(rctx context.Context, address *string, id *int, prove *bool) (*models.ProximaTimelocks, error) {
	// proximaTimelocks := map[string]interface{}{
	// 	"proof": map[string]string{
	// 			"proof": "",
	// 			"root": "",
	// 	},
	// 	}
	// args := map[string]interface{}{
	// 	"address":address,
	// 	"id":id,
	// }
//	timelocks, _ := datasources.dataRequest("timelocks", map[string]interface{}(args))
//	proximaTimelocks["timelocks"] = timelocks
	return nil, nil
}
