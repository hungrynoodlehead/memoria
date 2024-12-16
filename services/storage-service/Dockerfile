FROM golang AS base

WORKDIR /build
COPY go.mod .
RUN go mod verify && go mod download

COPY . .

FROM base AS build-prod

RUN CGO_ENABLED=0 go build -o app .

FROM base AS build-dev

RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o app-debug .

FROM scratch AS prod

EXPOSE 8080
ENV CONFIG_ENVIRONMENT=PRODUCTION

COPY --from=build-prod /build/app /

ENTRYPOINT ["/app"]

FROM golang:alpine AS dev

EXPOSE 8080 40000
ENV CONFIG_ENVIRONMENT=DEVELOPMENT

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY --from=build-dev /build/app-debug /app-debug
ENTRYPOINT ["dlv", "--continue", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app-debug"]