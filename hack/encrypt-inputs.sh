#!/usr/bin/env bash
set -euo pipefail

find inputs -name '*.txt' -not -name '*.enc.txt' -print0 \
  | while read -rd $'\0' decrypted
do
  encrypted="${decrypted%.txt}.enc.txt"
  if [ ! -f "$encrypted" ] || ! cmp -s "$decrypted" <(sops -d "$encrypted"); then
    echo "$decrypted -> $encrypted"
    sops -e "$decrypted" > "$encrypted"
    git add "$encrypted"
  fi
done
