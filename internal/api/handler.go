package api

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	BannerService
}

func New(bs BannerService) *Handler {
	return &Handler{
		BannerService: bs,
	}
}
func (h *Handler) successResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) logError(r *http.Request, err error) {
	log.Println(err)
}

func (h *Handler) errorResponse(w http.ResponseWriter, r *http.Request, status int, msg interface{}) {
	env := envelope{"error": msg}

	err := writeJSON(w, status, env, nil)
	if err != nil {
		h.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) userUnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	msg := "user is unauthorized"
	h.errorResponse(w, r, http.StatusUnauthorized, msg)
}

func (h *Handler) forbiddenAccessResponse(w http.ResponseWriter, r *http.Request) {
	msg := "user has no access"
	h.errorResponse(w, r, http.StatusForbidden, msg)
}

func (h *Handler) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logError(r, err)

	msg := "the server encountered a problem and could not process your request"
	h.errorResponse(w, r, http.StatusNotFound, msg)
}

func (h *Handler) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	msg := "the requested resource could not be found"
	h.errorResponse(w, r, http.StatusNotFound, msg)
}

func (h *Handler) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	h.errorResponse(w, r, http.StatusMethodNotAllowed, msg)
}

func (h *Handler) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (h *Handler) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	h.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (h *Handler) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	msg := "unable to update the record due to an edit conflict, please try again"
	h.errorResponse(w, r, http.StatusConflict, msg)
}
