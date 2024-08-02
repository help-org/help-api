FROM golang:1.22 AS build
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o directory-api

FROM alpine:latest AS final
WORKDIR /app
COPY --from=build /build/directory-api .
COPY --from=build /build/.env .

EXPOSE 8080

ENTRYPOINT ["./directory-api"]
