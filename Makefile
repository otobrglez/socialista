.PHONY: dependencies
dependencies:
	go get golang.org/x/tools/cmd/vet
	# go get github.com/golang/lint/golint

vet: dependencies
	go tool vet -v ./
