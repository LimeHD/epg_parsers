SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1.dev

all: build

build: bin/digea_parser 

bin/ept_parser:
	go get && go build -o ./bin ./ept_parser

bin/digea_parser:
	go get && go build -o ./bin ./digea_parser

help:
	@echo "make all  		: Build all parsers"
	@echo "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo "Please report any bug or error to the author."
