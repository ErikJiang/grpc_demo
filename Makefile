PROJECTNAME=$(shell basename "$(PWD)")
PROJECTNAME=grpc_demo
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`

all: help

.PHONY: help proto clean

## proto	generate proto file
proto:
	@protoc -I proto --go_out proto \
		--go_opt paths=source_relative \
		--go-grpc_out proto \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out proto \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		proto/models.proto \
		proto/user.proto \
		proto/team.proto

## clean	clean proto file
clean:
	@rm -R proto/*.pb.*

## help	print this help message and exit.
help: Makefile
	@echo "Choose a command run in "$(PROJECTNAME)":"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Valid target values are:"
	@echo ""
	@sed -n 's/^## //p' $<
