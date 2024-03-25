FROM golang:1.21 as common-build-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy

# Utils
RUN go install github.com/go-task/task/v3/cmd/task@latest

# Build
RUN task build
RUN chmod +x ./app

EXPOSE 3000

# Run
CMD ["./app"]