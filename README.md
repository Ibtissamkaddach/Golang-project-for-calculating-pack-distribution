# Golang-project-for-calculating-pack-distribution
Golang project for calculating pack distribution on Visual Studio Code (VS Code) with a fully functional HTTP API
# Pack Calculator API

## Description

This application calculates the optimal number of packs needed to fulfill an order.

## Endpoints

- `/calculate-packs?order=<number>`: Calculate the packs for the given order.

## How to Run

1. Clone the repository.
2. Run `go mod tidy` to install dependencies.
3. Run `go run main.go` to start the server.
4. Access the API at `http://localhost:8080/calculate-packs?order=<number>`.

## Example

```sh
curl "http://localhost:8080/calculate-packs?order=2750"
