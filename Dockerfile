FROM golang:1.17-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-w -s" ./cmd/server.go

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder /build/server .

ENV PORT=3000
ENV ADDRESS=0.0.0.0

# Export necessary port.
EXPOSE ${PORT}

# Command to run when starting the container.
CMD ["./server"]