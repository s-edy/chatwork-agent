all: gom gom-install build install

gom:
	go get -u github.com/mattn/gom

gom-install:
	gom install

build: chatwork-agent.go chatwork/*.go
	gom build $(GOFLAGS) -o chatwork-agent chatwork-agent.go

install:
	cp chatwork-agent $(GOPATH)/bin/chatwork-agent

fmt:
	go fmt ./...

clean:
	rm chatwork-agent
