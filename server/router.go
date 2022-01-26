package server

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/lightninglabs/lndclient"

	"gitlab.i.bitonic.nl/lex/lnurl-server/models"
	"gitlab.i.bitonic.nl/lex/lnurl-server/server/handlers"
)

/*
NewRouter creates a new router and needs LND Services to do so.
*/
func NewRouter(lndServices *lndclient.GrpcLndServices) *mux.Router {
	ctx := context.TODO()
	router := mux.NewRouter().StrictSlash(true)

	// Define the GET requests
	getRoutes := []models.Route{
		{
			Name:        "GetWithdrawLNURL",
			Method:      "GET",
			Pattern:     "/get-withdraw-lnurl",
			HandlerFunc: handlers.GetWithdrawLNURL(ctx),
		},
		{
			Name:        "GetWithdrawDetails",
			Method:      "GET",
			Pattern:     "/get-withdraw-details",
			HandlerFunc: handlers.GetWithdrawDetails(ctx),
		},
		{
			Name:        "WithdrawInvoice",
			Method:      "GET",
			Pattern:     "/withdraw-invoice",
			HandlerFunc: handlers.WithdrawInvoice(ctx, lndServices),
		},
	}

	// Add all GET requests
	for _, route := range getRoutes {
		handler := route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Queries(route.Queries...).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
