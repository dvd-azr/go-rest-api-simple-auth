package controllers

import (
	"net/http"
	"rest-api-simple-auth/helpers"
)

func Pong(w http.ResponseWriter, r *http.Request) {

	// helpers.SuccessResponse
	w.WriteHeader(http.StatusOK)
	helpers.SuccessJsonResponse(w, []struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}{
		{1, "a@mail.com"},
		{2, "b@mail.com"},
	})

}
