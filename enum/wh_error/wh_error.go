package wh_error

// 定義 error 對應成功訊息

var (
	// 基本錯誤
	BadRequestError      = &WhiteLabelError{errorMessage: "Bad Request", code: BAD_REQUEST}
	InvalidPayloadError  = &WhiteLabelError{errorMessage: "Invalid Payload", code: INVALID_PAYLOAD}
	RecordNotExistsError = &WhiteLabelError{errorMessage: "Record Not Exists", code: RECORD_NOT_EXISTS}
	InvalidTokenError    = &WhiteLabelError{errorMessage: "Invalid Token", code: INVALID_TOKEN}
	ResourceLockedError  = &WhiteLabelError{errorMessage: "Resource Locked", code: RESOURCE_LOCKED}
	InternalServerError  = &WhiteLabelError{errorMessage: "Internal Server Error", code: INTERNAL_SERVER_ERROR}
)
