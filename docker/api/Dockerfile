FROM golang:latest

COPY src/api /go/src/api/
WORKDIR /go/src/api/

RUN go get github.com/gin-gonic/gin \
    && go get github.com/codegangsta/gin

CMD ["gin", "-i", "run"]