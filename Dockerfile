FROM golang:1.16 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy the go source
COPY src/main.go src/main.go
COPY src/pkg/ src/pkg/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o rolloutproxy src/main.go

FROM gcr.io/distroless/static

WORKDIR /
COPY --from=builder /workspace/rolloutproxy .

ENTRYPOINT ["/rolloutproxy"]
