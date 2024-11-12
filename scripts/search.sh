#!/usr/bin/env bash

CurrWorkDir=$(pwd)
ScriptDir=$(dirname "$0")
cd "$ScriptDir" || exit 1
cd ..

FILES=$(find . -type f | grep -Ei '\.go$')

while IFS= read -r line; do
  FILE="$line"
  echo "$FILE"
  < "$FILE" grep -Ei "$1"

done <<< "$FILES"

cd "$CurrWorkDir" || exit 1
