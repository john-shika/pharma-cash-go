#!/usr/bin/env bash

FILES=$(find . -type f | grep -Ei '\.[a-zA-Z0-9]+\~$')

while IFS= read -r line; do
  FILE="$line"
  echo "$FILE"
  rm "$FILE"

done <<< "$FILES"
