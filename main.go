package main

import (
	"github.com/gin-gonic/gin"
	"os"
	_ "context"
	_ "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	proxima "github.com/proxima-one/proxima-db-client-go"
	resolver "github.com/proxima-one/binance-chain-subgraph/pkg/resolvers"
	datasource "github.com/proxima-one/binance-chain-subgraph/pkg/datasources"
	dataloader "github.com/proxima-one/binance-chain-subgraph/pkg/dataloader"
  gql "github.com/proxima-one/binance-chain-subgraph/pkg/gql"
	cache "github.com/patrickmn/go-cache"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	go r.POST("/query", BinanceSubgraphServer())
	go r.GET("/", playgroundHandler())
	r.Run(":4000")
}

func BinanceSubgraphServer() gin.HandlerFunc {
	db, _ := StartDatabase(datasource.BinanceTableList)
	ds, _ := StartBinanceDatasource(db)
	subgraphResolvers := CreateResolvers(db, ds)
		h := handler.GraphQL(gql.NewExecutableSchema(subgraphResolvers))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CreateResolvers(db *proxima.ProximaDB, ds *datasource.Datasource) (gql.Config) {
	c := cache.New(5*time.Minute, 10*time.Minute)
	loader, _  := SetDataloader(c, db, ds)
	return resolver.NewResolver(loader)
}

func StartDatabase(tableList []string) (*proxima.ProximaDB, error) {
	ip := getEnv("DB_ADDRESS" , "0.0.0.0")
	port :=  getEnv("DB_PORT", "50051")
	proximaDB := proxima.NewProximaDB(ip, port)
	_, err := proximaDB.OpenAll(tableList)
	if err != nil {
		return proximaDB, err
	}
	return proximaDB, nil
}

func StartBinanceDatasource(db *proxima.ProximaDB) (*datasource.Datasource, error) {
	ip := getEnv("BINANCE_NODE_ADDRESS" , "dex.binance.org")
	port :=  getEnv("BINANCE_NODE_PORT", "")
	uri := "https://"
	uri = uri + ip
	if len(port) > 0 {
		uri = "http://" + ip + ":" + port
	}
	ds, err := datasource.NewDatasource(db, uri)
	if err != nil {
		return nil, err
	}
	go ds.Start()
	return ds, nil
}

func SetDataloader(c *cache.Cache, db *proxima.ProximaDB, ds *datasource.Datasource) (*dataloader.Dataloader, error) {
	loader , err:= dataloader.NewDataloader(c, db, ds)
	if err != nil {
		return nil, err
	}
	return loader, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
	  return value
  }
    return defaultVal
}
