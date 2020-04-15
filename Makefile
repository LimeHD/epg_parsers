SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1.dev

all: build

build: digea_parser 

ept_parser:
	go get && go build ./ept_parser.go

digea_parser:
	go get && go build ./digea_parser.go

help:
	@echo "make all  		: Build all parsers"
	@echo "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo "Please report any bug or error to the author."
