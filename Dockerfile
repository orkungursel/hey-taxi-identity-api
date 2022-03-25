FROM golang:1.17-buster AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags build -ldflags="-s -w" -trimpath -o /bin/app cmd/main.go

FROM scratch
COPY --from=build /bin/app /bin/app

EXPOSE 8080
EXPOSE 50051

ENTRYPOINT ["/bin/app"]