# Build stage
FROM golang:1.27-alpine AS build
WORKDIR /src

# Download dependencies first, so this layer is cached until go.mod or go.sum change.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /streamgate .

# Runtime stage: static binary only, runs as a non-root user.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /streamgate /streamgate
EXPOSE 8080
ENTRYPOINT ["/streamgate"]
