package diagnostics

import(
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func NewDiagnostics() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz)
	router.HandleFunc("/info", ready)
    return router;
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}