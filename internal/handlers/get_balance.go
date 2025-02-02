package handlers

import (
	"Go-Tutorial-Api/api"
	"Go-Tutorial-Api/internal/tools"
	"encoding/json"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.ApiParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var database *tools.DatabaseInterface
	database, errorMsg := tools.NewDatabase()
	if errorMsg != nil {
		log.Error(errorMsg)
		api.InternalErrorHandler(w)
		return
	}
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoinDetails(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var response = api.ApiResponse{
		StatusCode: http.StatusOK,
		Balance:    (*tokenDetails).Coins,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
