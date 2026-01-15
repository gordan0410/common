package response

import (
	"context"
	"net/http"
	"time"

	"github.com/gordan0410/common/enum"
	"github.com/gordan0410/common/enum/wh_error"
)

type Response struct {
	Code         int         `json:"code"`
	ExternalCode string      `json:"excode,omitempty"`
	Msg          string      `json:"msg"`
	Data         interface{} `json:"data,omitempty"`
	TraceID      string      `json:"trace_id"`
	Timestamp    int64       `json:"timestamp"`
}

func NewResponse(ctx context.Context) *Response {
	traceIdAny := ctx.Value(enum.TraceId.ToString())
	traceId, _ := traceIdAny.(string)

	return &Response{
		Timestamp: time.Now().UnixMilli(),
		TraceID:   traceId,
	}
}

// GenerateTimestamp 寫入回傳時間
func (r *Response) GenerateTimestamp() {
	r.Timestamp = time.Now().UnixMilli()
}

// GetResponseByError 依照丟進來的 error 產生回傳
func GetResponseByError(ctx context.Context, err error) (int, *Response) {
	r := NewResponse(ctx)

	// 如果有自定義 error 則需要去抓對應的 msg
	if bsError, ok := wh_error.IsWhiteLabelError(err); ok {
		r.Msg = bsError.Error()
		r.Code = bsError.ErrorCode()

		return http.StatusBadRequest, r
	} else if err != nil {
		// 非自定義 error 統一回預設
		r.Msg = wh_error.BadRequestError.GetMessage()
		r.Code = wh_error.BadRequestError.GetCode()

		return http.StatusForbidden, r
	}

	r.Msg = API_CODE_MSG_MAP[CODE_SUCCESS]
	r.Code = CODE_SUCCESS

	return http.StatusOK, r
}

// GetSuccessResponse 產生成功回傳
func GetSuccessResponse(ctx context.Context) (int, *Response) {
	r := NewResponse(ctx)
	r.Msg = API_CODE_MSG_MAP[CODE_SUCCESS]
	r.Code = CODE_SUCCESS

	return http.StatusOK, r
}

// GetSuccessResponseWithData 產生成功回傳且塞data
func GetSuccessResponseWithData(ctx context.Context, data interface{}) (int, *Response) {
	r := NewResponse(ctx)
	r.Msg = API_CODE_MSG_MAP[CODE_SUCCESS]
	r.Code = CODE_SUCCESS
	r.Data = data

	return http.StatusOK, r
}
