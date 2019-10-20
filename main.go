package main

import (
	"log"
	"net/http"
	"os"
	_ "context"
	_ "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	proxima "github.com/proxima-one/proxima-db-client-go"
	binance_chain_resolvers "github.com/proxima-one/binance-chain-subgraph/pkg/resolvers"
	datasource "github.com/proxima-one/binance-chain-subgraph/pkg/datasources"
  gql "github.com/proxima-one/binance-chain-subgraph/pkg/gql"
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

func main() {
	dbConfig := make(map[string]string)
	dbConfig["ip"] = "0.0.0.0"//"db"
	dbConfig["port"] = "50051"
	proximaDB, _ := SetupDB(datasource.BinanceTableList, dbConfig)
	datasourceConfig := make(map[string]string)
	datasourceConfig["uri"] = "https://dex.binance.org"
	ds, _:= SetupDatasource(proximaDB, datasourceConfig)
	go ds.Start()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(binance_chain_resolvers.NewResolver(proximaDB))))



	// handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	// 			rc := graphql.GetResolverContext(ctx)
	// 			fmt.Println("Entered", rc.Object, rc.Field.Name)
	// 			res, err = next(ctx)
	// 			fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
	// 			return res, err
	// 		})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
