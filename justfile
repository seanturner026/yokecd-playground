_default:
    just --list

alias b := build
# Build the WASM binary
build:
    GOOS=wasip1 GOARCH=wasm go build -o .wasm/foo/production/app.wasm ./apps/foo/production/main.go

alias d := deploy
# Deploy the WASM binary
deploy: build
    yoke apply foo .wasm/foo/production/app.wasm
    yoke inspect foo

alias ku := kind-up
# Create the Kind Cluster
kind-up:
    kind create cluster --name local

alias kd := kind-down
# Cleanup the Kind Cluster
kind-down:
    kind delete clusters local
