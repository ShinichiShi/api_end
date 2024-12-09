package api
import (
	"encoding/json"
	"net/http"
)
//coin balance params
type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct{
	//success code 200	
	Code int

	Balance int64
}

//error struct
type Error struct {
	code int
	Message string
}
