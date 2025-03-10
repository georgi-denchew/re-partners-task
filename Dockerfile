FROM golang:1.24.1-bookworm AS base

FROM base AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Running tests as part of image build as there's no CI to execute tests
RUN go test -count=1 ./...

RUN CGO_ENABLED=0 go build -o orders-service

FROM scratch AS production

WORKDIR /prod

COPY --from=builder /build/orders-service ./

EXPOSE 3000

CMD ["/prod/orders-service"]
