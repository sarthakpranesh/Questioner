FROM golang
ADD . /go/src/github.com/sarthakpranesh/Questioner
WORKDIR /go/src/github.com/sarthakpranesh/Questioner
RUN go mod tidy
RUN go install
ENTRYPOINT /go/bin/Questioner
EXPOSE 8080
