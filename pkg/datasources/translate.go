package datasources

import (
  //json "github.com/json-iterator/go"
  json "github.com/json-iterator/go"
  "strconv"
  _ "fmt"
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
  "marketCandlesticks": marketCandlesticksTranslate,
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
//todo
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
  val := make([]map[string]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
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
  val := make([][]interface{}, 0)
  err := json.Unmarshal(res, &val)
  if err!= nil || val== nil || len(val) == 0  {
    return nil, err
  }
  return val, nil
}

func marketDepthTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil {
    return nil, err
  }
  return val, nil
}

func transactionQueryTranslate(res []byte) (interface{}, error) {
  val := make(map[string]interface{})
  err := json.Unmarshal(res, &val)
  if err!= nil || val == nil || val["tx"] == nil {
    return nil, err
  }
  v := []interface{}(val["tx"].([]interface{}))
  return v, nil
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
  v := []interface{}(val["trade"].([]interface{}))
  return v, nil
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
