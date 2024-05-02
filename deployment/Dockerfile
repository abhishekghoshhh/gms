FROM golang:1.22.1-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# download the dependencies
RUN go mod download

COPY . .

# build the go app
RUN go build -o app /app/cmd/gms/

# New Stage
FROM alpine:latest

WORKDIR /app/

COPY --from=builder app .

EXPOSE 8080

# Command to run the executable
CMD ["./app"]
