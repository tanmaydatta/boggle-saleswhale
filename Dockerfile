FROM golang:1.11-stretch AS build-env

ARG app_env
ENV ENV $app_env

ADD . ${GOPATH}/src/github.com/tanmaydatta/boggle

RUN go get ${GOPATH}/src/github.com/tanmaydatta/boggle/cmd

RUN cd ${GOPATH}/src/github.com/tanmaydatta/boggle/cmd && go build -o ../bin/boggle


EXPOSE 8080
WORKDIR ${GOPATH}/src/github.com/tanmaydatta/boggle
CMD ./bin/boggle
