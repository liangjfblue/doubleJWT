package uuid

import (
    guuid "github.com/satori/go.uuid"
    "strings"
)

// NewUUID 生成uuid
func NewUUID() string {
    u2 := guuid.NewV4()
    return strings.Replace(u2.String(), "-", "", -1)
}

