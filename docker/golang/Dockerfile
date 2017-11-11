FROM golang:1.9

# app specific env
ENV BUILDPATH /go/src/middleware
ENV TOOLS /var/exec
ENV GOBINARIES /go/bin

RUN apt-get update && apt-get upgrade -y
RUN apt-get install telnet vim -y

# Install reflex
WORKDIR $GOBINARIES

RUN go get gopkg.in/check.v1
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/yaml.v2

RUN go get github.com/patrickmn/go-cache
RUN go get github.com/cespare/reflex
RUN go get github.com/codegangsta/negroni
RUN go get github.com/pborman/uuid
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get github.com/agtorre/gocolorize
RUN go get github.com/fatih/structs
RUN go get github.com/casbin/casbin
RUN go get github.com/casbin/mongodb-adapter
RUN go get github.com/stretchr/testify/assert
RUN go get github.com/mattn/goveralls

# Send to GOROOT
RUN cd /go/src/ && mv github.com /usr/local/go/src/
RUN cd /go/src/ && mv gopkg.in   /usr/local/go/src/

ENV PORT 6060
# dockeer/app port
EXPOSE $PORT

# Make directories and add files as needed
RUN mkdir -p $TOOLS
ADD build.sh $TOOLS
ADD reflex.conf $TOOLS
RUN chmod +x $TOOLS/build.sh

# Execute reflex.
WORKDIR $BUILDPATH
CMD ["reflex","-c","/var/exec/reflex.conf"]