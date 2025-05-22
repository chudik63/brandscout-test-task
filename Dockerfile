FROM golang:1.23 AS build

WORKDIR /build

COPY go.mod go.sum ./     

RUN go mod download      

COPY . .       

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

EXPOSE 8080

COPY --from=build /build/main .

COPY --from=build /build/.env .env

CMD ["./main"]