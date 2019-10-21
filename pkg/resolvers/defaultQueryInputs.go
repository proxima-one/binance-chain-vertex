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
  "prove": true,
}

var FeesDefaultInputs  =  map[string]interface{}{
  "prove": true,
}

var AtomicSwapDefaultInputs = map[string]interface{} {
    "id": nil,
    "prove": false,
}

var ValidatorsDefaultInputs = map[string]interface{} {
  "prove": true,
}

var TokensDefaultInputs = map[string]interface{} {
  "prove": true,
  "offset": 0,
  "limit": 1000,
}

var MarketsDefaultInputs = map[string]interface{} {
  "prove": true,
  "offset": 0,
  "limit": 1000,
}

var MarketTickersDefaultInputs = map[string]interface{} {
  "prove": true,
  "offset": 0,
  "limit": 1000,
}

var MarketTickerDefaultInputs = map[string]interface{} {
    "symbol": "RAVEN-F66_BNB",
    "prove": true,
}

var MarketDepthDefaultInputs = map[string]interface{} {
    "symbol": "RAVEN-F66_BNB",
    "prove": true,
}

var MarketCandleSticksDefaultInputs = map[string]interface{} {
  "symbol": "RAVEN-F66_BNB",
  "interval": "5m",
    "prove": true,
}

var OrderDefaultInputs = map[string]interface{} {
    "orderId": "37D9383E6AD9AFEF6C5D8066ABA3ACA8C75D9F39-4017145",
    "prove": true,
}

var TransactionDefaultInputs = map[string]interface{} {
    "txHash": "A00D544D5640016D9B6B0D3F59E3AFC1D0157EF3D1C129758A791C135A3391A1",
    "prove": true,
}

var AccountDefaultInputs = map[string]interface{} {
  "address": "bnb1a03uvqmnqzl85csnxnsx2xy28m76gkkht46f2l",
  "prove": true,
}


var TimelocksDefaultInputs = map[string]interface{} {
    "address": "bnb1a03uvqmnqzl85csnxnsx2xy28m76gkkht46f2l",
    "prove": true,
    "id": nil,
}

var OrdersDefaultInputs = map[string]interface{} {
    "prove": true,
    "address": "bnb1a03uvqmnqzl85csnxnsx2xy28m76gkkht46f2l",
    "symbol": nil,
    "status": nil,
    "orderSide": nil,
    "end": nil,
    "start": nil,
    "open":nil,
    "total": nil,
}

var TransactionsDefaultInputs = map[string]interface{} {
    "prove": false,
    "address": "bnb1urm5fr5djgk5grysyrrk2nffr4spkd999nl7u7",
    "txType": nil,
    "txSide": nil,
    "endTime": nil,
    "startTime": nil,
    "blockHeight": 42743260,
    "offset": 0,
    "limit": 1000,
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
    "height": 42743260,
    "symbol": nil,
    "offset": 0,
    "limit": 1000,
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
