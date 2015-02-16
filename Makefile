OUT = marathon-cli

all: clean deps build

build:
	go build -o $(OUT)

clean:
	rm -f $(OUT) 

deps:
	go get github.com/Sirupsen/logrus
	go get github.com/codegangsta/cli 
	go get github.com/asteris-llc/gomarathon
