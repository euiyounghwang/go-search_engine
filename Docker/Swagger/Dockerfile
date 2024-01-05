
FROM golang:1.21.0 as builder
ARG CGO_ENABLED=0

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# RUN go build
# Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o swagger
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o FN-Go-Basic-Services .
RUN go build -o FN-Go-Basic-Services swagger.go

FROM alpine:latest as runtime
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# COPY --from=builder /app/FN-Go-Basic-Services .
COPY --from=builder /app/FN-Go-Basic-Services .

EXPOSE 9081

CMD ["./FN-Go-Basic-Services"]
# CMD ["./swagger"]

# within Docker
# apk 
# apk add curl