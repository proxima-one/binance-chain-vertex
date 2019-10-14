package datasources

import (
  url "net/url"
  time "time"
  "encoding/json"
  proxima "github.com/proxima-one/proxima-db-client-go"
)

type Datasource struct{
  uri *url.URL
  baseUri string
  proximaDB *proxima.ProximaDB
}

func (ds *Datasource) Start() {
  static_interval := 10000
  dynamic_interval := 5000
  go ds.staticUpdates(static_interval)
  go ds.dynamicUpdates(dynamic_interval)
}

func (ds *Datasource) staticUpdates(interval int) {
  for {
      go ds.updateFees()
      go ds.updateValidators()
      go ds.updateMarkets()
      go ds.updateTokens()
      time.Sleep(time.Duration(interval) * time.Millisecond)
  }
}

func (ds *Datasource) dynamicUpdates(interval int) {
  for {
    blockStats := ds.updateBlockStats()
    blockHeight := blockStats["latest_block_height"].(string)
    go ds.marketUpdates()
    go ds.transactionUpdates(blockHeight)
    go ds.updateTrades(blockHeight)
    time.Sleep(time.Duration(interval) * time.Millisecond)
  }
}

func NewDatasource(db *proxima.ProximaDB, rawURI string) (*Datasource, error) {
  uri, err := url.Parse(rawURI)
  if err != nil {
    return nil, err
  }
  return &Datasource{uri: uri, baseUri: rawURI, proximaDB: db}, nil
}

func (ds *Datasource) updateFees() {
  args := make(map[string]interface{});
  fees := ds.dataRequest("fees", args);
  ds.proximaDB.Set(Primary, "Fees", fees, args)
}

func (ds *Datasource) updateValidators() {
  args := make(map[string]interface{});
  validators := ds.dataRequest("validators", args)
  ds.proximaDB.Set(Primary, "Validators", validators, args)
}

func (ds *Datasource) updateTokens() {
  args := make(map[string]interface{});
  tokens := ds.dataRequest("tokens", args)
  ds.proximaDB.Set(Primary, "Tokens", tokens, args)
}

func (ds *Datasource) updateMarkets() {
  args := make(map[string]interface{});
  markets := ds.dataRequest("markets", args)
  ds.proximaDB.Set(Primary, "Markets", markets, args)
}

func (ds *Datasource) updateBlockStats() (map[string]interface{}) {
  args := make(map[string]interface{});
  blockStats := ds.dataRequest("blockStats", args).(map[string]interface{})
  ds.proximaDB.Set(Primary, "BlockStats", blockStats, args)
  return blockStats
}

func (ds *Datasource) marketUpdates() {
  marketTickers := ds.updateMarketTickers().([]map[string]interface{})
  args := make(map[string]interface{})
  for _, marketTicker := range marketTickers {
    ds.updateMarketDepth(marketTicker["symbol"].(string))
    ds.proximaDB.Set(MarketTickersBySymbol, marketTicker["symbol"].(string), marketTicker, args)
  }
}

func (ds *Datasource) updateMarketTickers() (interface{}) {
  args := make(map[string]interface{});
  marketTickers := ds.dataRequest("marketTickers", args).([]map[string]interface{})
  ds.proximaDB.Set(Primary, "MarketTickers", marketTickers, args)
  return marketTickers
}

func (ds *Datasource) updateMarketDepth(symbol string) {
  args:= map[string]interface{}{"symbol": symbol,}
  data := ds.dataRequest("marketDepth", args)
  ds.proximaDB.Set(MarketDepthBySymbol, symbol, data, args)
}

func (ds *Datasource) updateTrades(height string) {
  args := make(map[string]interface{});
  args["height"] = height
  trades := ds.dataRequest("trade", args)
  tradesByTime := make(map[string]interface{});
  tradesBySellerId := make(map[string]interface{});
  tradesByBuyerId := make(map[string]interface{});
  for _, trade := range trades.([]map[string]interface{}) {
    ds.proximaDB.Set(TradesByTradeId, trade["tradeId"], trade, args)
    tradesByTime[trade["time"].(string)] = append(tradesByTime[trade["time"].(string)].([]interface{}), trade)
    tradesByBuyerId[trade["buyerId"].(string)] = append(tradesByBuyerId[trade["buyerId"].(string)].([]interface{}), trade)
    tradesBySellerId[trade["sellerId"].(string)] = append(tradesBySellerId[trade["sellerId"].(string)].([]interface{}), trade)
  }
  ds.updateData(TradesByTime, tradesByTime)
  ds.updateData(TradesByBuyerId, tradesByBuyerId)
  ds.updateData(TradesBySellerId, tradesBySellerId)
}

