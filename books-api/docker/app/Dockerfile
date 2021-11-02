# Build environment
# -----------------

FROM golang:1.17-alpine as build-env
WORKDIR /books-api

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------

FROM alpine
RUN apk update && apk add --no-cache bash
COPY --from=build-env /books-api/bin/app /books-api/
COPY --from=build-env /books-api/bin/migrate /books-api/
COPY --from=build-env /books-api/migrations /books-api/migrations

EXPOSE 8080
CMD ["/books-api/app"]
