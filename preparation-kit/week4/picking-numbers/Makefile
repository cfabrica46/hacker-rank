BUILDPATH=$(CURDIR)
API_NAME=picking-numbers

test:
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage:
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -func coverfile_out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' > coverage.txt

.PHONY: test coverage all clean
