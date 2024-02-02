package handlers

import (
	_ "asd/internal/errs"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func (h *EchoHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", h.List)
		r.Get("/{id}", h.Get)
		r.Delete("/", h.Delete)
		r.Put("/", h.Save)
		r.Patch("/", h.Update)
		r.Post("/", h.Create)
	})

	return r
}

// Get godoc
// @Summary      Get handler
// @Description  Simple echo get
// @Tags         echo
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /echo/{id} [get]
func (h *EchoHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get"))
}

// List godoc
// @Summary      List handler
// @Description  Simple echo list
// @Tags         echo
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /echo [get]
func (h *EchoHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}

// Delete godoc
// @Summary      Delete handler
// @Description  Simple echo list
// @Tags         echo
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /echo [delete]
func (h *EchoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

// Save godoc
// @Summary      Save handler
// @Description  Simple echo put
// @Tags         echo
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /echo [put]
func (h *EchoHandler) Save(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put"))
}

// Create godoc
// @Summary      Create handler
// @Description  Simple echo post
// @Tags         echo
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /echo [post]
func (h *EchoHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post"))
}

// Update godoc
// @Summary      Update handler
// @Description  Simple echo patch
// @Tags         echo
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /echo [patch]
func (h *EchoHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("patch"))
}
