target = ./bin/goapi

all: build

build:
	go build -o $(target)

clean:
	rm -rfv pkg
	rm -rf $(target)
