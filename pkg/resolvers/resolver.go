package resolvers
//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	//datasources "github.com/proxima-one/binance-chain-subgraph/pkg/datasources"
	models "github.com/proxima-one/binance-chain-subgraph/pkg/models"
	gql "github.com/proxima-one/binance-chain-subgraph/pkg/gql"
	dataloader "github.com/proxima-one/binance-chain-subgraph/pkg/dataloader"
	//json "github.com/json-iterator/go"
	_ "fmt"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
type Resolver struct{
	loader *dataloader.Dataloader
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func NewResolver(loader *dataloader.Dataloader) (gql.Config) {
	r := Resolver{}
	r.loader = loader
	return gql.Config{
		Resolvers: &r,
	}
}

func (r *queryResolver) BlockStats(ctx context.Context, prove *bool) (*models.ProximaBlockStats, error) {
	args := BlockStatsDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	return r.loader.LoadProximaBlockStats(args)
}

func (r *queryResolver) Fees(rctx context.Context, prove *bool) (*models.ProximaFees, error) {
	args :=  FeesDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove}
	return r.loader.LoadProximaFees(args)
}

func (r *queryResolver) Tokens(rctx context.Context, limit *int, offset *int, prove *bool) (*models.ProximaTokens, error) {
	args :=  TokensDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (limit != nil ) { args["limit"] = *limit }
	if (offset != nil ) { args["offset"] = *offset }
	return r.loader.LoadProximaTokens(args)
}

func (r *queryResolver) Account(rctx context.Context, address *string, prove *bool) (*models.ProximaAccount, error) {
	args :=  AccountDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (address != nil ) { args["address"] = *address }
	return r.loader.LoadProximaAccount(args)
}

func (r *queryResolver) Orders(rctx context.Context, address *string, symbol *string, start *string, end *string, orderSide *int, open *bool, status *string, total *int, limit *int, offset *int, prove *bool) ([]*models.ProximaOrder, error) {
	args :=  OrdersDefaultInputs
	if (prove != nil ) { args["prove"] = *prove }
	if (symbol != nil ) { args["symbol"] = *symbol }
	if (limit != nil ) { args["limit"] = *limit }
	if (offset != nil ) { args["offset"] = *offset }
	if (address != nil ) { args["address"] = *address }
	if (open != nil ) { args["open"] = *open }
	if (start != nil ) { args["start"] = *start }
	if (end != nil ) { args["end"] = *end }
	if (total != nil ) { args["total"] = *total }
	if (orderSide != nil ) { args["orderSide"] = *orderSide }
	return r.loader.LoadProximaOrders(args)
}
//fetch from network first
func (r *queryResolver) Order(rctx context.Context, orderID *string, prove *bool) (*models.ProximaOrder, error) {
	args :=  OrderDefaultInputs
	if (prove != nil ) { args["prove"] = *prove }
	if (orderID != nil ) { args["orderId"] = *orderID }
	return r.loader.LoadProximaOrder(args)
}

func (r *queryResolver) Transactions(rctx context.Context, address *string, txType *string, txAsset *string, txSide *int, blockHeight *string, startTime *string, endTime *string, limit *int, offset *int, prove *bool) ([]*models.ProximaTransaction, error) {
	args :=  TransactionsDefaultInputs
	if (txType != nil ) { args["txType"] = *txType }
	if (txAsset != nil ) { args["txAsset"] = *txAsset }
	if (txSide != nil ) { args["txSide"] = *txSide }
	if (blockHeight != nil ) { args["blockHeight"] = *blockHeight }
	if (prove != nil ) { args["prove"] = *prove }
	if (limit != nil ) { args["limit"] = *limit }
	if (offset != nil ) { args["offset"] = *offset }
	if (address != nil ) { args["address"] = *address }
	if (endTime != nil ) { args["endTime"] = *endTime }
	return r.loader.LoadProximaTransactions(args)
}

func (r *queryResolver) Transaction(rctx context.Context, txHash *string, prove *bool) (*models.ProximaTransaction, error) {
	args:=  TransactionDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (txHash != nil ) { args["txHash"] = *txHash}
	return r.loader.LoadProximaTransaction(args)
}

func (r *queryResolver) Markets(rctx context.Context, limit *int, offset *int, prove *bool) (*models.ProximaMarkets, error) {
	args := MarketsDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (limit != nil ) { args["limit"] = *limit }
	if (offset != nil ) { args["offset"] = *offset }
	return r.loader.LoadProximaMarkets(args)
}

