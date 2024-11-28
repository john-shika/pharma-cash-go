package motd

import (
	"bytes"
	"fmt"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

func Motd() string {
	name := globals.Get("nokowebapi.name").(string)
	description := globals.Get("nokowebapi.description").(string)
	version := globals.Get("nokowebapi.version").(string)
	author := globals.Get("nokowebapi.author").(string)
	url := globals.Get("nokowebapi.url").(string)
	license := globals.Get("nokowebapi.license").(string)
	licenseUrl := globals.Get("nokowebapi.license_url").(string)
	banners := []string{
		"\n" +
			" /$$   /$$           /$$                 /$$      /$$           /$$        /$$$$$$            /$$\n" +
			"| $$$ | $$          | $$                | $$  /$ | $$          | $$       /$$__  $$          |__/\n" +
			"| $$$$| $$  /$$$$$$ | $$   /$$  /$$$$$$ | $$ /$$$| $$  /$$$$$$ | $$$$$$$ | $$  \\ $$  /$$$$$$  /$$\n" +
			"| $$ $$ $$ /$$__  $$| $$  /$$/ /$$__  $$| $$/$$ $$ $$ /$$__  $$| $$__  $$| $$$$$$$$ /$$__  $$| $$\n" +
			"| $$  $$$$| $$  \\ $$| $$$$$$/ | $$  \\ $$| $$$$_  $$$$| $$$$$$$$| $$  \\ $$| $$__  $$| $$  \\ $$| $$\n" +
			"| $$\\  $$$| $$  | $$| $$_  $$ | $$  | $$| $$$/ \\  $$$| $$_____/| $$  | $$| $$  | $$| $$  | $$| $$\n" +
			"|__/  \\__/ \\______/ |__/  \\__/ \\______/ |__/     \\__/ \\_______/|_______/ |__/  |__/| $$____/ |__/\n" +
			"                                                                                   | $$          \n" +
			"                                                                                   | $$          \n" +
			"                                                                                   |__/          \n" +
			"",
		"\n" +
			" _____     _       _ _ _     _   _____     _ \n" +
			"|   | |___| |_ ___| | | |___| |_|  _  |___|_|\n" +
			"| | | | . | '_| . | | | | -_| . |     | . | |\n" +
			"|_|___|___|_,_|___|_____|___|___|__|__|  _|_|\n" +
			"                                      |_|    \n" +
			"",
		"\n" +
			"   _  __     __      _      __    __   ___        _ \n" +
			"  / |/ /__  / /_____| | /| / /__ / /  / _ | ___  (_)\n" +
			" /    / _ \\/  '_/ _ \\ |/ |/ / -_) _ \\/ __ |/ _ \\/ / \n" +
			"/_/|_/\\___/_/\\_\\\\___/__/|__/\\__/_.__/_/ |_/ .__/_/  \n" +
			"                                         /_/        \n" +
			"",
	}

	size := len(banners)
	m := nokocore.RandomRangeInt(0, size)
	buffer := bytes.Buffer{}

	buffer.WriteString(fmt.Sprintf("%s\n", banners[m]))
	buffer.WriteString(fmt.Sprintf("name: %s\n", name))
	buffer.WriteString(fmt.Sprintf("description: %s\n", description))
	buffer.WriteString(fmt.Sprintf("version: %s\n", version))
	buffer.WriteString(fmt.Sprintf("author: %s\n", author))
	buffer.WriteString(fmt.Sprintf("url: %s\n", url))
	buffer.WriteString(fmt.Sprintf("license_url: %s\n", licenseUrl))
	buffer.WriteString(fmt.Sprintf("license: %s\n", license))

	return buffer.String()
}
