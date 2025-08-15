package model

import (
	"errors"
)

type JxErr struct {
	HttpCode  int
	ErrorCode string
	ID        string
	EN        string
}

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrForbidden           = errors.New("forbidden")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrBadRequest          = errors.New("bad request")

	ErrLoginRequired        = errors.New("login required")
	ErrSubscriptionRequired = errors.New("subscription required")
)

var (
	ERR_MAP = map[error]JxErr{
		ErrInternalServerError: {
			HttpCode: 500, ErrorCode: "internal_server_error", ID: "Kesalahan internal server", EN: "Internal server error",
		},
		ErrUnprocessableEntity: {
			HttpCode: 422, ErrorCode: "unprocessable_entity", ID: "Tidak dapat diproses", EN: "Unprocessable entity",
		},
		ErrForbidden: {
			HttpCode: 403, ErrorCode: "forbidden", ID: "Tidak dapat diakses", EN: "Forbidden access",
		},
		ErrUnauthorized: {
			HttpCode: 401, ErrorCode: "unauthorized", ID: "Tidak ter autorisasi", EN: "Unauthorized",
		},
		ErrTooManyRequests: {
			HttpCode: 429, ErrorCode: "too_many_requests", ID: "Terlalu banyak request", EN: "Too many requests",
		},
		ErrBadRequest: {
			HttpCode: 400, ErrorCode: "bad_request", ID: "Permintaan buruk", EN: "Bad request",
		},

		ErrLoginRequired: {
			HttpCode: 400, ErrorCode: "login_required",
			ID: "Kamu harus login terlebih dahulu", EN: "You have to login first",
		},
		ErrSubscriptionRequired: {
			HttpCode: 400, ErrorCode: "subscription_required",
			ID: "Kamu harus membeli subscription terlebih dahulu", EN: "You have to purchase subscription first",
		},
	}
)
