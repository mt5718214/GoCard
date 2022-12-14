FROM golang:1.18
COPY . /src
WORKDIR /src
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install
EXPOSE 8080
CMD [ "go", "run", "." ]