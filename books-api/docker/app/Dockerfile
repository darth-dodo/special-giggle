# Build environment
# -----------------

FROM golang:1.17-alpine as build-env
WORKDIR /books-api

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
# ----------------------

FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /books-api/bin/app /books-api/

EXPOSE 8080
CMD ["/books-api/app"]