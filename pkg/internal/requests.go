
package datasources
//
// import (
//     "io/ioutil"
//     "net/http"
//     "fmt"
//   )
//
// func binance_request(requestType string, baseUri string, args map[string]interface{}) ([]byte, error) {
//   uri := binance_datasource_uri[requestType](baseUri, args)
//   resp, err := http.Get(uri)
//   if err!=nil {
//     return nil, err
//   }
//   body, e := ioutil.ReadAll(resp.Body)
//   if e != nil {
//     return nil, e
//   }
//   return body, nil
// }
//
//
// func (ds *Datasource) updateEntity(name, table, ) {
//
// }
//
//
// func (ds *Datasource) updateFees() {
//   args := make(map[string]interface{});
//   fees := ds.dataRequest("fees", args); //errs
//   ds.proximaDB.Set(Primary, "Fees", fees, args)
// }
//
// func (ds *Datasource) updateValidators() {
//   args := make(map[string]interface{});
//   validators := ds.dataRequest("validators", args) //errs
//   ds.proximaDB.Set(Primary, "Validators", validators, args)
// }
//
// func (ds *Datasource) updateTokens() {
//   args := make(map[string]interface{});
//   tokens := ds.dataRequest("tokens", args) //errs
//   ds.proximaDB.Set(Primary, "Tokens", tokens, args)
// }
//
// func (ds *Datasource) updateMarkets() {
//   args := make(map[string]interface{});
//   markets := ds.dataRequest("markets", args) //errs
//   ds.proximaDB.Set(Primary, "Markets", markets, args)
// }
//
// func (ds *Datasource) updateBlockStats() (map[string]interface{}) {
//   args := make(map[string]interface{});
//   blockStats := ds.dataRequest("blockStats", args).(map[string]interface{})
//   ds.proximaDB.Set(Primary, "BlockStats", blockStats, args)
//   return blockStats
// }
//
//
//
// func (ds *Datasource) updateMarketTickers() (interface{}) {
//   args := make(map[string]interface{});
//   marketTickers := ds.dataRequest("marketTickers", args).([]map[string]interface{})
//   ds.proximaDB.Set(Primary, "MarketTickers", marketTickers, args)
//   return marketTickers
// }
//
// func (ds *Datasource) updateMarketDepth(symbol string) {
//   args:= map[string]interface{}{"symbol": symbol,}
//   data := ds.dataRequest("marketDepth", args)
//   ds.proximaDB.Set(MarketDepthBySymbol, symbol, data, args)
// }
//
//
//
// func (ds *Datasource) updateAccounts(accounts map[string]bool) {
//   dbArgs := make(map[string]interface{});
//   account := make(map[string]interface{});
//   addressKey := "address"
//   for address, _ := range accounts {
//     requestArgs := map[string]interface{}{
//       addressKey : address,
//     }
//     account = ds.dataRequest("account", requestArgs).(map[string]interface{})
//     ds.proximaDB.Set(AccountsByAddress, address, account, dbArgs)
//   }
// }
