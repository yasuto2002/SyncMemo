package handler

import (
	"fmt"
	"net/http"
	"syncmemo/auth"
	"syncmemo/repository/response"

	"github.com/go-playground/validator/v10"
)

type LoginCheck struct {
	Validator *validator.Validate
}

func (lc *LoginCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mail, ok := auth.GetUserID(ctx)
	if !ok {
		RespondJSON(ctx, rw, &ErrResponse{
			Message: fmt.Errorf("user_id not found").Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, rw, response.LoginCheck{Mail: mail}, http.StatusOK)
}
