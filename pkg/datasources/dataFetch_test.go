package datasources

import (
  "testing"
  "fmt"
  "log"
  helpers "github.com/proxima-one/binance-chain-subgraph/pkg/common"
)




func TestFetch(t *testing.T) {
  baseUri := "https://dex.binance.org"
  fmt.Println(helpers.FetchTestCases)
  for key, value := range helpers.FetchTestCases {
    log.Println(key, value)
    result, err := BinanceRequest(key, baseUri , value)
    log.Println(string(result))
    if err != nil {
      t.Error("Cannot batch values: ", err)
    }
    val, tErr := BinanceTranslate(key, result)
    if tErr != nil {
      t.Error("Cannot batch values: ", tErr)
    }
    log.Println(val)
  }
}
