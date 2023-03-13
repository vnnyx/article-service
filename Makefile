tidy:
	go mod tidy

generate:
	go run github.com/99designs/gqlgen generate

wire:
	wire ./graph/di