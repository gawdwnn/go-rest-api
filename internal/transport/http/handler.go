package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comment service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetUpRoutes - sets up all routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "i am alive!")
	})
}
