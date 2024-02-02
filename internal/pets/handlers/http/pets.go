package http

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/rusneustroevkz/http-server/internal/errs"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	"net/http"
)

type PetsHTTPHandler struct {
	log logger.Logger
}

func NewPetsHTTPHandler(log logger.Logger) *PetsHTTPHandler {
	return &PetsHTTPHandler{log: log}
}

func (*PetsHTTPHandler) Pattern() string {
	return "/pets"
}

func (h *PetsHTTPHandler) Routes() *chi.Mux {
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
// @Description  Simple pets get
// @Tags         pets
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /pets/{id} [get]
func (h *PetsHTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get"))
}

// List godoc
// @Summary      List handler
// @Description  Simple pets list
// @Tags         pets
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /pets [get]
func (h *PetsHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}

// Delete godoc
// @Summary      Delete handler
// @Description  Simple pets list
// @Tags         pets
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /pets [delete]
func (h *PetsHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

// Save godoc
// @Summary      Save handler
// @Description  Simple pets put
// @Tags         pets
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /pets [put]
func (h *PetsHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put"))
}

// Create godoc
// @Summary      Create handler
// @Description  Simple pets post
// @Tags         pets
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /pets [post]
func (h *PetsHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post"))
}

// Update godoc
// @Summary      Update handler
// @Description  Simple pets patch
// @Tags         pets
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /pets [patch]
func (h *PetsHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("patch"))
}
