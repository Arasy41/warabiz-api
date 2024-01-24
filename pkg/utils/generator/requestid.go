package generator

import (
	"fmt"

	"warabiz/api/pkg/utils/encryption"
)

func GenerateRequestID(sessionID string, ticks int) string {
	return encryption.GenerateSHA256(fmt.Sprintf("%s:%v", sessionID, ticks))
}
