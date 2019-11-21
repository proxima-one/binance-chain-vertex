package datasources

import (
  "testing"
  "fmt"
  //"log"
  "time"
  helpers "github.com/proxima-one/binance-chain-subgraph/pkg/common"
)
//docker run -d --rm -v /Users/chasesmith/binance-chain-data:/opt/bnbchaind -e "BNET=prod" -p 27146:27146 -p 127.0.0.1:27147:27147 --security-opt no-new-privileges --ulimit nofile=16000:16000 varnav/binance-node



func TestFetch(t *testing.T) {
  baseUri := "http://localhost:8080" 
  //baseUri := "https://dex.binance.org"
  for key, value := range helpers.FetchTestCases {
    fmt.Println(key)
    start := time.Now()
    for i := 0; i < 50; i++ {
    result, err := BinanceRequest(key, baseUri , value)
    if i%1000 == 0 {
      fmt.Println(string(result))
    }
    if err != nil {
      t.Error(err)
    }
  }
  end := time.Now()
  elapsed := end.Sub(start)
  fmt.Println(elapsed)
  fmt.Println(50000/int(elapsed.Milliseconds()))
}
}
