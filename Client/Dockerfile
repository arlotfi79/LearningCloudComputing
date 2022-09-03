FROM golang:1.18-alpine AS builder

WORKDIR /app

# Copy (go.mod, go.sum) and install dependencies in /app
COPY Client/go.mod Client/go.sum ./
RUN go mod download

# copy sourse code to /app
COPY . .

WORKDIR /app/Client

# Build app
RUN go build -o main .

# ---------- multi-stage --------------
FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/Client/main /app/Client/main

# Run the binary program produced by `go install`
CMD ["/app/Client/main"]