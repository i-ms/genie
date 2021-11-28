## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool genie and copies it to myapp
## on windows , the executable is named genie.exe
build_cli:
	@go build -o ../genieSample/genie ./cmd/cli