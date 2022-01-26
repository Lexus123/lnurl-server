package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fiatjaf/go-lnurl"
	"github.com/lightninglabs/lndclient"

	"gitlab.i.bitonic.nl/lex/lnurl-server/models"
)

/*
GetWithdrawLNURL handles requests made to /get-withdraw-lnurl
*/
func GetWithdrawLNURL(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// The URL where the wallet can retrieve withdraw details
		url := fmt.Sprintf("http://%s/get-withdraw-details", os.Getenv("LNURL_SERVER_HOST"))

		encodedUrl, _ := lnurl.LNURLEncode(url)

		// Create the response
		response := models.NewWithdrawQR(encodedUrl)

		output, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}

	return http.HandlerFunc(fn)
}

/*
GetWithdrawDetails handles requests made to /get-withdraw-details
*/
func GetWithdrawDetails(ctx context.Context) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Create the response
		response := models.NewWithdrawResponse("username", "withdraw")

		output, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}

	return http.HandlerFunc(fn)
}

/*
WithdrawInvoice handles requests made to /withdraw-invoice
*/
func WithdrawInvoice(ctx context.Context, lndServices *lndclient.GrpcLndServices) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Get query param "payment request" pr
		invoice := r.FormValue("pr")

		// Create the response
		response := models.NewStatusResponse("OK", "")

		lndServices.Client.PayInvoice(ctx, invoice, 100, nil)

		output, _ := json.Marshal(response)

		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}

	return http.HandlerFunc(fn)
}
