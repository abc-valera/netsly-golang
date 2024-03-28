#!/bin/bash

# Turn on git hooks for the project
git config --local core.hooksPath .githooks

# Install go tools
echo "Downloading tools 📦 (It can take some time...)"
export GOBIN=$PWD/bin
go install github.com/pressly/goose/v3/cmd/goose
go install github.com/volatiletech/sqlboiler/v4
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
go install github.com/vektra/mockery/v2
go install github.com/99designs/gqlgen

# Pull docker images
echo "Pulling docker images 🐳 (It can take even more time.....)"
docker pull redocly/cli
docker pull postgres:15-alpine
docker pull redis/redis-stack:latest

# Create docker network
docker network create netsly-network

echo "Workspace initialized 🚀"
echo "You can start coding now! 👨‍💻 / 👩‍💻"