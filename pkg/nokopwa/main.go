package main

import (
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"nokowebapi/globals"
	"nokowebapi/motd"
	"nokowebapi/nokocore"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	programFiles := []string{
		"C:\\Program Files",
		"C:\\Program Files (x86)",
		// %USERPROFILE%\AppData\Local\Programs
	}

	chromeApps := []string{
		"Microsoft\\Edge\\Application\\msedge.exe",
		"Google\\Chrome\\Application\\chrome.exe",
		"Mozilla Firefox\\firefox.exe",
		"Opera\\opera.exe",
	}

	userProfileDir := os.Getenv("USERPROFILE")
	localAppDataDir := filepath.Join(userProfileDir, "AppData", "Local")

	programFiles = append(programFiles, filepath.Join(localAppDataDir, "Programs"))

	var err error
	var path string
	var found bool

	nokocore.KeepVoid(err, path, found)

	for i, chromeApp := range chromeApps {
		nokocore.KeepVoid(i)

		for j, programFile := range programFiles {
			nokocore.KeepVoid(j)

			webViewPath := filepath.Join(programFile, chromeApp)
			if path, err = exec.LookPath(webViewPath); err != nil {
				continue
			}

			found = true
			break
		}

		if found {
			break
		}
	}

	fmt.Println(userProfileDir)
	fmt.Println(path)

	motd.Motd()

	temp := make(map[string]any)
	nokocore.NoErr(mapstructure.Decode(globals.GetJwtConfig(), &temp))
	fmt.Println(nokocore.ShikaYamlEncode(temp))
}
