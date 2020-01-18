package fixer

import (
	"encoding/json"
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/parnurzeal/gorequest"
)

var cache Currencies
var fixerApiKey string

func SetKey(apikey string) {
	fixerApiKey = apikey
}

func Update() bool {
	URL := "http://data.fixer.io/api/latest?access_key=" + fixerApiKey + "&format=1"
	resp, _, _ := gorequest.New().Get(URL).End()
	var newCache Currencies
	err := json.NewDecoder(resp.Body).Decode(&newCache)
	if err != nil {
		sentry.CaptureException(err)
	}
	if newCache.Success {
		cache = newCache
		return true
	}
	sentry.CaptureException(errors.New("JSON response isn't success"))
	return false
}

func getCurrencies() interface{} {
	return cache
}
