FROM golang:alpine

ARG app_env
ENV APP_ENV $app_env

RUN apk add --no-cache git
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/pilu/fresh

WORKDIR /go/src/github.com/gegen07/cartola-coltec/app/cartola-coltec-service

COPY . .

CMD ["fresh"]

## for production
# CMD dep ensure && go build && ./cartola-coltec

EXPOSE 8080