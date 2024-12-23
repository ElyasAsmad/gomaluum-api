FROM golang:1.23.2-alpine AS build
LABEL org.opencontainers.image.source=https://github.com/nrmnqdds/gomaluum-api

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags="-s -w" -o /app/gomaluum main.go

# use debug so can docker exec
FROM gcr.io/distroless/static-debian11:debug AS final
COPY --from=build /app/gomaluum /

ENV APP_ENV=production
ENV PORT=1323
ENV HOSTNAME=0.0.0.0

USER nonroot:nonroot

EXPOSE 1323

ENTRYPOINT ["/gomaluum", "-p"]
