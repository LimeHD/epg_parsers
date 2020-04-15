SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1.dev

PROJECTNAME=$(shell basename "$(PWD)")
PROJECTDIR=$(shell pwd)

all: help

build-digea:
	cd $(PROJECTDIR)/bin/ && go build -o ./builds digea.go

digea:
	./bin/builds/digea --format csv --output $(PROJECTDIR)/bin/output

help:
	@echo -e "build-*  	: Build parser, where * - is parsername. Example: build-digea"
	@echo -e "digea  	: Run parser Digea & store to csv file\n"
	@echo -e "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo -e "Please report any bug or error to the author."