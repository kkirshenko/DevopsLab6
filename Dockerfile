FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
# The repository currently does not contain go.sum; resolve and record the
# complete dependency checksums before compiling the application.
RUN go mod tidy && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /app/server ./cmd/server

FROM alpine:3.20
RUN adduser -D -H app
USER app
COPY --from=build /app/server /server
EXPOSE 8080 9090
ENTRYPOINT ["/server"]
