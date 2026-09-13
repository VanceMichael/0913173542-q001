FROM golang:1.23-alpine AS build
WORKDIR /src
ENV GOPROXY=https://proxy.golang.com.cn,direct
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /race-review .
FROM alpine:3.21
COPY --from=build /race-review /race-review
VOLUME ["/data"]
EXPOSE 8080
ENTRYPOINT ["/race-review"]
