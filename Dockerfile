FROM golang:1.14 AS builder
WORKDIR /builder/src/
COPY . .
RUN GOOS=linux go build driver.go

EXPOSE 8080
CMD ["/builder/src/driver"]