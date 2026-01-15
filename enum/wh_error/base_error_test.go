package wh_error

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWhiteLabelError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  WhiteLabelErrorI
		want1 bool
	}{
		{
			name: "normal error",
			args: args{
				err: fmt.Errorf("is not wh error"),
			},
			want:  nil,
			want1: false,
		},
		{
			name: "wh error ",
			args: args{
				err: BadRequestError,
			},
			want:  BadRequestError,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsWhiteLabelError(tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsWhiteLabelError got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsWhiteLabelError got1 = %v, want %v", got1, tt.want1)
			}
			if got != nil {
				fmt.Println(got.Error())
			}
		})
	}
}

// TestWhiteLabelError_Error 測試 WithDetail 不會影響到原本的 singleton
func TestWhiteLabelError_WithDetail(t *testing.T) {

	// 呼叫 WithDetail 並不該影響 original
	origin := InvalidPayloadError
	newErr := origin.WithDetail("missing required field")

	// 驗證原本未被修改
	assert.Equal(t, "", origin.detail, "original.detail should not be modified")

	// 驗證新錯誤有正確 detail
	assert.Equal(t, "missing required field", newErr.detail)
	assert.Equal(t, origin.code, newErr.code)
	assert.Equal(t, origin.errorMessage, newErr.errorMessage)

	// 確認不是同一個 pointer（真的複製了）
	assert.NotSame(t, origin, newErr)
}
