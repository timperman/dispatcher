FROM golang:1.5.1

EXPOSE 8081
COPY . /go/src/github.com/timperman/dispatcher

RUN go get -d -v ./...

RUN go test ./... && go install ./...

ENTRYPOINT [ "/go/bin/dispatcher" ]
