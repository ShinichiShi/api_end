package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ShinichiShi/lessGO.git/api"
	"github.com/ShinichiShi/lessGO.git/internal/tools"
	"github.com/gorilla.schema"
	"github.com/shinichishi/api_end/api"
	"github.com/shinichishi/api_end/internal/tools"
	log "github.com/sirupsen/logrus"
)
func GetCoinBalance( w http.ResponseWriter, r *http.Request){
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.newDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var database *tools.DatabaseInterface
	database,err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var tokenDetails *tools.CoinDetails
	tokenDetails= (*database).GetUserCoins(params.Username)
	if tokenDetails == nil{
		log.Error((err))
		api.InternalErrorHandler(w)
		return
	}
	var response = api.CoinBalanceResponse{
		Balance: (*&tokenDetails.Coins),
		Code: http.StatusOK,
	}
	w.Header().Set("Content-Type","application/json")
	err=json.NewEncoder(w).Encode(response)
	if err !=nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}