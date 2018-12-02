FROM golang:alpine as golang


FROM smizy/mecab:0.996-alpine as mecab


FROM golang as runtime

ENV runtime_deps "openssl musl-dev libstdc++ git g++"
ENV CGO_LDFLAGS "-L/usr/local/lib -lmecab -lstdc++"
ENV CGO_CFLAGS "-I/usr/local/include"

RUN apk add --no-cache ${runtime_deps}

COPY --from=mecab / /


FROM runtime as app
ENV GO111MODULE=on \
    GOOS=linux
WORKDIR /go/github.com/nekottyo/newlogosorder

COPY go.mod go.sum ./
RUN go mod download

COPY  . .
RUN go build --ldflags '-extldflags "-static"' -o app

CMD ["./app"]

#FROM scratch
#COPY --from=app /go/github.com/nekottyo/newlogosorder /app
#
#CMD ["./app"]
