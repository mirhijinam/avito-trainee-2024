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

func (h *Handler) responseLog(r *http.Request, err error) {
	log.Println(err)
}

func (h *Handler) responseCreator(w http.ResponseWriter, r *http.Request, status int, msg interface{}) {
	env := envelope{"answer": msg}

	err := writeJSON(w, status, env, nil)
	if err != nil {
		h.responseLog(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) successResponse(w http.ResponseWriter, r *http.Request) {
	msg := "banner was created successfully"
	h.responseCreator(w, r, http.StatusOK, msg)
}

func (h *Handler) userUnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	msg := "error! user is unauthorized"
	h.responseCreator(w, r, http.StatusUnauthorized, msg)
}

func (h *Handler) forbiddenAccessResponse(w http.ResponseWriter, r *http.Request) {
	msg := "error! user has no access"
	h.responseCreator(w, r, http.StatusForbidden, msg)
}

func (h *Handler) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.responseLog(r, err)

	msg := "error! the server encountered a problem and could not process your request"
	h.responseCreator(w, r, http.StatusNotFound, msg)
}

func (h *Handler) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	msg := "error! the requested resource could not be found"
	h.responseCreator(w, r, http.StatusNotFound, msg)
}

func (h *Handler) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("error! the %s method is not supported for this resource", r.Method)
	h.responseCreator(w, r, http.StatusMethodNotAllowed, msg)
}

func (h *Handler) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.responseCreator(w, r, http.StatusBadRequest, "error! "+err.Error())
}

func (h *Handler) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	h.responseCreator(w, r, http.StatusUnprocessableEntity, errors)
}

func (h *Handler) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	msg := "error! unable to update the record due to an edit conflict, please try again"
	h.responseCreator(w, r, http.StatusConflict, msg)
}
