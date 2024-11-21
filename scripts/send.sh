#!/usr/bin/env bash

CurrWorkDir=$(pwd)
ScriptDir=$(dirname "$0")
cd "$ScriptDir" || exit 1
cd ..

ZIPFILE="app.zip"

scp "$ZIPFILE" "udin@pharma:~"

cd "$CurrWorkDir" || exit 1
