#!/usr/bin/env bash

go install ./cmd/gopx-web && export $(cat init/.env | xargs) && sudo gopx-web