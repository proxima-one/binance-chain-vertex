package dataloader

import (
  cache "github.com/patrickmn/go-cache"
  datasources "github.com/proxima-one/binance-chain-subgraph/pkg/datasources"
  models "github.com/proxima-one/binance-chain-subgraph/pkg/models"
  json "github.com/json-iterator/go"
  proxima "github.com/proxima-one/proxima-db-client-go"

  //"time"
)

//gets added context
type Dataloader struct{
	db *proxima.ProximaDB
  cache *cache.Cache
  datasource *datasources.Datasource
}

func NewDataloader(c *cache.Cache, db *proxima.ProximaDB, datasource *datasources.Datasource) (*Dataloader, error) {
	return &Dataloader{
		db: db,
    cache: c,
    datasource: datasource,
	}, nil
}

func (d *Dataloader) LoadProximaBlockStats(args map[string]interface{}) (*models.ProximaBlockStats, error) {
    var result *proxima.ProximaDBResult;
  if cached, found := d.cache.Get("BlockStats"); found {
		result = cached.(*proxima.ProximaDBResult)
  } else {
    result, _ = d.db.Get(datasources.Primary, "BlockStats", args)
    d.cache.Set("BlockStats", result, BlockStatsCacheExpiration)
  }
    value := models.BlockStats{}
    json.Unmarshal(result.GetValue(), &value)
    proof := GenerateProof(result.GetProof())
    blockStats := &models.ProximaBlockStats{BlockStats: &value, Proof: &proof}
    return blockStats, nil
    //updateBlockStats()
}

func (d *Dataloader) LoadProximaTokens(args map[string]interface{}) (*models.ProximaTokens, error) {
  var result *proxima.ProximaDBResult;
  if cached, found := d.cache.Get("Tokens"); found {
    result = cached.(*proxima.ProximaDBResult)
  } else {
    result, _ = d.db.Get(datasources.Primary, "Tokens", args)
    d.cache.Set("Tokens", result, TokensCacheExpiration)
    }
  	value := []*models.Token{}
  	json.Unmarshal(result.GetValue(), &value)
  	proof := GenerateProof(result.GetProof())
  	tokens :=  &models.ProximaTokens{Tokens: value, Proof: &proof}
    return tokens, nil
}

  func (d *Dataloader) LoadProximaFees(args map[string]interface{}) (*models.ProximaFees, error) {
   var result *proxima.ProximaDBResult;
  if cached, found := d.cache.Get("Fees"); found {
    result= cached.(*proxima.ProximaDBResult)
  } else {
    result, _ = d.db.Get(datasources.Primary, "Fees", args)
    d.cache.Set("Fees", result, FeesCacheExpiration)
    }
  	value := []*models.Fee{}
  	json.Unmarshal(result.GetValue(), &value)
  	proof := GenerateProof(result.GetProof())
  	fees :=  &models.ProximaFees{Fees: value, Proof: &proof}
    return fees, nil
  }

