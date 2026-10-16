FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Version and build time shown on startup, build time defaults to the time of the image build
ARG BUILD_VERSION=dev
ARG BUILD_TIME
RUN go build -ldflags "-X 'main.BuildVersion=${BUILD_VERSION}' -X 'main.BuildTime=${BUILD_TIME:-$(date +'%Y-%m-%dT%H:%M:%S%z')}'" -o weather-exporter .



FROM alpine:latest
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

COPY --from=builder /app/weather-exporter .

EXPOSE 8080 8443
USER app
CMD ["./weather-exporter"]
