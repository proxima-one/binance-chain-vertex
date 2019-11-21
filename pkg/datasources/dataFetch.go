package datasources

import (
  "net/http"
  "io/ioutil"
    json "github.com/json-iterator/go"
  "fmt"
  "errors"
)

func  (d *Datasource) AccountFetch(args map[string]interface{}) (map[string]interface{}, error){
  account, err:= d.DataRequest("account", args)
  if err != nil {
    return nil, err
  }
  d.proximaDB.Set(AccountsByAddress, args["address"], account, args)
  return account.(map[string]interface{}), nil
}

func  (d *Datasource) OrderFetch(args map[string]interface{}) (map[string]interface{}, error) {
  order, err:= d.DataRequest("order", args)
  if err != nil {
    return nil, err
  }
  d.proximaDB.Set(OrdersByOrderId, args["orderId"], order.(map[string]interface{}), args)
  return order.(map[string]interface{}), nil
}
func  (d *Datasource) OrdersFetch(args map[string]interface{}) ([]byte, error) {
  orders, err := d.DataRequest("orders", args)
  if err != nil {
    return nil, err
  }
  byteValue, _ := json.Marshal(orders)
  return []byte(byteValue), nil
}

func (d *Datasource) TradesFetch(args map[string]interface{}) ([]byte, error) {
  trades, err := d.DataRequest("trades", args)
  if err != nil {
    return nil, err
  }
  byteValue, _ := json.Marshal(trades)
  return []byte(byteValue), nil
}

func (d *Datasource) MarketCandlesticksFetch(args map[string]interface{}) ([]map[string]interface{}, error) {
  ms, err := d.DataRequest("marketCandleSticks", args)
  if err != nil || ms == nil {
    return nil, err
  }
  marketCandleSticks := ms.([]map[string]interface{})
  d.proximaDB.Set(MarketCandleSticks, args["symbol"].(string) + args["interval"].(string), marketCandleSticks, args)
  return marketCandleSticks, nil
}

func  (d *Datasource) TransactionFetch(args map[string]interface{}) (map[string]interface{}, error){
  transaction, err := d.DataRequest("transaction_tx", args)
  if err != nil {
    return nil, err
  }
  d.proximaDB.Set(TransactionsByTxHash, args["txHash"], transaction.(map[string]interface{}), args)
  return transaction.(map[string]interface{}), nil
}

func  (d *Datasource) TransactionsFetch(args map[string]interface{}) ([]byte, error) {  //"net/http"
  transactions, err := d.DataRequest("transactions", args)
  if err != nil {
    return nil, err
  }
  byteValue, _ := json.Marshal(transactions)
  return []byte(byteValue), nil
}

func (ds *Datasource) MarketTickersFetch() ([]map[string]interface{}, error) {
  args := make(map[string]interface{});
  marketTickers, err := ds.DataRequest("marketTickers", args)
  if err != nil {
    return nil, err
  }
  ds.proximaDB.Set(Primary, "MarketTickers", marketTickers.([]map[string]interface{}), args)
  return marketTickers.([]map[string]interface{}), nil
}

func (ds *Datasource) MarketDepthFetch(symbol string) (map[string]interface{}, error) {
  args:= map[string]interface{}{"symbol": symbol,}
  data, err := ds.DataRequest("marketDepth", args)
  if err != nil {
    return nil, err
  }
  ds.proximaDB.Set(MarketDepthBySymbol, symbol, data.(map[string]interface{}), args)
  return data.(map[string]interface{}), nil
}

func  (d *Datasource)  AtomicSwapFetch(args map[string]interface{}) ([]byte, error){
  atomicSwap, err := d.DataRequest("atomicSwap", args)
  if err != nil {
    return nil, err
  }
  return atomicSwap.([]byte), nil
}

func (d *Datasource)  AtomicSwapsFetch(args map[string]interface{})([]byte, error) {
  atomicSwaps, err := d.DataRequest("atomicSwaps", args)
  if err != nil {
    return nil, err
  }
  byteValue, _ := json.Marshal(atomicSwaps)
  return []byte(byteValue), nil
}

func  (d *Datasource) TimelockFetch(args map[string]interface{}) ([]byte, error){
  timelock, err := d.DataRequest("timelock", args)
  if err != nil {
    return nil, err
  }
  return timelock.([]byte), nil
}

func (ds *Datasource) BlockStatsFetch() (map[string]interface{}, error) {
  args := make(map[string]interface{});
  blockStats, err := ds.DataRequest("blockStats", args)
  if err != nil || blockStats == nil {
    return nil, err
  }
  ds.proximaDB.Set(Primary, "BlockStats", blockStats.(map[string]interface{}), args)
  return blockStats.(map[string]interface{}), nil
}

func (ds *Datasource) ValidatorsFetch() ([]map[string]interface{}, error) {
  args := make(map[string]interface{});
  validators, err := ds.DataRequest("validators", args)
  if err != nil || validators == nil {
    return nil, err
  }
  ds.proximaDB.Set(Primary, "Validators", validators.([]map[string]interface{}), args)
  return validators.([]map[string]interface{}), nil
}

func (ds *Datasource) TokensFetch() ([]map[string]interface{}, error) {
  args := make(map[string]interface{});
  tokens, err := ds.DataRequest("tokens", args) //errs
  if err != nil || tokens == nil {
    return nil, err
  }
  ds.proximaDB.Set(Primary, "Tokens", tokens, args)
  return tokens.([]map[string]interface{}), nil
}

func (ds *Datasource) MarketsFetch()  ([]map[string]interface{}, error){
  args := make(map[string]interface{});
  markets, err := ds.DataRequest("markets", args) //errs
  if err != nil || markets == nil {
    return nil, err
  }
  ds.proximaDB.Set(Primary, "Markets", markets, args)
  return markets.([]map[string]interface{}), nil
}

func (ds *Datasource) FeesFetch() ([]map[string]interface{}, error) {
  args := make(map[string]interface{});
  fees, err := ds.DataRequest("fees", args);
  if err != nil || fees == nil {
    return nil, err
  }
  ds.proximaDB.Set(Primary, "Fees", fees, args)
  return fees.([]map[string]interface{}), nil
}

func (ds *Datasource) DataRequest(requestType string, args map[string]interface{}) (interface{}, error) {
  resp, err := BinanceRequest(requestType, ds.baseUri, args)
  if err != nil {
    return nil, err
  }
  val, tErr := BinanceTranslate(requestType, resp)
  if val == nil {
    fmt.Println(requestType)
    fmt.Println(args)
    fmt.Println(string(resp))
    return nil, errors.New("Error with translation of vars")
  }

  if tErr != nil {
    return nil, tErr
  }
  return val, nil
}

func BinanceRequest(requestType string, baseUri string, args map[string]interface{}) ([]byte, error) {
  uri := binance_datasource_uri[requestType](baseUri, args)

  resp, httpErr := http.Get(uri)
  if httpErr != nil {
    //fmt.Println(uri)
    return nil, httpErr
  }
  body, ioErr := ioutil.ReadAll(resp.Body)
  if ioErr != nil {
    //fmt.Println(uri)
    return nil, ioErr
  }
  return body, nil
}