func (ds *Datasource) transactionUpdates(blockHeight string)  {
  args := make(map[string]interface{});
  args["blockHeight"] = blockHeight
  transactions := ds.dataRequest("transaction", args).([]map[string]interface{})
  ds.proximaDB.Set(TransactionsByBlockHeight, blockHeight, transactions, args)
  transactionsByFromAddr := make(map[string]interface{});
  transactionsByToAddr := make(map[string]interface{});
  transactionsByTimeStamp := make(map[string]interface{});
  orderIds := make(map[string]bool)
  accounts := make(map[string]bool)

  for _, transaction := range transactions {
    ds.proximaDB.Set(TransactionsByTxHash, transaction["txHash"], transaction, args)
    accounts[transaction["fromAddr"].(string)] = true
    accounts[transaction["toAddr"].(string)] = true
    orderIds[transaction["orderId"].(string)] = true

    transactionsByTimeStamp[transaction["timeStamp"].(string)] = append(transactionsByTimeStamp[transaction["timeStamp"].(string)].([]interface{}), transaction)
    transactionsByFromAddr[transaction["fromAddr"].(string)] = append(transactionsByFromAddr[transaction["fromAddr"].(string)].([]interface{}), transaction)
    transactionsByToAddr[transaction["toAddr"].(string)] = append(transactionsByToAddr[transaction["toAddr"].(string)].([]interface{}), transaction)
  }

  ds.updateData(TransactionsByTimeStamp, (transactionsByTimeStamp))
  ds.updateData(TransactionsByToAddr, map[string]interface{}(transactionsByToAddr))
  ds.updateData(TransactionsByFromAddr, map[string]interface{}(transactionsByFromAddr))

  ds.updateAccounts(accounts)
  ds.updateOrdersByOrderId(orderIds)
}

func (ds *Datasource) updateAccounts(accounts map[string]bool) {
  dbArgs := make(map[string]interface{});
  account := make(map[string]interface{});
  addressKey := "address"
  for address, _ := range accounts {
    requestArgs := map[string]interface{}{
      addressKey : address,
    }
    account = ds.dataRequest("account", requestArgs).(map[string]interface{})
    ds.proximaDB.Set(AccountsByAddress, address, account, dbArgs)
  }
}
//Orders need to be done by account //
func (ds *Datasource) updateOrdersByOrderId(orderIds map[string]bool) {
  dbArgs := make(map[string]interface{});
  order:= make(map[string]interface{});
  orderKey := "orderId"
  for orderId, _ := range orderIds {
    requestArgs := map[string]interface{}{
      orderKey : orderId,
    }
    order = ds.dataRequest("order", requestArgs).(map[string]interface{})
    ds.proximaDB.Set(OrdersByOrderId, orderId, order, dbArgs)
  }
}

func (ds *Datasource) updateData(table_name string, dataMap map[string]interface{}) (bool, error) {
  args:= make(map[string]interface{})
  for key, value := range dataMap {
    result, _ := ds.proximaDB.Get(table_name, key, args) //get the value
    val := make([]interface{}, 0)
    json.Unmarshal(result.GetValue(), &val)
    value = append(val, value)
    ds.proximaDB.Set(table_name, key, value, args)
  }
  return true, nil
}

func (ds *Datasource) batchUpdate(table_name, dataMap map[string]interface{}) (bool, error) {
  requests := make([]interface{}, 0)
  args:= make(map[string]interface{})
  for key, value := range dataMap {
    requests = append(requests, map[string]interface{}{"table": table_name, "key": key, "value": value, "prove":false}) //get the value
  }
  ds.proximaDB.Batch(requests, args)
  return true, nil
}

func (ds *Datasource) dataRequest(requestType string, args map[string]interface{}) (interface{}) {
  resp := binance_request(requestType, ds.baseUri, args)
  val := binance_translate(requestType, resp)
  return val
}
