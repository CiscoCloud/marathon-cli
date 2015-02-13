OUT = marathon-cli

all: clean deps build

build:
	go build -o $(OUT) $(OUT).go

clean:
	rm -f $(OUT) 

deps:
	go get github.com/Sirupsen/logrus
	go get github.com/codegangsta/cli 
	go get github.com/jbdalido/go-marathon
