#!/usr/bin/env bash

set -x

CurrWorkDir=$(pwd)
ScriptDir=$(dirname "$0")
cd "$ScriptDir" || exit 1
cd ..

APPEXE="exe"
APPDIR="app"
DSTEXE="/opt/pharma/exe"

ZIPFILE="app.zip"
SERVICE="pharma.service"

# remove app dir cache
if [ -d "$APPDIR" ]; then
  rm -rf "$APPDIR"

fi

# unpack app zip
unzip "$ZIPFILE" -d "$APPDIR"

# build new version
cd "$APPDIR"
go build -o "$APPEXE" .

# remove previous version
if [ -f "$DSTEXE" ]; then
  sudo rm -f "$DSTEXE"

fi

# install
sudo cp "$APPEXE" "$DSTEXE"
sudo systemctl restart "$SERVICE"

cd "$CurrWorkDir" || exit 1
