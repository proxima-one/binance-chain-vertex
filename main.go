package main

import (
	//"log"
	//"net/http"
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

const defaultPort = "4000"

func SetupDB(tableList []string, config map[string]string) (*proxima.ProximaDB, error) {
	proximaDB := proxima.NewProximaDB(config["ip"], config["port"])
	_, err := proximaDB.OpenAll(tableList)
	if err != nil {
		return proximaDB, err
	}
	return proximaDB, nil
}

func SetupDatasource(db *proxima.ProximaDB, config map[string]string) (*datasource.Datasource, error) {
	ds, err := datasource.NewDatasource(db, config["uri"])
	if err != nil {
		return nil, err
	}
	return ds, nil
}

func SetDataloader(c *cache.Cache, db *proxima.ProximaDB, ds *datasource.Datasource, config map[string]string) (*dataloader.Dataloader, error) {
	loader , err:= dataloader.NewDataloader(c, db, ds)
	if err != nil {
		return nil, err
	}
	return loader, nil
}

func SetupResolver() (gql.Config) {
	dbConfig := make(map[string]string)
	dbConfig["ip"] = "0.0.0.0"//"db"
	dbConfig["port"] = "50051"
	proximaDB, _ := SetupDB(datasource.BinanceTableList, dbConfig)
	datasourceConfig := make(map[string]string)
	datasourceConfig["uri"] = "https://dex.binance.org"
	ds, _:= SetupDatasource(proximaDB, datasourceConfig)
	go ds.Start()
	c := cache.New(5*time.Minute, 10*time.Minute)
	loader, _  := SetDataloader(c, proximaDB, ds, dbConfig)
	return resolver.NewResolver(loader)
}

func graphqlHandler() gin.HandlerFunc {
		resolver := SetupResolver()
	h := handler.GraphQL(gql.NewExecutableSchema(resolver))

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

func main() {


	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	go r.POST("/query", graphqlHandler())
	//go r.GET("/query", graphqlHandler())
	go r.GET("/", playgroundHandler())
	//fmt.Println("Test with Get      : curl -g 'http://localhost:4000/graphql?query={hello}'")
	r.Run(":4000")


	// handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	// 			rc := graphql.GetResolverContext(ctx)
	// 			fmt.Println("Entered", rc.Object, rc.Field.Name)
	// 			res, err = next(ctx)
	// 			fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
	// 			return res, err
	// 		})

//	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
//	log.Fatal(http.ListenAndServe(":"+port, nil))
}
