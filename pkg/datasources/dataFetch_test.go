package datasources

import (
  "testing"
  "fmt"
  helpers "github.com/proxima-one/binance-chain-subgraph/pkg/common"
)




func TestFetch(t *testing.T) {
  baseUri := "https://dex.binance.org"
  for key, value := range helpers.FetchTestCases {
    //fmt.Println(key, value)
    if err != nil {
      t.Error("Cannot batch values: ", err)
    }
  }
}
