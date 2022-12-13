FROM golang:1.18
COPY . /src
WORKDIR /src
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install
EXPOSE 8080
CMD [ "go", "run", "." ]