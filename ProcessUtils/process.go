package ProcessUtils

import (
	"github.com/pandengyang/utils/FileUtils"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"
)

func ProcessExists(pid int) (exist bool) {
	if err := syscall.Kill(pid, 0); err == nil {
		return true
	}

	return false
}

func ProcessExistsByPidfile(pidfile string) (exist bool) {
	exist, err := FileUtils.FileExists(pidfile)
	if err != nil {
		return false
	}

	if !exist {
		return false
	}

	bytes, err := ioutil.ReadFile(pidfile)
	if err != nil {
		return false
	}

	pid, err := strconv.Atoi(strings.Trim(string(bytes), "\r\n "))
	if err != nil {
		return false
	}

	return ProcessExists(pid)
}
