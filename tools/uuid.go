package tools

import (
	"github.com/satori/go.uuid"
)

// UUidV4 return v4 uuid
func UUidV4() string {
	u4 := uuid.NewV4()
	return u4.String()
}
