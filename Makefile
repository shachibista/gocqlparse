gen:
	@antlr4 Cql3.g4 -no-listener -visitor -o internal/parser

test: gen
	go test -v ./...
