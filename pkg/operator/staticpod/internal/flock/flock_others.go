// +build !linux

package flock

import (
	"fmt"
	"time"
)

// TryLock is not supported/implemented on this platform
func (f *FLock) TryLock(timeout time.Duration) error {
	return fmt.Errorf("flock is not supported on this platform")
}

// Unlock is not supported/implemented on this platform
func (f *FLock) Unlock() error {
	return fmt.Errorf("flock is not supported on this platform")
}
