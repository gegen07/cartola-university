FROM golang:alpine

RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/pilu/fresh

COPY ./app /go/src/github.com/gegen07/cartola-coltec
WORKDIR /go/src/github.com/gegen07/cartola-coltec

CMD dep ensure && fresh

## for production
# CMD dep ensure && go build && ./cartola-coltec

EXPOSE 8080