func (r *queryResolver) MarketTicker(rctx context.Context, symbol *string, prove *bool) (*models.ProximaMarketTicker, error) {
	args :=  MarketTickerDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (symbol != nil ) { args["symbol"] = *symbol }
	return r.loader.LoadProximaMarketTicker(args)
}

func (r *queryResolver) MarketTickers(rctx context.Context, limit *int, offset *int, prove *bool) (*models.ProximaMarketTickers, error) {
	args :=  MarketTickersDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (limit != nil ) { args["limit"] = *limit }
	if (offset != nil ) { args["offset"] = *offset }
	return r.loader.LoadProximaMarketTickers(args)
}

func (r *queryResolver) MarketDepth(rctx context.Context, symbolPair *string, limit *int, prove *bool) (*models.ProximaMarketDepth, error) {
	args :=  MarketTickerDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (symbolPair != nil ) { args["symbol_pair"] = *symbolPair }
	if (limit != nil ) { args["limit"] = *limit }
	return r.loader.LoadProximaMarketDepth(args)
}

func (r *queryResolver) MarketCandleSticks(rctx context.Context, symbol *string, startTime *string, endTime *string, interval *string, limit *int, prove *bool) (*models.ProximaMarketCandleSticks, error) {
	args:= MarketCandleSticksDefaultInputs
	if (prove != nil ) { args["prove"] = *prove }
	if (limit != nil ) { args["limit"] = *limit }
	if (interval != nil ) { args["interval"] = *interval }
	if (startTime != nil ) { args["start"] = *startTime }
	if (endTime != nil ) { args["endTime"] = *endTime }
	if (symbol != nil ) { args["symbol"] = *symbol }
	return r.loader.LoadProximaMarketCandleSticks(args)
}

func (r *queryResolver) Trades(rctx context.Context, address *string, symbol *string, quoteAssetSymbol *string, blockHeight *string, startTime *string, endTime *string, buyerOrderID *string, sellerOrderID *string, orderSide *int, limit *int, offset *int, prove *bool) ([]*models.ProximaTrade, error) {
	args :=  TradesDefaultInputs
 	if (address != nil ) { args["address"] = *address }
	if (symbol != nil ) { args["symbol"] = *symbol }
	if (quoteAssetSymbol != nil ) { args["quoteAssetSymbol"] = *quoteAssetSymbol }
	if (blockHeight != nil ) { args["blockHeight"] = *blockHeight }
	if (startTime != nil ) { args["start"] = *startTime }
	if (endTime != nil ) { args["endTime"] = *endTime }
	if (buyerOrderID != nil ) { args["buyerOrderID"] = *buyerOrderID }
	if (sellerOrderID != nil ) { args["sellerOrderID"] = *sellerOrderID }
	if (orderSide != nil ) { args["orderSide"] = *orderSide }
	if (limit != nil ) { args["limit"] = *limit }
	if (offset != nil ) { args["offset"] = *offset }
	if (prove != nil ) { args["prove"] = *prove }
	return r.loader.LoadProximaTrades(args)
}

func (r *queryResolver) AtomicSwaps(rctx context.Context, fromAddress *string, toAddress *string, startTime *string, endTime *string, limit *int, offset *int, prove *bool) ([]*models.ProximaAtomicSwap, error) {
	args:= AtomicSwapsDefaultInputs
	if (prove != nil ) { args["prove"] = *prove }
	if (limit != nil ) { args["limit"] = *limit }
	if (offset != nil ) { args["offset"] = *offset }
	if (startTime != nil ) { args["startTime"] = *startTime }
	if (endTime != nil ) { args["endTime"] = *endTime }
	if (fromAddress != nil ) { args["fromAddress"] = *fromAddress }
	if (toAddress != nil ) { args["toAddress"] = *toAddress }
	return r.loader.LoadProximaAtomicSwaps(args)
}

func (r *queryResolver) AtomicSwap(rctx context.Context, id *string, prove *bool) (*models.ProximaAtomicSwap, error) {
	args :=  AtomicSwapDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	if (id != nil ) { args["id"] = *id }
	return r.loader.LoadProximaAtomicSwap(args)
}

func (r *queryResolver) Validators(rctx context.Context, prove *bool) (*models.ProximaValidators, error) {
	args :=  ValidatorsDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	return r.loader.LoadProximaValidators(args)
}

func (r *queryResolver) Timelocks(rctx context.Context, address *string, prove *bool) (*models.ProximaTimelocks, error) {
	args := TimelocksDefaultInputs
	if (prove != nil ) { args["prove"] = *prove }
	if (address != nil ) { args["address"] = *address }
	return r.loader.LoadProximaTimelocks(args)
}
