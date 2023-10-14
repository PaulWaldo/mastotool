run-race:
	GORACE="halt_on_error=1" go run -race .

run-race-log:
	GORACE="log_path=race/report halt_on_error=1" go run -race .

test-race:
	go test -race ./...

test:
	go test ./...

tidy:
	go mod tidy
	go fmt ./...
	go vet ./...

package:
	fyne package -icon assets/Icon.png -name Mastotool

mobilesim:
	go run -tags mobile main.go

showcov:
	go test ./... -coverprofile=coverage.out -covermode=count
	go tool cover -html=coverage.out
	rm coverage.out

testdocs:
	gotestdox ./...