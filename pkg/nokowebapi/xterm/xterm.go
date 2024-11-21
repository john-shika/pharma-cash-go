package xterm

import (
	"github.com/mattn/go-colorable"
	"os"
)

var Stdin = os.Stdin
var Stdout = colorable.NewColorableStdout()
var Stderr = colorable.NewColorableStderr()
