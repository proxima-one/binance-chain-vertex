package datasources

import (
  //json "github.com/json-iterator/go"
  json "github.com/json-iterator/go"
  "strconv"
  "fmt"
  //"log"
)

func BinanceTranslate(requestType string, res []byte) (interface{}, error) {
  return TRANSLATE_RES[requestType](res)
}

var TRANSLATE_RES = map[string]func([]byte) (interface{}, error) {
  "fees": feesTranslate,
  "tokens": tokensTranslate,
  "blockStats": blockStatsTranslate,
  "markets": marketsTranslate,
  "validators":validatorsTranslate,
  "marketDepth":marketDepthTranslate,
  "marketTickers":marketTickersTranslate,
  "marketCandleSticks": marketCandlesticksTranslate,
  "account": accountTranslate,
  "transaction": transactionsTranslate,
  "transaction_tx": transactionTranslate,
  "transactions" : transactionQueryTranslate,
  "order": orderTranslate,
  "atomicSwap": atomicSwapTranslate,
  "timelocks": timeLocksTranslate,
  "atomicSwaps": atomicSwapsTranslate,
  "trade": tradeTranslate,
  "trades": tradesTranslate,
  "orders": ordersTranslate,
 }

func feesTranslate(res []byte) (interface{}, error) {
  fees:= make([]map[string]interface{}, 0)
  err:=json.Unmarshal(res, &fees)
  if err!= nil {
    return nil, err
  }
  for _ , fee := range fees {
    if fee["msg_type"] == nil && fee["dex_fee_fields"] != nil {
      fee["msg_type"] = "dex_fee_fields"
    }

    if fee["msg_type"] == "" && fee["fixed_fee_fields"] != nil {
     fee["msg_type"] = "fixed_fee_fields"
    }
  }
  return fees, nil
}

func blockStatsTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  sync_val := val["sync_info"]
  if sync_val == nil {
      sync_val = make(map[string]interface{})
      return sync_val, nil
  } else {
    v := sync_val.(map[string]interface{})
    temp:= int64(v["latest_block_height"].(float64))
    v["latest_block_height"] = strconv.FormatInt(temp, 10)
    return v, nil
  }
}

func validatorsTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  if val != nil && val["validators"] != nil {
    v := make([]map[string]interface{}, len(val["validators"].([]interface{})))
    for i, value := range val["validators"].([]interface{}) {
      vals := value.(map[string]interface{})
      v[i] = vals
    }
    return v, nil
  }
  return nil, nil
}

func tokensTranslate(res []byte) (interface{}, error) {
  val := make([]map[string]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func marketsTranslate(res []byte) (interface{}, error) {
  val := make([]map[string]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func marketTickersTranslate(res []byte) (interface{}, error) {
  val := make([]map[string]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func marketTickerTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func marketCandlesticksTranslate(res []byte) (interface{}, error) {
  val := make([]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err != nil || val == nil || len(val) == 0 {
    return nil, err
  }
  returnVal := make([]map[string]interface{}, len(val))
  for i, v := range val {
    vals := v.([]interface{})
    value := make(map[string]interface{})
    value["close"] =  vals[0].(float64)
    value["closingTime"] = fmt.Sprintf("%v", vals[1])
    value["high"], _ = strconv.ParseFloat(vals[2].(string), 64)
    value["low"], _ = strconv.ParseFloat(vals[3].(string), 64)
    value["numberOfTrades"], _ = strconv.Atoi(vals[4].(string))
    value["open"], _ = strconv.ParseFloat(vals[5].(string), 64)
    value["openTime"] = fmt.Sprintf("%v", vals[6])
    value["quoteAssetVolume"], _ = strconv.ParseFloat(vals[7].(string), 64)
    value["volume"] =  vals[8].(float64)
    returnVal[i] = value
  }
  return returnVal, nil
}

func marketDepthTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil || val == nil {
    return nil, err
  }
  if val["asks"] == nil {
    val["asks"] = make([]map[string]string,0)
  } else {
    val_Ask := val["asks"].([]interface{})
    length := len(val_Ask)
    vAsk := make([]map[string]string,length)
    for i, v :=  range val_Ask {
      value := v.([]interface{})
      vAsk[i] = map[string]string{"price": fmt.Sprintf("%v", value[0]), "qty":fmt.Sprintf("%v", value[1])}
    }
    val["asks"] = vAsk
  }
  if val["bids"] == nil {
    val["bids"] = make([]map[string]string,0)
  } else {
    val_Bid := val["bids"].([]interface{})
    length := len(val_Bid)
    vBid := make([]map[string]string,length)
    for i, v :=  range val_Bid {
      value := v.([]interface{})
      vBid[i] = map[string]string{"price": fmt.Sprintf("%v", value[0]), "qty":fmt.Sprintf("%v", value[1])}
    }
    val["bids"] = vBid
  }
  delete(val, "height")
  return val, nil
}

func transactionQueryTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil || val == nil || val["tx"] == nil {
    return nil, err
  }
  v := val["tx"].([]interface{})
  returnVal := make([]map[string]interface{}, len(v))
  for i, va := range v {
    value := va.(map[string]interface{})

    h:= int64(value["blockHeight"].(float64))
    value["blockHeight"] = strconv.FormatInt(h, 10)
    value["code"] = fmt.Sprintf("%v", value["code"])
    value["timeStamp"] = fmt.Sprintf("%v", value["timeStamp"])
    value["confirmBlocks"] = fmt.Sprintf("%v", value["confirmBlocks"])
    value["txAge"] = fmt.Sprintf("%v", value["txAge"])
    value["sequence"] = fmt.Sprintf("%v", value["sequence"])
    value["data"] = fmt.Sprintf("%v", value["data"])
    returnVal[i] = value
  }
  return returnVal, nil
}

func accountTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func transactionsTranslate(res []byte) (interface{}, error) {
  val := make([]map[string]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func tradeTranslate(res []byte) (interface{}, error) {
  val := make([]map[string]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}


func tradesTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil || val == nil || val["trade"] == nil {
    return nil, err
  }
  v := val["trade"].([]interface{})
  returnVal := make([]map[string]interface{}, len(v))
  for i, va := range v {
    value := va.(map[string]interface{})

    h:= int64(value["blockHeight"].(float64))
    value["blockHeight"] = strconv.FormatInt(h, 10)

    value["time"] = fmt.Sprintf("%v", value["time"])
    returnVal[i] = value
  }
  return returnVal, nil
}

func orderTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func ordersTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil || val == nil || val["order"] == nil  {
    return nil, err
  }
  return val["order"].([]interface{}), nil
}

func transactionTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func atomicSwapTranslate(res []byte) (interface{}, error) {
  return res, nil
}

func atomicSwapsTranslate(res []byte) (interface{}, error) {
  return res, nil
}

func timeLocksTranslate(res []byte) (interface{}, error) {
  return res, nil
}
