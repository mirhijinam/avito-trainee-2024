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

func (h *Handler) responseCreator(w http.ResponseWriter, r *http.Request, status int, env map[string]interface{}) {
	err := writeJSONBody(w, status, env, nil)
	if err != nil {
		h.responseLog(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) successBannerCreationResponse(w http.ResponseWriter, r *http.Request) {
	ans := "answer"
	msg := "error! user is unauthorized"
	env := envelope{ans: msg}
	h.responseCreator(w, r, http.StatusOK, env)
}

func (h *Handler) successGetBannerListResponse(w http.ResponseWriter, r *http.Request, env map[string]interface{}) {
	h.responseCreator(w, r, http.StatusOK, env)
}

func (h *Handler) userUnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	ans := "answer"
	msg := "error! user is unauthorized"
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusUnauthorized, env)
}

func (h *Handler) forbiddenAccessResponse(w http.ResponseWriter, r *http.Request) {
	ans := "answer"
	msg := "error! user has no access"
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusForbidden, env)
}

func (h *Handler) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.responseLog(r, err)
	ans := "answer"
	msg := "error! the server encountered a problem and could not process your request"
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusNotFound, env)
}

func (h *Handler) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	ans := "answer"
	msg := "error! the requested resource could not be found"
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusNotFound, env)
}

func (h *Handler) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	ans := "answer"
	msg := fmt.Sprintf("error! the %s method is not supported for this resource", r.Method)
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusMethodNotAllowed, env)
}

func (h *Handler) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ans := "answer"
	msg := "error! " + err.Error()
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusBadRequest, env)
}

func (h *Handler) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	ans := "answer"
	msg := errors
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusUnprocessableEntity, env)
}

func (h *Handler) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	ans := "answer"
	msg := "error! unable to update the record due to an edit conflict, please try again"
	env := envelope{ans: msg}

	h.responseCreator(w, r, http.StatusConflict, env)
}
