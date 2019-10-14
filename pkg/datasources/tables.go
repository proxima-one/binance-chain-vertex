package datasources

var Primary string = "Binance-Primary"
var MarketDepthBySymbol string = "Binance-MarketDepthBySymbol"
var MarketTickersBySymbol string =  "Binance-MarketTickersBySymbol"
var AccountsByAddress string = "Binance-AccountsByAddress"
var OrdersByOrderId string = "Binance-OrdersByOrderId"
var OrdersByOwner string = "Binance-OrdersByOwner"
var TransactionsByTxHash string = "Binance-TransactionsByTxHash"
var TransactionsByFromAddr string = "Binance-TransactionsByFromAddr"
var TransactionsByToAddr string = "Binance-TransactionsByToAddr"
var TransactionsByBlockHeight string = "Binance-TransactionsByBlockHeight"
var TransactionsByTimeStamp string = "Binance-ransactionsByTimeStamp"
var TradesByTime string = "Binance-TradesByTime"
var TradesByTradeId string = "Binance-TradesByTradeId"
var TradesByBuyerId string = "Binance-TradesByBuyerId"
var TradesBySellerId string = "Binance-TradesBySellerId"
var AtomicSwapsBySwapId string = "Binance-AtomicSwapsBySwapId"
var AtomicSwapsByFromAddr string = "Binance-AtomicSwapsByFromAddr"
var AtomicSwapsByToAddr string = "Binance-AtomicSwapsByToAddr"
var AtomicSwapsByTimestamp string = "Binance-AtomicSwapsByTimestamp"

var BinanceTableList []string = []string{
  Primary,
  MarketDepthBySymbol,
  MarketTickersBySymbol,
  AccountsByAddress,
  OrdersByOrderId,
  OrdersByOwner,
  TransactionsByTxHash,
  TransactionsByToAddr,
  TransactionsByFromAddr,
  TransactionsByTimeStamp,
  TradesByTradeId,
  TradesByBuyerId,
  TradesBySellerId,
  TradesByTime,
  AtomicSwapsBySwapId,
  AtomicSwapsByToAddr,
  AtomicSwapsByFromAddr,
  AtomicSwapsByTimestamp,
}
