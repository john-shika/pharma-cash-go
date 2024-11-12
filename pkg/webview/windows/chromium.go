package windows

import "os"

/*
 * C:\\Program Files
 * C:\\Program Files (x86)
 *
 * Microsoft\Edge\Application\msedge.exe
 * Google\Chrome\Application\chrome.exe
 *
 * msedge.exe --app=http://localhost --user-data-dir=C:\\Users\%USER%\.webview\microsoft-edge\profiles --profile-directory=localhost
 * chrome.exe --app=http://localhost --user-data-dir=C:\\Users\%USER%\.webview\google-chrome\profiles --profile-directory=localhost
 *
 * %USERPROFILE%\AppData\Local\NokoWebView\Profiles\Google Chrome\Default
 * %USERPROFILE%\AppData\Local\NokoWebView\Profiles\Microsoft Edge\Default
 *
 * %USERPROFILE%\AppData\Local %LOCALAPPDATA%
 * %USERPROFILE%\AppData\Roaming %APPDATA%
 * C:\\ProgramData %PROGRAMDATA%
 *
 * // firefox can't customizable profile as well
 * C:\\"Program Files"\"Mozilla Firefox"\firefox.exe -P app -new-instance -url http://localhost
 *
 */

var programFiles = []string{
	"C:\\Program Files",
	"C:\\Program Files (x86)",
	// %USERPROFILE%\AppData\Local\Programs
}

var chromeApps = []string{
	"Microsoft\\Edge\\Application\\msedge.exe",
	"Google\\Chrome\\Application\\chrome.exe",
}

var userProfile = os.Getenv("USERPROFILE")
