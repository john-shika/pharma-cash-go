#!/usr/bin/env bash

CurrWorkDir=$(pwd)
ScriptRootDir=$(dirname "$0")
cd "$ScriptRootDir" || exit 1
cd ..

ZIPFILE="app.zip"

if [ -f "$ZIPFILE" ]; then
    rm "$ZIPFILE"
fi

FILES=(
    "app"
    "pkg"
    "main.go"
    "go.mod"
    "go.work"
    "nokowebapi.yaml.example"
)

zip -r0q9yo "$ZIPFILE" "${FILES[@]}"

cd "$CurrWorkDir" || exit 1
