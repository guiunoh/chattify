##
## build
##
FROM golang:latest AS builder
ARG package=./cmd/thumbnailer

WORKDIR /build
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -tags=jsoniter -a -ldflags '-w -s' -o main ${package}

##
## deploy
##
FROM alpine:latest
ARG USER=deploy
ARG config=config.yaml
ENV HOME /home/$USER

RUN apk add --update sudo
RUN adduser -D $USER \
    && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER

USER $USER
WORKDIR $HOME

COPY --from=builder /build/main .
COPY --from=builder /build/config/config*.yaml ./config/
COPY --from=builder /build/entrypoint.sh .

RUN sudo chown -R $USER:$USER $HOME
RUN chmod +x ./entrypoint.sh

ENV APP_CONFIG_FILE ${config}
ENV GIN_MODE release
ENTRYPOINT ["./entrypoint.sh"]
