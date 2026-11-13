FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /app/server ./cmd/server

FROM alpine:3.20
RUN adduser -D -H app
USER app
COPY --from=build /app/server /server
EXPOSE 8080 9090
ENTRYPOINT ["/server"]
