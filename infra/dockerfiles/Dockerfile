FROM golang:1.23.2 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./build/ ./app/cmd

FROM scratch
COPY --from=builder /app/build /build
ENTRYPOINT ["/build"]
