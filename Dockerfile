FROM golang:1.14 AS builder
WORKDIR /builder/src/
COPY . .
RUN GOOS=linux go build driver.go

FROM scratch
WORKDIR /app
COPY --from=builder /builder/src/driver /app/driver
COPY drivers.json .
EXPOSE 8080
CMD ["/app/driver"]