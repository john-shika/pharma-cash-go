#!/usr/bin/env bash

set -ex

CurrWorkDir=$(pwd)
ScriptRootDir=$(dirname "$0")
cd "$ScriptRootDir" || exit 1
cd ..

AppExe="exe"
AppDir="app"
DstExe="/opt/pharma/exe"
DstSqlite="/opt/pharma/migrations/dev.sqlite3"

ZipFile="app.zip"
Service="pharma.service"

if [ -d "$AppDir" ]; then
  rm -rf "$AppDir"

fi

unzip "$ZipFile" -d "$AppDir"

cd "$AppDir" || exit 1
go build -o "$AppExe" .

if [ -f "$DstExe" ]; then
  sudo rm -f "$DstExe"

fi

sudo rm "$DstSqlite"

sudo cp "$AppExe" "$DstExe"
sudo systemctl restart "$Service"

cd "$CurrWorkDir" || exit 1
