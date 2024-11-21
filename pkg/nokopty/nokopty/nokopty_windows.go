//go:build windows

package nokopty

import "golang.org/x/sys/windows"

type OsVersionEx struct {
	Major int
	Minor int
	Patch int
}

func GetOsVersion() *OsVersionEx {
	major, minor, patch := windows.RtlGetNtVersionNumbers()
	return &OsVersionEx{
		Major: int(major),
		Minor: int(minor),
		Patch: int(patch),
	}
}
