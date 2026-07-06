FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /cais ./cmd/cais

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=build /cais /usr/local/bin/cais
ENTRYPOINT ["cais"]