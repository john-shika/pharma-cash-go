package motd

import (
	"fmt"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

func Motd() {
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
	m := nokocore.RandomRangeInt(0, len(banners))
	fmt.Println(banners[m])
	fmt.Println("name:", name)
	fmt.Println("description:", description)
	fmt.Println("version:", version)
	fmt.Println("author:", author)
	fmt.Println("url:", url)
	fmt.Println("license_url:", licenseUrl)
	fmt.Println("license:", license)
}
