# syntax=docker/dockerfile:1

FROM golang:1.21.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=1 GOOS=linux go build -o /go-go-htmx

EXPOSE 3000

# Run
CMD ["/go-go-htmx"]