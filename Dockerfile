# Docker builder for Golang
FROM golang:latest as builder

RUN mkdir -p /vibes-api
WORKDIR /vibes-api
COPY . .

ENV GO111MODULE=on

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -o build/out .

# Docker run Golang app
FROM alpine
RUN apk update && \
   apk add ca-certificates && \
   update-ca-certificates && \
   rm -rf /var/cache/apk/*

COPY --from=builder /vibes-api/build/out /app
RUN chmod 700 /app
EXPOSE 3000
CMD ["./app"]