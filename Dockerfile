## CI
FROM golang:1.18 AS test
COPY . /src
WORKDIR /src
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install
CMD [ "go", "run", "." ]


## release
FROM golang:1.18 AS builder
WORKDIR /src
# GOOS是作業系統，這邊如果要改成在linux之下則可以修改為linux
# GOARCH則是平台架構，可以用的有386、amd64、ARM
ENV GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN  go mod download
COPY . .
RUN go vet ./... && go build -o /bin/server

FROM alpine:latest AS release
# Copy from builder
COPY --from=builder /bin/server /bin/server
COPY --from=builder /src/.env .env
EXPOSE 8080
ENTRYPOINT ["./bin/server"]