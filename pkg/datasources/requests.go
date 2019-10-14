
package datasources

import (
    "io/ioutil"
    "net/http"
  )

var binance_datasource_uri = map[string]func(string, map[string]interface{}) (string) {
  "fees" : fees_request_uri,
  "blockStats" : blockStats_request_uri,
  "validators" : validators_request_uri,
  "tokens" : tokens_request_uri,
  "markets" : markets_request_uri,
  "marketTickers": marketTickers_request_uri,
  "marketDepth": marketDepth_request_uri,
  "marketCandleSticks": marketCandlesticks_request_uri,
  "transaction": transaction_request_uri,
  "trade": trade_request_uri,
  "order": order_request_uri,
  "account" : account_request_uri,
  "atomicSwap" : atomicSwap_request_uri,
}

func binance_request(requestType string, baseUri string, args map[string]interface{}) ([]byte) {
  uri := binance_datasource_uri[requestType](baseUri, args)
  resp, _ := http.Get(uri)
  body, _ := ioutil.ReadAll(resp.Body)
  return body
}

func fees_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/fees"
}

func tokens_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/tokens"
}

func markets_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/markets"
}

func marketTickers_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/ticker/24hr"
}

func marketDepth_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/depth?symbol=" + args["symbol"].(string)
}

func marketCandlesticks_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/klines?symbol="+ args["symbol"].(string) + "&interval=" + args["interval"].(string)
}

func validators_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/validators"
}

func blockStats_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/node-info"
}

func account_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/account/" + args["address"].(string)
}

func trade_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/trades?height=" + args["height"].(string)
}

func transaction_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v2/transactions-in-block/" + args["blockHeight"].(string)
}

func order_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/orders/" + args["orderId"].(string)
}

func atomicSwap_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/atomic-swaps/" + args["swapId"].(string)
}
