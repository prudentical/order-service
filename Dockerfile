FROM golang:1.22.2-alpine3.19 as build
WORKDIR /app

ENV GOPROXY https://goproxy.io

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service ./cmd/web-service/


FROM gcr.io/distroless/static:nonroot
WORKDIR /app

COPY --from=build /app/database/migrations ./database/migrations
COPY --from=build /app/config.yml .
COPY --from=build /app/order-service .

EXPOSE 8001

CMD ["./order-service"]