FROM golang:1.22.1-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-w -s" -o apiserver ./cmd/api/main.go

FROM scratch

COPY --from=builder /build/apiserver /
COPY --from=builder /build/docs /docs
COPY --from=builder /build/.env /

EXPOSE 8080

ENTRYPOINT ["/apiserver"]
