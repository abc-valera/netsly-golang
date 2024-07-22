#!/usr/bin/bash

# This a post install script for the dev container.

echo "Running post-create script 🛠️"

# Install all the tools that weren't installed earlier
go install mvdan.cc/gofumpt@latest

# Turn on git hooks for the project
git config --local core.hooksPath .githooks

# Install project tools and dependencies
echo "Downloading tools and dependencies 📦 (It can take some time...)"
go mod download

export GOBIN=$PWD/bin
go install github.com/pressly/goose/v3/cmd/goose
go install github.com/volatiletech/sqlboiler/v4
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3
go install github.com/vektra/mockery/v2

npm install --save-dev prettier prettier-plugin-go-template prettier-plugin-tailwindcss

# Pull docker images
echo "Pulling docker images 🐳 (It can take even more time.....)"
docker pull redocly/cli
docker pull postgres:15-alpine
docker pull redis/redis-stack:latest

# Create docker network
docker network create netsly-network

echo "Workspace initialized 🚀"
echo "You can start coding now! 👨‍💻 / 👩‍💻"
