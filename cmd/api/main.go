package main
import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/ShinichiShi/api_end/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)//setup logger..so tht u get file and line no.
	var r *chi.Mux = chi.NewRouter()//pointer to mux type....which is struct to handle api
	handlers.Handler(r)
	fmt.Println("Starting go trial api service")
	fmt.Println(`
	|||||||||
	Go apiiii
	|||||||||
	`)
	err:=http.ListenAndServe("localhost:8000",r)//handler which our mux type satisfies
	if err!=nil{
			log.Error(err)
	}
}