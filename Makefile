PROJECTNAME=$(shell basename "$(PWD)")
PROJECTNAME=grpc_demo
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`

all: help

.PHONY: help proto clean

## proto	compiling proto file
proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product.proto

## clean	clean proto file
clean:
	@rm -R proto/*.pb.go

## help	print this help message and exit.
help: Makefile
	@echo "Choose a command run in "$(PROJECTNAME)":"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Valid target values are:"
	@echo ""
	@sed -n 's/^## //p' $<
