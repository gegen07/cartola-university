FROM golang:alpine

RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/pilu/fresh

WORKDIR /go/src/github.com/gegen07/cartola-coltec/app/cartola-coltec-service

COPY . /go/src/github.com/gegen07/cartola-coltec/app/cartola-coltec-service

CMD dep ensure && fresh

## for production
# CMD dep ensure && go build && ./cartola-coltec

EXPOSE 8080