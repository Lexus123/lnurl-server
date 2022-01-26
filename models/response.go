package models

import (
	"fmt"
	"os"

	lnurl "github.com/fiatjaf/go-lnurl"
)

func NewWithdrawResponse(k1, description string) lnurl.LNURLWithdrawResponse {
	return lnurl.LNURLWithdrawResponse{
		Tag:                "withdrawRequest",
		K1:                 k1,
		Callback:           fmt.Sprintf("http://%s/withdraw-invoice", os.Getenv("LNURL_SERVER_HOST")),
		MaxWithdrawable:    3000000,
		MinWithdrawable:    3000000,
		DefaultDescription: description,
	}
}

func NewStatusResponse(status, reason string) lnurl.LNURLResponse {
	return lnurl.LNURLResponse{
		Status: status,
		Reason: reason,
	}
}

type WithdrawQR struct {
	LightningLNURL string `json:"lightningLnurl"` // For QR generation
	RawLNURL       string `json:"rawLnurl"`       // Raw LNURL for copy-paste
}

func NewWithdrawQR(encodedUrl string) WithdrawQR {
	return WithdrawQR{
		LightningLNURL: fmt.Sprintf("lightning:%s", encodedUrl),
		RawLNURL:       encodedUrl,
	}
}
