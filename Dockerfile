FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/YerkhanG/Assignment3GO
RUN cd /build && git clone https://github.com/YerkhanG/Assignment3GO.git
RUN cd /build/Assignment3GO/main && go build

EXPOSE 3000
ENTRYPOINT ["/build/Assignment3/main/main"]







