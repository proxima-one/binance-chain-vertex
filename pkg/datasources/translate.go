package datasources

import (
  "encoding/json"
  "strconv"
  _ "fmt"
)

func binance_translate(requestType string, res []byte) (interface{}) {
  return TRANSLATE_RES[requestType](res)
}

var TRANSLATE_RES = map[string]func([]byte) (interface{}) {
  "fees": feesTranslate,
  "tokens": tokensTranslate,
  "blockStats": blockStatsTranslate,
  "markets": marketsTranslate,
  "validators":validatorsTranslate,
  "marketDepth":marketDepthTranslate,
  "marketTickers":marketTickersTranslate,
  "marketCandlesticks":  marketCandlesticksTranslate,
  "account": accountTranslate,
  "transaction": transactionTranslate,
  "order": orderTranslate,
  "trade": tradeTranslate,
  "atomicSwap": atomicSwapTranslate,
}
//todo
func feesTranslate(res []byte) (interface{}) {
  fees:= make([]map[string]interface{}, 0)
  json.Unmarshal(res, &fees)
  for _ , fee := range fees {
    if fee["msg_type"] == nil && fee["dex_fee_fields"] != nil {
      fee["msg_type"] = "dex_fee_fields"
    }

    if fee["msg_type"] == "" && fee["fixed_fee_fields"] != nil {
     fee["msg_type"] = "fixed_fee_fields"
    }
  }
  return fees
}

func blockStatsTranslate(res []byte) (interface{}) {
  val := make(map[string]interface{})
  json.Unmarshal(res, &val)
  sync_val := val["sync_info"]
  if sync_val == nil {
      sync_val = make(map[string]interface{})
      return sync_val
  } else {
    v := sync_val.(map[string]interface{})
    temp:= int64(v["latest_block_height"].(float64))
    v["latest_block_height"] = strconv.FormatInt(temp, 10)
    return v
  }

}

func validatorsTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}

func tokensTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}

func marketsTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}

func marketTickersTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}

func marketTickerTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}

func marketCandlesticksTranslate(res []byte) (interface{}) {
  val := make(map[string]interface{})
  json.Unmarshal(res, &val)
  return val
}

func marketDepthTranslate(res []byte) (interface{}) {
  val := make(map[string]interface{})
  json.Unmarshal(res, &val)
  return val
}

func accountTranslate(res []byte) (interface{}) {
  val := make(map[string]interface{})
  json.Unmarshal(res, &val)
  return val
}

func tradeTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}

func orderTranslate(res []byte) (interface{}) {
  return res
}

func transactionTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}

func atomicSwapTranslate(res []byte) (interface{}) {
  val := make([]map[string]interface{}, 0)
  json.Unmarshal(res, &val)
  return val
}
