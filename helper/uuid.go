package helper

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func GetUUID() string {
	u := uuid.New()
	return hex.EncodeToString(u[:])
}
