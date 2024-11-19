//go:build linux

package nokopty

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

type OsVersionEx struct {
	Major int
	Minor int
	Patch int
}

func GetOsVersion() *OsVersionEx {
	var uname syscall.Utsname
	err := syscall.Uname(&uname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting uname: %v\n", err)
		return nil
	}

	release := toString(uname.Release[:])
	versionParts := strings.Split(release, ".")

	if len(versionParts) < 3 {
		fmt.Fprintf(os.Stderr, "unexpected version format: %s\n", release)
		return nil
	}

	major, _ := parseInt(versionParts[0])
	minor, _ := parseInt(versionParts[1])
	patch, _ := parseInt(strings.Split(versionParts[2], "-")[0])

	return &OsVersionEx{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func toString(ca [65]int8) string {
	n := 0
	for ; n < len(ca); n++ {
		if ca[n] == 0 {
			break
		}
	}
	return string(ca[:n])
}

func parseInt(str string) (int, error) {
	var val int
	_, err := fmt.Sscanf(str, "%d", &val)
	if err != nil {
		return 0, err
	}
	return val, nil
}
