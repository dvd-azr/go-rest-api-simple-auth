BINARY_NAME=rest-api-simple-auth.exe

#  build: builds all binaries
    # @go mod vendor
build:
   	@go build -o tmp/${BINARY_NAME} .
	@echo Celeritas built!
# 
run:
	@echo Staring dvd...&
    @go run main.go
	@echo dvd started!
