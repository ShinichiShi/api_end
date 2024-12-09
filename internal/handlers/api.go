package handlers
import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/ShinichiShi/api_end/internal/middleware"
)

func Handler(r *chi.Mux){//handles the mux type juz created
	//if the fn is capital it bcms public can be accessed by other files...or else cant
	//fn called bfr actual fn called middleware ish
	// add middleware to route by r.useconst
	// global middleware: applied tto all middlepoints 
	r.Use(chimiddle.StripSlashes) //makes sure trailing slashes are igntored localhost:8000/api/endpoint/ errrr
	
	r.Route("/account",func(router chi.Router){ //anonymous fn added
		 //middleware to check if user is auth firsrt to access data first
		router.Use(middleware.Authorization)

		router.Get("/coins",GetCoinBalance)//middleware auth is called b4 account handler
	})

}

