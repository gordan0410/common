package wh_error

import (
	"errors"
	"fmt"
)

// WhiteLabelErrorI 須符合 go 本身 error
type WhiteLabelErrorI interface {
	Error() string
	ErrorCode() int
}

type WhiteLabelError struct {
	WhiteLabelErrorI
	errorMessage string
	code         int
	detail       string
}

func (base *WhiteLabelError) Error() string {
	return fmt.Sprintf("error code: %d, message: %s, detail: %s", base.code, base.errorMessage, base.detail)
}

func (base *WhiteLabelError) ErrorCode() int {
	return base.code
}

func (base *WhiteLabelError) WithDetail(detail string) *WhiteLabelError {
	clone := *base // 複製一份值（shallow copy）
	clone.detail = detail
	return &clone
}

func (base *WhiteLabelError) GetMessage() string {
	return base.errorMessage
}

func (base *WhiteLabelError) GetCode() int {
	return base.code
}

func IsWhiteLabelError(err error) (WhiteLabelErrorI, bool) {
	var bsError WhiteLabelErrorI
	if errors.As(err, &bsError) {
		return bsError, true
	}

	return nil, false
}
