#!/usr/bin/env bash

export $(cat init/.env | xargs) && go run cmd/gopx-web/*.go