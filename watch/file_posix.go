//go:build illumos || linux || darwin || freebsd || netbsd || openbsd
// +build illumos linux darwin freebsd netbsd openbsd

package watch

import "os"

func IsDeletePending(_ *os.File) (bool, error) {
	return false, nil
}
