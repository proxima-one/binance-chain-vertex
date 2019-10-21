
package datasources

import (
    "net/url"
   "fmt"
   //"reflect"
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
  "transaction_tx" : transaction_tx_request_uri,
  "trade": trade_request_uri,
  "order": order_request_uri,
  "account" : account_request_uri,
  "atomicSwap" : atomicSwap_request_uri,
  "timelocks": timelocks_request_uri,
  "trades" :trades_request_uri,
  "atomicSwaps": atomicSwaps_request_uri,
  "transactions":transactions_request_uri,
  "orders": orders_request_uri,
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
  return baseUri + "/api/v1/trades?height=" + fmt.Sprintf("%v", args["height"])
}

func transaction_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v2/transactions-in-block/" + fmt.Sprintf("%v", args["blockHeight"])
}

func transaction_tx_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/tx/" + args["txHash"].(string)
}

func timelocks_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/timelocks/" + args["address"].(string)
}
//trades request /api/v1/tradess?address=bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38
func trades_request_uri(baseUri string, args map[string]interface{}) (string) {
  return ProcessURL(baseUri + "/api/v1/trades", args)
}

//atomicSwaps ///atomic-swaps?toAddress=bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38
func atomicSwaps_request_uri(baseUri string, args map[string]interface{}) (string) {
  return ProcessURL(baseUri + "/api/v1/atomic-swaps", args)

}

//transactions ///?address=bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38
func transactions_request_uri(baseUri string, args map[string]interface{}) (string) {
  return ProcessURL(baseUri + "/api/v1/transactions", args)
}

//orders https://dex.binance.org/api/v1/orders/open?address=bnb1xlvns0n2mxh77mzaspn2hgav4rr4m8eerfju38
func order_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/orders/" + args["orderId"].(string)
}

func orders_request_uri(baseUri string, args map[string]interface{}) (string) {
  return ProcessURL(baseUri + "/api/v1/orders/close", args)
}

func atomicSwap_request_uri(baseUri string, args map[string]interface{}) (string) {
  return baseUri + "/api/v1/atomic-swaps/" + args["swapId"].(string)
}


func ProcessURL(baseUri string, args map[string]interface{}) (string) {
  urlArgs := processProximaArgs(args)
  urlValue := url.Values{}
  for key, value := range urlArgs {
    urlValue.Set(key, value)
  }
  u := &url.URL{
    Host: baseUri,
		RawQuery: urlValue.Encode(),
	}
  u.Opaque = baseUri
  return u.String()

}

func processProximaArgs(args map[string]interface{}) (map[string]string) {
  cleanedArgs := make(map[string]string)
  val := fmt.Sprintf("%v", interface{}(nil))
  for key, value := range args {
    v:= fmt.Sprintf("%v", value)
    if v != val && key != "prove" {
      cleanedArgs[key] = v
    }
  }


  return cleanedArgs
}
