package ProcessUtils

import (
	"syscall"
)

func ProcessExists(pid int) (alive bool) {
	if err := syscall.Kill(pid, 0); err == nil {
		alive = true
	}

	return alive
}
