FROM golang:latest

ADD . /go/src/Assignment3
WORKDIR /go/src/Assignment3
RUN go get Assignment3
RUN go install
ENTRYPOINT ["/go/bin/Assignment3"]





