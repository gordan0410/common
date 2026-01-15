package helper

import "testing"

func TestHMACSignSHA256(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		secretKey string
		want      string
	}{
		{
			name:      "basic HMAC-SHA256",
			data:      "hello world",
			secretKey: "secret",
			want:      "734cc62f32841568f45715aeb9f4d7891324e6d948e4c6c60c0621cdac48623a",
		},
		{
			name:      "empty data",
			data:      "",
			secretKey: "secret",
			want:      "f9e66e179b6747ae54108f82f8ade8b3c25d76fd30afde6c395822c530196169",
		},
		{
			name:      "empty secret",
			data:      "hello",
			secretKey: "",
			want:      "4352b26e33fe0d769a8922a6ba29004109f01688e26acc9e6cb347e5a5afc4da",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HMACSignSHA256(tt.data, tt.secretKey)
			if got != tt.want {
				t.Errorf("HMACSignSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHMACSignSHA1(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		secretKey string
		want      string
	}{
		{
			name:      "basic HMAC-SHA1",
			data:      "hello world",
			secretKey: "secret",
			want:      "03376ee7ad7bbfceee98660439a4d8b125122a5a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HMACSignSHA1(tt.data, tt.secretKey)
			if got != tt.want {
				t.Errorf("HMACSignSHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyHMACSignSHA256(t *testing.T) {
	data := "test data"
	secretKey := "test secret"
	validSignature := HMACSignSHA256(data, secretKey)

	tests := []struct {
		name      string
		data      string
		secretKey string
		signature string
		want      bool
	}{
		{
			name:      "valid signature",
			data:      data,
			secretKey: secretKey,
			signature: validSignature,
			want:      true,
		},
		{
			name:      "invalid signature",
			data:      data,
			secretKey: secretKey,
			signature: "invalid",
			want:      false,
		},
		{
			name:      "wrong secret key",
			data:      data,
			secretKey: "wrong secret",
			signature: validSignature,
			want:      false,
		},
		{
			name:      "wrong data",
			data:      "wrong data",
			secretKey: secretKey,
			signature: validSignature,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VerifyHMACSignSHA256(tt.data, tt.secretKey, tt.signature)
			if got != tt.want {
				t.Errorf("VerifyHMACSignSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}
