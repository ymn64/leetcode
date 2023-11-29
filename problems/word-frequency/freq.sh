#!/bin/bash

declare -A freq

while read -r -a words; do
  for word in "${words[@]}"; do
    ((freq[$word]++))
  done
done <words.txt

for word in "${!freq[@]}"; do
  echo "$word ${freq[$word]}"
done | sort -k2 -n -r
