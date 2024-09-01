FROM golang:1.23-alpine
RUN apk update && apk add --no-cache make
RUN mkdir /go/src/app/
WORKDIR /go/src/app/
COPY ./ /go/src/app/

# RUN make tidy

# ENTRYPOINT [ "go", "run", "./api/" ]
