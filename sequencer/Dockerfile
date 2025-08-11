FROM golang:1.23.0-alpine

# Create a non-root user and group with UID and GID 1000
# RUN addgroup -g 1000 appgroup && adduser -D -u 1000 -G appgroup appuser

WORKDIR /app

ARG APP_MODE
ARG AWS_REGION
ARG AWS_DB_SECRET_NAME
ARG CONTRACT_ADDRESS
ARG ETHEREUM_RPC
ARG PRIVATE_KEY

ENV APP_MODE=${APP_MODE}
ENV AWS_REGION=${AWS_REGION}
ENV AWS_DB_SECRET_NAME=${AWS_DB_SECRET_NAME}
ENV CONTRACT_ADDRESS=${CONTRACT_ADDRESS}
ENV ETHEREUM_RPC=${ETHEREUM_RPC}
ENV PRIVATE_KEY=${PRIVATE_KEY}

# for internal health check
RUN apk add --no-cache curl 

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Create Go build cache directory and set ownership
# RUN mkdir -p /app/.cache/go-build && chown -R appuser:appgroup /app/.cache

# Set Go build cache environment variable
# ENV GOCACHE=/app/.cache/go-build

EXPOSE 8080

# Switch to non-root user
# USER appuser:appgroup

CMD ["go", "run", "main.go"]