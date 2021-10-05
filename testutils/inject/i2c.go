package inject

import (
	"go.viam.com/core/board"
)

// I2C is an injected I2C.
type I2C struct {
	board.I2C
	OpenHandleFunc func() (board.I2CHandle, error)
}

// OpenHandle calls the injected OpenHandle or the real version.
func (s *I2C) OpenHandle() (board.I2CHandle, error) {
	if s.OpenHandleFunc == nil {
		return s.I2C.OpenHandle()
	}
	return s.OpenHandleFunc()
}