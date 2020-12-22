package main
import(
	"go_training/todo/controller"
	"net/http"
	"github.com/gorilla/mux"

)

func main()  {
	 api:=mux.NewRouter()
	 
	 api.HandleFunc("/create", controller.Create).Methods(http.MethodGet)
	
	 api.HandleFunc("/create", controller.Create).Methods(http.MethodPost)
	 
	 api.HandleFunc("/readall", controller.ReadAll).Methods(http.MethodGet)
	
	 api.Path("/delete/{id:[0-50]+}").HandlerFunc(controller.Delete).Methods(http.MethodGet)
	
	 api.Path("/update/{id:[0-50]+}").HandlerFunc(controller.Update).Methods(http.MethodGet)	 
	 
	 api.Path("/update/{id:[0-50]+}").HandlerFunc(controller.Update).Methods(http.MethodPost)
	
	 server:=http.Server{
		Addr: ":8080",
		Handler: api,
	}
	server.ListenAndServe()
}