package common

//FetchQueries


var FetchTestCases = map[string]map[string]interface{}{
  "marketCandleSticks":MarketCandleSticksDefaultInputs,
  "transaction": TransactionDefaultInputs ,
  "transaction_tx":TransactionDefaultInputs,
  "order": OrderDefaultInputs,
 "account":AccountDefaultInputs,
 //"atomicSwap" ,
 //"timelocks": TimelocksDefaultInputs,
  "trades":TradesDefaultInputs ,
  //"atomicSwaps": AtomicSwapsDefaultInputs,
//  "transactions":TransactionsDefaultInputs,
  "marketDepth":MarketDepthDefaultInputs,
  "orders": OrdersDefaultInputs,
  "validators":make(map[string]interface{}),
  "tokens":make(map[string]interface{}),
  "markets":make(map[string]interface{}),
  "fees":make(map[string]interface{}),
  "marketTickers":make(map[string]interface{}),
  };

  var AccountDefaultInputs = map[string]interface{} {
    "address": "bnb1a03uvqmnqzl85csnxnsx2xy28m76gkkht46f2l",
    "prove": true,
  }

  var MarketDepthDefaultInputs = map[string]interface{} {
      "symbol": "RAVEN-F66_BNB",
      "prove": true,
  }

  var MarketCandleSticksDefaultInputs = map[string]interface{} {
      "symbol": "RAVEN-F66_BNB",
      "interval": "5m",
      "prove": false,
  }

  var TransactionDefaultInputs = map[string]interface{} {
      "txHash": "A00D544D5640016D9B6B0D3F59E3AFC1D0157EF3D1C129758A791C135A3391A1",
      "prove": true,
  }

  var TimelocksDefaultInputs = map[string]interface{} {
      "address": "bnb1a03uvqmnqzl85csnxnsx2xy28m76gkkht46f2l",
      "prove": false,
      "id": nil,
  }

  var OrderDefaultInputs = map[string]interface{} {
      "orderId": "37D9383E6AD9AFEF6C5D8066ABA3ACA8C75D9F39-4017145",
      "prove": true,
  }

  var OrdersDefaultInputs = map[string]interface{} {
      "prove": false,
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
      "address": nil,
      "txType": nil,
      "txSide": nil,
      "endTime": nil,
      "startTime": nil,
      "blockHeight": 52743260,
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
      "height": nil,
      "symbol": nil,
      "offset": 0,
      "limit": 1000,
  }

  var AtomicSwapDefaultInputs = map[string]interface{} {
      "id": nil,
      "prove": false,
  }

  var AtomicSwapsDefaultInputs = map[string]interface{} {
    "fromAddress":"bnb1a03uvqmnqzl85csnxnsx2xy28m76gkkht46f2l",
    "toAddress": nil,
    "endTime": nil,
    "startTime": nil,
    "symbol": nil,
    "offset": 0,
    "limit": 1000,
    "prove": false,
  }
