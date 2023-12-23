
FROM golang:1.21.0 as builder
ARG CGO_ENABLED=0

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# RUN go build
# Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o FN-Go-Basic-Services .

FROM alpine:latest as runtime
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/FN-Go-Basic-Services .
# COPY . FN-Go-Basic-Services

CMD ["./FN-Go-Basic-Services"]

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 9080
