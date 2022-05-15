FROM golang:1.18-alpine as builder
LABEL stage=builder

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app .

RUN go build -o /build/app ./cmd/server

FROM alpine:3.15

COPY --from=builder ./build/ /bin/
CMD ["/bin/app"]