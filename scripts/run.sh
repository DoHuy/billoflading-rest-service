#!/bin/bash
# shellcheck disable=SC2164

# export config env
config_file='./resources/config.json'

for keyval in $(grep -E '": [^\{]' $config_file | sed -e 's/: /=/' -e "s/\(\,\)$//"); do
  eval export $keyval
done

# run build
go run ./cmd

