#!/usr/bin/env bash

# Export the necessary environment variables
export $(cat scripts/.env | xargs)

# Install the server executable
go install ./cmd/gopx-web

# Run the server
sudo gopx-web