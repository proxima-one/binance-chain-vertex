package datasources

import (
  "testing"
  "fmt"
  //"log"
  "time"
  helpers "github.com/proxima-one/binance-chain-subgraph/pkg/common"
)




func TestFetch(t *testing.T) {
  baseUri := "https://dex.binance.org"
  for key, value := range helpers.FetchTestCases {
    fmt.Println(key)
    start := time.Now()
    for i := 0; i < 5; i++ {
    _, err := BinanceRequest(key, baseUri , value)
    if err != nil {
      t.Error(err)
    }
  }
  end := time.Now()
  elapsed := end.Sub(start)
  fmt.Println(elapsed)
  fmt.Println(elapsed/5)
}
}
