package http

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/rusneustroevkz/http-server/internal/errs"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	"net/http"
)

type ProductHTTPHandler struct {
	log logger.Logger
}

func NewProductsHTTPHandler(log logger.Logger) *ProductHTTPHandler {
	return &ProductHTTPHandler{log: log}
}

func (*ProductHTTPHandler) Pattern() string {
	return "/products"
}

func (h *ProductHTTPHandler) Routes() *chi.Mux {
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
// @Description  Simple products get
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /products/{id} [get]
func (h *ProductHTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get"))
}

// List godoc
// @Summary      List handler
// @Description  Simple products list
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /products [get]
func (h *ProductHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}

// Delete godoc
// @Summary      Delete handler
// @Description  Simple products list
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /products [delete]
func (h *ProductHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

// Save godoc
// @Summary      Save handler
// @Description  Simple products put
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /products [put]
func (h *ProductHTTPHandler) Save(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("put"))
}

// Create godoc
// @Summary      Create handler
// @Description  Simple products post
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /products [post]
func (h *ProductHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post"))
}

// Update godoc
// @Summary      Update handler
// @Description  Simple products patch
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}  errs.Error
// @Failure      404  {object}  errs.Error
// @Failure      500  {object}  errs.Error
// @Router       /products [patch]
func (h *ProductHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("patch"))
}
