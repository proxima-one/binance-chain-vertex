package dataloader

import (
 "time"
 cache "github.com/patrickmn/go-cache"
)


const TokensCacheExpiration = cache.NoExpiration

const FeesCacheExpiration = cache.NoExpiration

const MarketsCacheExpiration = cache.NoExpiration

const ValidatorsCacheExpiration = cache.NoExpiration

const TransactionsCacheExpiration = 5*time.Minute

const OrdersCacheExpiration = 5*time.Minute

const TradesCacheExpiration = 5*time.Minute

const MarketCandleSticksCacheExpiration = 1*time.Minute

const BlockStatsCacheExpiration = 5*time.Second

const MarketTickersCacheExpiration = 5*time.Second

const MarketTickerCacheExpiration = 5*time.Second

const MarketDepthCacheExpiration = 5*time.Second
