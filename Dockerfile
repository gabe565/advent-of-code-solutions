#syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.27.3-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache \
  CGO_ENABLED=0 GOARCH="$TARGETARCH"\
  go build -ldflags='-w -s' -trimpath


FROM alpine:3.24
WORKDIR /
COPY --from=build /app/advent-of-code-solutions /
ENTRYPOINT ["/advent-of-code-solutions"]
