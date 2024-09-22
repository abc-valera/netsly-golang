#!/usr/bin/bash

# This a post install script for the dev container.

echo "Running post-create script 🛠️"

# Install all the dependencies
echo "Downloading tools and dependencies 📦 (It can take some time...)"
go mod download
npm install --save-dev prettier prettier-plugin-go-template prettier-plugin-tailwindcss

# Install project tools
export GOBIN=$PWD/bin
go install github.com/golangci/golangci-lint/cmd/golangci-lint
go install github.com/mgechev/revive
go install mvdan.cc/gofumpt
go install github.com/air-verse/air
go install github.com/vektra/mockery/v2

# Pull docker images
echo "Pulling docker images 🐳 (It can take even more time.....)"
docker pull redocly/cli

# Create docker network
docker network create netsly-network

# Turn on git hooks for the project
git config --local core.hooksPath .githooks

# Create .env files
cp env/.example.dev.env env/dev.env
cp env/.example.test.env env/test.env

echo "Workspace initialized 🚀"
echo "You can start coding now! 👨‍💻 / 👩‍💻"
