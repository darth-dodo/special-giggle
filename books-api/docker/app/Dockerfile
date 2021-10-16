FROM golang:1.17-alpine
WORKDIR /books-api/

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

CMD ["./bin/app"]
EXPOSE 8080