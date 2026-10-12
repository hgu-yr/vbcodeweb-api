# 构建阶段
FROM golang:1.26-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/vbcodeweb .

# 运行阶段
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -u 10001 app
WORKDIR /app
COPY --from=build /out/vbcodeweb /app/vbcodeweb
USER app
EXPOSE 8899
# 运行时通过挂载把 prod.yaml 放到 /app/prod.yaml，并设置 APP_ENV=prod
ENV APP_ENV=prod GIN_MODE=release
ENTRYPOINT ["/app/vbcodeweb"]
