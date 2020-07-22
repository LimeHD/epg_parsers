SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.2.dev

all: clean build

build: bin/digea_parser bin/ept_parser bin/star_parser

parse_all:
	./bin/star_parser
	./bin/ept_parser
	./bin/digea_parser
	./download_stv && ./bin/mejor_parser
	./bin/mejor_cards_parser
	./download_tv_pack && ./bin/standard_parser

clean:
	rm -f bin/*

bin/star_parser:
	go get && go build -o ./bin ./star_parser

bin/ept_parser:
	go get && go build -o ./bin ./ept_parser

bin/digea_parser:
	go get && go build -o ./bin ./digea_parser

bin/mejor_paser:
	./download_stv && ./bin/mejor_parser

bin/mejor_cards_parser:
	./bin/mejor_cards_parser

help:
	@echo "make all  		: Build all parsers"
	@echo "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo "Please report any bug or error to the author."
