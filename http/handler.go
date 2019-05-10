package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/threeaccents/botnet"
)

//Handler is
type Handler struct {
	*mux.Router
	CommanderService botnet.Commander
	Storage          botnet.Storager
}

//NewHandler is
func NewHandler(cc botnet.Commander, storage botnet.Storager) *Handler {
	h := &Handler{
		Router:           mux.NewRouter(),
		CommanderService: cc,
		Storage:          storage,
	}

	const listBotsPath = "/bots"
	h.Handle(listBotsPath, h.handleListBots()).Methods("GET")

	const checkBotsHealthPath = "/bots/health"
	h.Handle(checkBotsHealthPath, h.handleCheckBotsHealth()).Methods("POST")

	const commandPath = "/commands"
	h.Handle(commandPath, h.handleCommand()).Methods("POST")

	// const listProfessionalAppointmentsPath = "/v1/professionals/{professional_id}/appointments"
	// h.Handle(listProfessionalAppointmentsPath, h.authMiddleware(h.handleListProfessionalAppointments())).Methods("GET")

	return h
}

// ServeHTTP is
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}
