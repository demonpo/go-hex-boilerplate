FROM golang:1.19 as common-build-stage
WORKDIR /app
COPY . .
RUN go mod download

# Utils
RUN go install github.com/go-task/task/v3/cmd/task@latest

# Build
RUN task build
RUN chmod +x ./app

EXPOSE 3000

# Run
CMD ["./app"]