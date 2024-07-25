FROM golang as builder

RUN mkdir app
COPY ./ app
WORKDIR app
RUN CGO_ENABLED=0 go build -o dist/api cmd/main.go

FROM golang as runner

RUN mkdir app
COPY --from=builder ./go/app/dist/api app/
COPY ./internal/config/config.yaml app/

RUN chmod +x app
WORKDIR app

ENTRYPOINT [ "./api", "--config-dir", "." ]