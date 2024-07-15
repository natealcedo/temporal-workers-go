# Build stage
FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN make deps
RUN make build

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 8080
CMD ["./myapp"]