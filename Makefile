init:
	which govendor || go get github.com/kardianos/govendor
	govendor sync

test: init
	go test -v ./...