func (d *Dataloader) LoadProximaValidators(args map[string]interface{}) (*models.ProximaValidators, error) {
 var result *proxima.ProximaDBResult;
if cached, found := d.cache.Get("Validators"); found {
  result = cached.(*proxima.ProximaDBResult)
} else {
  result, _ = d.db.Get(datasources.Primary, "Validators", args)
  d.cache.Set("Validators", result, ValidatorsCacheExpiration)
  }
  value := []*models.Validator{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  validators :=  &models.ProximaValidators{Validators: value, Proof: &proof}
  return validators, nil
}

func (d *Dataloader) LoadProximaMarkets(args map[string]interface{}) (*models.ProximaMarkets, error) {
 var result *proxima.ProximaDBResult;
if cached, found := d.cache.Get("Markets"); found {
  result = cached.(*proxima.ProximaDBResult)
} else {
  result, _ = d.db.Get(datasources.Primary, "Markets", map[string]interface{}(args))
  d.cache.Set("Markets", result, MarketsCacheExpiration)
  }
  value := []*models.Market{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  markets :=  &models.ProximaMarkets{Markets: value, Proof: &proof}
  return markets, nil
}

func (d *Dataloader) LoadProximaMarketTickers(args map[string]interface{}) (*models.ProximaMarketTickers, error) {
 var result *proxima.ProximaDBResult;
if cached, found := d.cache.Get("MarketTickers"); found {
  result= cached.(*proxima.ProximaDBResult)
} else {
  result, _ = d.db.Get(datasources.Primary, "MarketTickers", map[string]interface{}(args))
  d.cache.Set("MarketTickers", result, MarketTickersCacheExpiration)
  }
  value := []*models.MarketTicker{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  marketTickers :=  &models.ProximaMarketTickers{MarketTickers: value, Proof: &proof}
  return marketTickers, nil
}

func (d *Dataloader) LoadProximaMarketTicker(args map[string]interface{}) (*models.ProximaMarketTicker, error) {
 var result *proxima.ProximaDBResult;
if cached, found := d.cache.Get("MarketTicker" + args["symbol"].(string)); found {
  result = cached.(*proxima.ProximaDBResult)
} else {
    result, _ = d.db.Get(datasources.MarketTickersBySymbol, args["symbol"], args)
    d.cache.Set("MarketTicker" + args["symbol"].(string), result, MarketTickerCacheExpiration)
  }
  value := models.MarketTicker{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  marketTicker :=  &models.ProximaMarketTicker{MarketTicker: &value, Proof: &proof}
  return marketTicker, nil
}

func (d *Dataloader) LoadProximaMarketDepth(args map[string]interface{}) (*models.ProximaMarketDepth, error) {
 var result *proxima.ProximaDBResult;
if cached, found := d.cache.Get("MarketDepth" + args["symbol"].(string)); found {
  result = cached.(*proxima.ProximaDBResult)
} else {
    result, _ = d.db.Get(datasources.MarketDepthBySymbol, args["symbol"], args)
    if result == nil {
      d.datasource.MarketDepthFetch(args["symbol"].(string))
      result, _ = d.db.Get(datasources.MarketDepthBySymbol, args["symbol"], args)
    }
    d.cache.Set("MarketDepth" + args["symbol"].(string), result, MarketDepthCacheExpiration)
  }
  value := models.MarketDepth{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  marketDepth :=  &models.ProximaMarketDepth{MarketDepth: &value, Proof: &proof}
  return marketDepth, nil
}

func (d *Dataloader) LoadProximaMarketCandleSticks(args map[string]interface{}) (*models.ProximaMarketCandleSticks, error) {
 var result *proxima.ProximaDBResult;
 if cached, found := d.cache.Get("MarketCandleSticks" + args["symbol"].(string) + args["interval"].(string)); found {
  result = cached.(*proxima.ProximaDBResult)
  } else {
  result, _ = d.db.Get(datasources.MarketCandleSticks, args["symbol"].(string) + args["interval"].(string), args)
  if result == nil {
    d.datasource.MarketCandlesticksFetch(args)
    result, _ = d.db.Get(datasources.MarketCandleSticks, args["symbol"].(string) + args["interval"].(string), args)
  }
  d.cache.Set("MarketCandleSticks" + args["symbol"].(string) + args["interval"].(string), result, MarketCandleSticksCacheExpiration)
  }
  value :=[]*models.CandleStick{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  marketCandlesticks :=  &models.ProximaMarketCandleSticks{MarketCandlesticks: value, Proof: &proof}
  return marketCandlesticks, nil
}

func (d *Dataloader) LoadProximaTransaction(args map[string]interface{}) (*models.ProximaTransaction, error) {
  var result *proxima.ProximaDBResult;
  result, _ = d.db.Get(datasources.TransactionsByTxHash, args["txHash"].(string), args)
  if result == nil {
    d.datasource.TransactionFetch(args)
    result, _ = d.db.Get(datasources.TransactionsByTxHash, args["txHash"].(string), args)
  }
  value := models.Transaction{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  transaction :=  &models.ProximaTransaction{Transaction: &value, Proof: &proof}
  return transaction, nil
}
//Transactions
func (d *Dataloader) LoadProximaTransactions(args map[string]interface{}) ([]*models.ProximaTransaction, error) {
  result, err := d.datasource.TransactionsFetch(args)
  if err != nil {
    return nil, err
  }
  value := []*models.Transaction{}
  json.Unmarshal(result, &value)
  proof:= models.Proof{Proof: nil,  Root: nil}
  transactions := make([]*models.ProximaTransaction, len(value))
  for i, transaction := range value {
    transactions[i] = &models.ProximaTransaction{Transaction: transaction, Proof: &proof}
  }
  return transactions, nil
}

func (d *Dataloader) LoadProximaAtomicSwap(args map[string]interface{}) (*models.ProximaAtomicSwap, error) {
  result, err := d.datasource.AtomicSwapFetch(args)
  if err != nil {
    return nil, err
  }
  value := models.AtomicSwap{}
  json.Unmarshal(result, &value)
  proof:= models.Proof{Proof: nil,  Root: nil}
  atomicSwap :=  &models.ProximaAtomicSwap{AtomicSwap: &value, Proof: &proof}
  return atomicSwap, nil
}
//AtomicSwap //No cache, No DB NOTDONE
func (d *Dataloader) LoadProximaAtomicSwaps(args map[string]interface{}) ([]*models.ProximaAtomicSwap, error) {
  result, err := d.datasource.AtomicSwapsFetch(args)
  if err != nil {
    return nil, err
  }
  value := []*models.AtomicSwap{}
  json.Unmarshal(result, &value)
  proof:= models.Proof{Proof: nil,  Root: nil}
  atomicSwaps := make([]*models.ProximaAtomicSwap, len(value))
  for i, atomicSwap := range value {
    atomicSwaps[i] = &models.ProximaAtomicSwap{AtomicSwap: atomicSwap, Proof: &proof}
  }
  return atomicSwaps, nil
}

func (d *Dataloader) LoadProximaOrder(args map[string]interface{}) (*models.ProximaOrder, error) {
  var result *proxima.ProximaDBResult;
  result, _ = d.db.Get(datasources.OrdersByOrderId, args["orderId"].(string), args)
  if result == nil {
    d.datasource.OrderFetch(args)
    result, _ = d.db.Get(datasources.OrdersByOrderId, args["orderId"].(string), args)
  }
  value := models.Order{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  order :=  &models.ProximaOrder{Order: &value, Proof: &proof}
  return order, nil
}
//Orders //No Cache, No DB
func (d *Dataloader) LoadProximaOrders(args map[string]interface{}) ([]*models.ProximaOrder, error) {
  result, err := d.datasource.OrdersFetch(args)
  if err != nil {
    return nil, err
  }
  value := []*models.Order{}
  json.Unmarshal(result, &value)
  proof:= models.Proof{Proof: nil,  Root: nil}
  orders := make([]*models.ProximaOrder, len(value))
  for i, order := range value {
    orders[i] = &models.ProximaOrder{Order: order, Proof: &proof}
  }
  return orders, nil
}

func (d *Dataloader) LoadProximaAccount(args map[string]interface{}) (*models.ProximaAccount, error) {
  var result *proxima.ProximaDBResult;
  result, err := d.db.Get(datasources.AccountsByAddress, args["address"].(string), args)
  if result == nil ||  err != nil{
    d.datasource.AccountFetch(args)
    result, _ = d.db.Get(datasources.AccountsByAddress, args["address"].(string), args)
  }
  value := models.Account{}
  json.Unmarshal(result.GetValue(), &value)
  proof := GenerateProof(result.GetProof())
  account :=  &models.ProximaAccount{Account: &value, Proof: &proof}
  return account, nil
}
//Trades //Cache, DB
func (d *Dataloader) LoadProximaTrades(args map[string]interface{}) ([]*models.ProximaTrade, error) {
  result, err := d.datasource.TradesFetch(args)
  if err != nil {
    return nil, err
  }
  value := []*models.Trade{}
  json.Unmarshal(result, &value)
  proof:= models.Proof{Proof: nil,  Root: nil}
  trades := make([]*models.ProximaTrade, len(value))
  for i, trade := range value {
    trades[i] = &models.ProximaTrade{Trade: trade, Proof: &proof}
  }
  return trades, nil
}

func (d *Dataloader) LoadProximaTimelocks(args map[string]interface{}) (*models.ProximaTimelocks, error) {
  result, err := d.datasource.TimelockFetch(args)
  if err != nil {
    return nil, err
  }
  value := []*models.Timelock{}
  json.Unmarshal(result, &value)
  proof:= models.Proof{Proof: nil,  Root: nil}
  timelocks :=  &models.ProximaTimelocks{Timelocks: value, Proof: &proof}
  return timelocks, nil
}
