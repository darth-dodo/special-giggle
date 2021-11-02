FROM golang:1.17.6-apline as builder
COPY go.mod go.sum /special-giggle/products-api/
RUN go mod download
COPY . /special-giggle/products-api/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/products-api github.com/darth-dodo/special-giggle/products-api


FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder ./special-giggle/products-api/ /usr/bin/products-api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/products-api"]
