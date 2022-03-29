FROM golang:1.18 AS build
WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM prom/busybox
COPY --from=build /app/app /app
CMD ["/app"]
EXPOSE 2000
EXPOSE 12000
USER 1:1
