FROM golang:1.19 as common-build-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN ls .
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./src/app.go
RUN chmod +x ./app
RUN ls .
RUN ls -l ./app
EXPOSE 3000

# Run
CMD ["./app"]