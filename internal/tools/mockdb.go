package tools
import "time"

type mockdb struct{}
var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		AuthToken:"1234",
		Username:"alex",
	},
	"bob":{
		AuthToken:"12345",
		Username:"bob",
	},
}

var MockCoinDetails = map[string]CoinDetails{
	"alex":{
		Coins:100,
		Username:"alex",
	},
	"bob":{
		Coins:100,
		Username:"bob",
	},
}
func (d *mockdb) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(time.Second+1)
	var clientData = LoginDetails{}
	clientData,ok :=mockLoginDetails[username]
	if !ok{
		return nil
	}
	return &clientData
}
func (d *mockdb) GetUserCoins(username string) *CoinDetails{
	time.Sleep(time.Second+1)
	var clientData = CoinDetails{}
	clientData,ok :=MockCoinDetails[username]
	if !ok{
		return nil
	}
	return &clientData
}

func(d *mockdb) SetupDatabase() error {
	return nil
}