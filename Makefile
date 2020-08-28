API_NAME=toggl-card

help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t

## run all tests, builds and executes the api
all: test run	

## builds api binary file
build:	
	@cd cmd/$(API_NAME) && go build -o ../../bin/$(API_NAME)

## run all implemented tests
test:	
	@cd internal/card && go test -v
	@echo "-------------------------"
	@cd internal/deck && go test -v
	@echo "-------------------------"
	@cd pkg/server && go test -v -parallel=3

## cleans and deletes current binary
clean:
	@cd bin/ && go clean && rm -f $(API_NAME)

## builds binary file and runs it
run:	
	@cd cmd/$(API_NAME) && go build -o ../../bin/$(API_NAME)
	@cd bin && ./$(API_NAME)

## get all dependencies and install them 
get:	
	@go get -u github.com/google/uuid 
	@go get -u github.com/gorilla/mux 
	@go get -u github.com/stretchr/testify 

## generate documentation from code commets
doc:	
	@cd internal/card && go doc -all
	@echo "-------------------------"
	@cd internal/deck && go doc -all
	@echo "-------------------------"
	@cd pkg/server && go doc -all