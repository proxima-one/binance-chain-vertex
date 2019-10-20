package resolvers
//import "fmt"

func GenerateInputs(inputs map[string]interface{}, defaultInputs map[string]interface{}) (map[string]interface{}) {
  argInputs := make(map[string]interface{})
  for key, value := range defaultInputs {
    argInputs[key] = value
      if inputs[key] != nil {
        argInputs[key] = inputs[key]
      }
    }
    return argInputs
  }


var BlockStatsDefaultInputs = map[string]interface{} {
  "prove": false,
}

var FeesDefaultInputs  =  map[string]interface{}{
  "prove": false,
}

var AccountDefaultInputs = map[string]interface{} {
  "address": "bnb1a03uvqmnqzl85csnxnsx2xy28m76gkkht46f2l",
  "prove": false,
}

var ValidatorsDefaultInputs = map[string]interface{} {
  "prove": false,
}

var TokensDefaultInputs = map[string]interface{} {
  "prove": false,
  "offset": 0,
  "limit": 1000,
}

var MarketsDefaultInputs = map[string]interface{} {
  "prove": false,
  "offset": 0,
  "limit": 1000,
}

var MarketTickersDefaultInputs = map[string]interface{} {
  "prove": false,
  "offset": 0,
  "limit": 1000,
}

var MarketTickerDefaultInputs = map[string]interface{} {
    "symbol": "RAVEN-F66_BNB",
    "prove": false,
}

var MarketDepthDefaultInputs = map[string]interface{} {
    "symbol_pair": "RAVEN-F66_BNB",
    "prove": false,
}

var MarketCandleSticksDefaultInputs = map[string]interface{} {
    "prove": false,
}

var OrderDefaultInputs = map[string]interface{} {
    "orderId": "E0F7448E8D922D440C9020C7654D291D601B34A5",
    "prove": false,
}

var OrdersDefaultInputs = map[string]interface{} {
    "prove": false,
    "address": "414FB3BBA216AF84C47E07D6EBAA2DCFC3563A2F",
    "symbol": nil,
    "status": nil,
    "orderSide": nil,
    "end": nil,
    "start": nil,
    "open":nil,
    "total": nil,
}

var TransactionDefaultInputs = map[string]interface{} {
    "txHash": "A00D544D5640016D9B6B0D3F59E3AFC1D0157EF3D1C129758A791C135A3391A1",
    "prove": false,
}

var TransactionsDefaultInputs = map[string]interface{} {
    "prove": false,
    "address": "414FB3BBA216AF84C47E07D6EBAA2DCFC3563A2F",
    "txType": nil,
    "txSide": nil,
    "endTime": nil,
    "startTime": nil,
    "blockHeight": "42250183",
    "offset": 0,
    "limit": 1000,
}

var TradeDefaultInputs = map[string]interface{} {
    "tradeId": nil,
    "prove": false,
}

var TradesDefaultInputs = map[string]interface{} {
    "address": "bnb1urm5fr5djgk5grysyrrk2nffr4spkd999nl7u7",
    "quoteAssetSymbol": nil,
    "buyerOrderId": nil,
    "sellerOrderId": nil,
    "orderSide": nil,
    "prove": false,
    "endTime": nil,
    "startTime": nil,
    "blockHeight": "42250183",
    "symbol": nil,
    "offset": 0,
    "limit": 1000,
}

var AtomicSwapDefaultInputs = map[string]interface{} {
    "id": nil,
    "prove": false,
}

var AtomicSwapsDefaultInputs = map[string]interface{} {
  "fromAddress":"bnb1urm5fr5djgk5grysyrrk2nffr4spkd999nl7u7",
  "toAddress": nil,
  "endTime": nil,
  "startTime": nil,
  "symbol": nil,
  "offset": 0,
  "limit": 1000,
  "prove": false,
}

var TimelocksDefaultInputs = map[string]interface{} {
    "address": nil,
    "prove": false,
    "id": nil,
}
