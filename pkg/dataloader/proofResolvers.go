package dataloader

import (
  proxima "github.com/proxima-one/proxima-db-client-go"
  models "github.com/proxima-one/binance-chain-subgraph/pkg/models"
  "fmt"

)

func GenerateProof(rawProof *proxima.ProximaDBProof) (models.Proof){
  proof := string(rawProof.GetProof())
  root :=  fmt.Sprintf("%x\n", rawProof.GetRoot());
  return models.Proof{Proof: &proof,  Root: &root}
}
