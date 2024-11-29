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

	ErrPhoneNumberAlreadyRegistered    = errors.New("phone number already registered")
	ErrInvalidOtp                      = errors.New("invalid otp")
	ErrOtpNotFound                     = errors.New("otp not found")
	ErrPasswordConfirmationMissmatch   = errors.New("password confirmation missmatch")
	ErrInvalidPhoneOrPassword          = errors.New("invalid phone or password")
	ErrInvalidOldPassword              = errors.New("invalid old password")
	ErrMissingGenderParams             = errors.New("missing gender params")
	ErrInsuficientMembershipAccess     = errors.New("insufficient membership access")
	ErrLoginRequired                   = errors.New("login required")
	ErrInvalidOrderStatus              = errors.New("invalid order status")
	ErrOrderBenefitOnprocess           = errors.New("order benefit onprocess")
	ErrInvalidDatingPlanOwner          = errors.New("invalid dating plan owner")
	ErrDatingFeedbackExist             = errors.New("dating feedback exist")
	ErrWaitngForMarriageProposalAnswer = errors.New("waitng for marriage proposal answer")
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

		ErrPhoneNumberAlreadyRegistered: {
			HttpCode: 422, ErrorCode: "phone_number_already_registered",
			ID: "Nomor telephone sudah terdaftar", EN: "Phone already registered",
		},
		ErrInvalidOtp: {
			HttpCode: 422, ErrorCode: "invalid_otp",
			ID: "Otp tidak benar", EN: "Invalid otp",
		},
		ErrOtpNotFound: {
			HttpCode: 422, ErrorCode: "otp_not_found",
			ID: "Otp tidak benar", EN: "Invalid otp",
		},
		ErrPasswordConfirmationMissmatch: {
			HttpCode: 400, ErrorCode: "password_confirmation_missmatch",
			ID: "Konfirmasi password tidak cocok", EN: "Password confirmation missmatch",
		},
		ErrInvalidPhoneOrPassword: {
			HttpCode: 401, ErrorCode: "invalid_phone_or_password",
			ID: "Nomor telpon atau password salah", EN: "Invalid phone or password",
		},
		ErrInvalidOldPassword: {
			HttpCode: 400, ErrorCode: "invalid_old_password",
			ID: "Password lama anda salah", EN: "Invalid old password",
		},
		ErrMissingGenderParams: {
			HttpCode: 400, ErrorCode: "missing_gender_params",
			ID: "Jenis kelamin belum diatur", EN: "Missing gender params",
		},
		ErrInsuficientMembershipAccess: {
			HttpCode: 403, ErrorCode: "insufficient_membership_access",
			ID: "Akses keanggotaan tidak cukup", EN: "Insufficient membership akses",
		},
		ErrLoginRequired: {
			HttpCode: 403, ErrorCode: "login_required",
			ID: "Membutuhkan login", EN: "Login required",
		},
		ErrInvalidOrderStatus: {
			HttpCode: 422, ErrorCode: "invalid_order_status",
			ID: "Status order tidak benar", EN: "Invalid order status",
		},
		ErrOrderBenefitOnprocess: {
			HttpCode: 422, ErrorCode: "order_benefit_onprocess",
			ID: "Order sedang dalam proses", EN: "Order benefit onprocess",
		},
		ErrInvalidDatingPlanOwner: {
			HttpCode: 422, ErrorCode: "invalid_dating_plan_owner",
			ID: "Pemilik dating plan tidak benar", EN: "Invalid dating plan owner",
		},
		ErrDatingFeedbackExist: {
			HttpCode: 422, ErrorCode: "dating_feedback_exist",
			ID: "Masukan dating plan sudah ada", EN: "Dating feedback exist",
		},
		ErrWaitngForMarriageProposalAnswer: {
			HttpCode: 422, ErrorCode: "waitng_for_marriage_proposal_answer",
			ID: "Menunggu jawaban proposal pernikahan", EN: "Waitng for marriage proposal answer",
		},
	}
)
