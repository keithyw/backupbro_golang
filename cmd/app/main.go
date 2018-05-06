package main

import (
	"backupBro/pkg/mongo"
	"log"
	"backupBro/pkg/crypto"
	"backupBro/pkg/server"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27107")
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()
	h := crypto.Hash{}
	u := mongo.NewUserService(ms.Copy(), "backupbro", "user", &h)
	s := server.NewServer(u)
	s.Start()
}
