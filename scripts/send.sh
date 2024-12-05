#!/usr/bin/env bash

CurrWorkDir=$(pwd)
ScriptRootDir=$(dirname "$0")
cd "$ScriptRootDir" || exit 1
cd ..

ZIPFILE="app.zip"

scp "$ZIPFILE" "udin@pharma:~"

cd "$CurrWorkDir" || exit 1
