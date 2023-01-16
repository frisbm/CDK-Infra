.PHONY: help
help: # Print this help message
	@grep '^[a-z]' Makefile

.PHONY: install
install: # Install required dependencies for the project
	go mod tidy; \
	  go mod download; \
	  go mod verify
	npm i -g aws-cdk
	npm i -g cdk-dia
	brew install

.PHONY: diagrams
diagrams: # Create Diagrams for Readme
	bash ./scripts/generate-diagrams.sh
