# ===================================================================================
# === Stage 1: Build the Go service code into 'server' exe ==========================
# ===================================================================================

FROM golang:1.19-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go
# RUN go install github.com/cosmtrek/air@latest
# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# ================================================================================================
# === Stage 2: Get server binary into a lightweight container ====================================
# ================================================================================================

FROM alpine:latest AS production

COPY --from=builder /app .

# CMD ["air"]
CMD ["./app"]