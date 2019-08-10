FROM golang:alpine as builder
LABEL stage=intermediate
RUN apk update && apk add --no-cache git
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o gombot-go .
FROM scratch
LABEL maintainer="Genie Chae ygenie.chae@gmail.com"
COPY --from=builder /build/gombot-go /app/
EXPOSE 9801
VOLUME [ "/app/data" ]
WORKDIR /app
ENTRYPOINT ["/app/gombot-go"]