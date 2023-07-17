test:
	go test ./... -v --cover -race

test-profile:
	mkdir -p tmp
	go test ./... --coverprofile=tmp/coverage.txt
	go tool cover -html=tmp/coverage.txt

.PHONY: test
.PHONY: test-profile
