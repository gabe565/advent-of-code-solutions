#!/usr/bin/env bash
set -euo pipefail

find inputs -name '*.enc.txt' -print0 | while read -rd $'\0' encrypted; do
  decrypted="${encrypted%.enc.txt}.txt"
  if [ ! -f "$decrypted" ] || ! cmp -s "$decrypted" <(sops -d "$encrypted"); then
    echo "$encrypted -> $decrypted"
    sops -d "$encrypted" > "$decrypted"
  fi
done
