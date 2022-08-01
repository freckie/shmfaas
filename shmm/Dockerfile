# Author: freckie@frec.kr
#
# docker build --platform linux/amd64 .

FROM golang:alpine AS builder

# for gcc
RUN apk add musl-dev libc-dev build-base

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=1

ENV SQLITE3=shmm.db3 \
    PORT=8080

WORKDIR /build
COPY . ./

RUN go mod download
RUN go build -a -ldflags '-linkmode external -extldflags "-static"' -o main .

WORKDIR /dist
RUN cp /build/main .

### 

FROM scratch

EXPOSE 8080

WORKDIR /app
COPY --from=builder /dist/main .

ENTRYPOINT [ "./main" ]
