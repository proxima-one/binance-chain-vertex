package datasources

import (
  url "net/url"
  time "time"
  //json "github.com/json-iterator/go"
  //"net/http"
  //"io/ioutil"
  json "github.com/json-iterator/go"
  proxima "github.com/proxima-one/proxima-db-client-go"
  _ "fmt"
)

type Datasource struct{
  uri *url.URL
  baseUri string
  proximaDB *proxima.ProximaDB
}

func (ds *Datasource) Start() {


  go ds.StaticUpdates()
  go ds.DynamicUpdates()
}

func (ds *Datasource) StaticUpdates() {
    static_interval := 24*time.Hour
  for {
      go ds.FeesFetch()
      go ds.ValidatorsFetch()
      go ds.MarketsFetch()
      go ds.TokensFetch()
      time.Sleep(static_interval)
  }
}

func (ds *Datasource) DynamicUpdates() {
  dynamic_interval := 5*time.Second
  for {
    go ds.MarketUpdates()
    blockStats, err := ds.BlockStatsFetch() //err
    if err == nil {
      blockHeight := blockStats["latest_block_height"].(string)
      go ds.TransactionUpdates(blockHeight)
      go ds.UpdateTrades(blockHeight)
    }
    time.Sleep(dynamic_interval)
  }
}

func NewDatasource(db *proxima.ProximaDB, rawURI string) (*Datasource, error) {
  uri, err := url.Parse(rawURI)
  if err != nil {
    return nil, err
  }
  return &Datasource{uri: uri, baseUri: rawURI, proximaDB: db}, nil
}

func (ds *Datasource) MarketUpdates() (bool, error) {
  marketTickers, err := ds.MarketTickersFetch()
  if err != nil {
    return false, err
  }
  args := make(map[string]interface{})
  for _, marketTicker := range marketTickers {
    ds.MarketDepthFetch(marketTicker["symbol"].(string))
    ds.proximaDB.Set(MarketTickersBySymbol, marketTicker["symbol"].(string), marketTicker, args)
  }
  return true, nil
}

func (ds *Datasource) UpdateTrades(height string) (bool, error){
  args := make(map[string]interface{});
  args["height"] = height
  trades, err := ds.DataRequest("trade", args)
  if err != nil {
    return false, err
  }
  tradesByTime := make(map[string]interface{});
  tradesBySellerId := make(map[string]interface{});
  tradesByBuyerId := make(map[string]interface{});
  for _, trade := range trades.([]map[string]interface{}) {
    ds.proximaDB.Set(TradesByTradeId, trade["tradeId"], trade, args)
    tradesByTime[trade["time"].(string)] = append(tradesByTime[trade["time"].(string)].([]interface{}), trade)
    tradesByBuyerId[trade["buyerId"].(string)] = append(tradesByBuyerId[trade["buyerId"].(string)].([]interface{}), trade)
    tradesBySellerId[trade["sellerId"].(string)] = append(tradesBySellerId[trade["sellerId"].(string)].([]interface{}), trade)
  }
  ds.UpdateData(TradesByTime, tradesByTime)
  ds.UpdateData(TradesByBuyerId, tradesByBuyerId)
  ds.UpdateData(TradesBySellerId, tradesBySellerId)
  return true, nil
}

func (ds *Datasource) TransactionUpdates(blockHeight string)  (bool, error)  {
  args := make(map[string]interface{});
  args["blockHeight"] = blockHeight
  txs, err := ds.DataRequest("transaction", args)
  if err != nil {
    return false, err
  }
  transactions := txs.([]map[string]interface{})
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

  ds.UpdateData(TransactionsByTimeStamp, (transactionsByTimeStamp))
  ds.UpdateData(TransactionsByToAddr, map[string]interface{}(transactionsByToAddr))
  ds.UpdateData(TransactionsByFromAddr, map[string]interface{}(transactionsByFromAddr))

  ds.AccountsFetchAll(accounts)
  ds.OrdersFetchAll(orderIds)
  return true, nil
}


func (ds *Datasource) UpdateData(table_name string, dataMap map[string]interface{}) (bool, error) {
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

func (ds *Datasource) AccountsFetchAll(accounts map[string]bool) {
  args := make(map[string]interface{})
  for address, _ := range accounts {
    args["address"] = address
    ds.AccountFetch(args)
  }
}

func (ds *Datasource) OrdersFetchAll(orders map[string]bool) {
  args := make(map[string]interface{})
  for orderId, _ := range orders {
    args["orderId"] = orderId
    ds.OrderFetch(args)
  }
}

func (ds *Datasource) BatchUpdate(table_name, dataMap map[string]interface{}) (bool, error) {
  requests := make([]interface{}, 0)
  args:= make(map[string]interface{})
  for key, value := range dataMap {
    requests = append(requests, map[string]interface{}{"table": table_name, "key": key, "value": value, "prove":false}) //get the value
  }
  ds.proximaDB.Batch(requests, args)
  return true, nil
}
