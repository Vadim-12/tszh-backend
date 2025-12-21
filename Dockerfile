FROM golang:1.24 AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./cmd

FROM gcr.io/distroless/base-debian12
ENV PORT=8080
EXPOSE 8080

COPY --from=build /out/app /app
COPY --from=build /go/bin/migrate /migrate
COPY schema /schema

USER nonroot:nonroot
ENTRYPOINT ["/app"]