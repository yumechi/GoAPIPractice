FROM golang:latest

COPY api /go/src/api/
WORKDIR /go/src/api/

RUN go get github.com/gin-gonic/gin \
    && go get github.com/codegangsta/gin

CMD ["gin", "-p", "37301", "-i", "run"]
