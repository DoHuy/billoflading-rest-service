#!/bin/bash
# shellcheck disable=SC2164
(
  swag init --parseDependency=true -o docs -d external/delivery/rest -g ../../../cmd/main.go
)