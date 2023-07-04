//go:build linux || darwin || freebsd || netbsd || openbsd || illumos
// +build linux darwin freebsd netbsd openbsd illumos

package watch

import "os"

func IsDeletePending(_ *os.File) (bool, error) {
	return false, nil
}
