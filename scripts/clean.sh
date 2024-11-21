#!/usr/bin/env bash

CurrWorkDir=$(pwd)
ScriptDir=$(dirname "$0")
cd "$ScriptDir" || exit 1
cd ..

FILES=$(find . -type f | grep -Ei '\.([a-zA-Z0-9]+)\~$')

while IFS= read -r line; do
  FILE="$line"
  if [ -f "$FILE" ]; then
    echo "$FILE"
    rm "$FILE"

  fi

done <<< "$FILES"

cd "$CurrWorkDir" || exit 1
