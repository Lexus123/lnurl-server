package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lightninglabs/lndclient"
)

/*
Host will init the routes and fire up a server
*/
func Host() {
	// Setup the config for an LND Services client
	conf := &lndclient.LndServicesConfig{
		LndAddress:  fmt.Sprintf("%s:%s", os.Getenv("LND_HOST"), os.Getenv("LND_RPC_PORT")),
		Network:     lndclient.NetworkMainnet,
		MacaroonDir: os.Getenv("MACAROON_DIR"),
		TLSPath:     os.Getenv("TLS_CERT_PATH"),
	}

	// Pass the config and get a LND Services client
	lndServices, err := lndclient.NewLndServices(conf)
	if err != nil {
		log.Panicf("Host (NewLndServices) error: %s", err.Error())
	} else {
		log.Println("Connected with LND")
	}

	// Setup routes and make lndclient available to them
	router := NewRouter(lndServices)

	log.Print("Server running on port 8003")
	log.Fatal(http.ListenAndServe(":8003", router))
}
