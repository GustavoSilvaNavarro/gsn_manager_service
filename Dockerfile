# Stage 1: Build the application using the custom base image
FROM gsn_manager_service_base as builder

WORKDIR /usr/src/app
COPY . .

# Build the binary with optimizations for production
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go-app ./src


FROM alpine:latest

RUN apk --no-cache add curl netcat-openbsd

RUN adduser default --disabled-password
USER default

COPY --from=builder /go-app /go-app

EXPOSE 8080

CMD ["/go-app"]
