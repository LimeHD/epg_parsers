SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.2.dev

all: clean build

build: go_get_all bin/digea_parser bin/ept_parser bin/star_parser

go_get_all:
	go get -v all

parse_all:
	./bin/star_parser
	./bin/ept_parser
	./bin/digea_parser
	./download_stv && ./bin/stv_parser
	./bin/dvbs_parser
	./download_tv_pack ./output/TV_Pack.xml && ./bin/tv_pack_parser ./output/TV_Pack.xml
	./bin/alfaomega

clean:
	rm -f bin/*

bin/star_parser:
	go build -o ./bin ./star_parser

bin/ept_parser:
	go build -o ./bin ./ept_parser

bin/digea_parser:
	go build -o ./bin ./digea_parser

help:
	@echo "make all  		: Build all parsers"
	@echo "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo "Please report any bug or error to the author."
