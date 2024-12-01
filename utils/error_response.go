package utils

import (
	"errors"
	"koriebruh/management/dto"
	"net/http"
)

var (
	ErrBadRequest          = errors.New("BAD REQUEST")
	ErrInternalServerError = errors.New("INTERNAL SERVER ERROR")
	ErrNotFound            = errors.New("NOT FOUND ERROR")
)

func ErrorResponseWeb(errIs error, err error) dto.WebResponse {
	if errors.Is(errIs, ErrBadRequest) {
		return dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
	} else if errors.Is(errIs, ErrNotFound) {
		return dto.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		}
	} else {
		return dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR sini",
			Data:   err.Error(),
		}
	}
}
