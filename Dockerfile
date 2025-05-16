# Stage 1: Build
FROM golang:1.21-bullseye AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# Stage 2: Run (minimal)
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/main /bin/main
CMD ["/bin/main", "/etc/env/.env"]