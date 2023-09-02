FROM golang:1.21-bullseye AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /bin/app

FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=build /bin/app /bin

CMD ["/bin/app"]