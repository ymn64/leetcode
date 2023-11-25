#!/bin/bash

declare -A columns

while IFS=' ' read -r -a row; do
  for ((i = 0; i < ${#row[@]}; i++)); do
    columns[$i]+="${row[i]} "
  done
done <file.txt

for ((i = 0; i < ${#columns[@]}; i++)); do
  echo "${columns[$i]:0:-1}"
done
