FROM golang:1.11-stretch AS build-env

ARG app_env
ENV ENV $app_env

ADD . ${GOPATH}/src/github.com/tanmaydatta/boggle

RUN go get ${GOPATH}/src/github.com/tanmaydatta/boggle/cmd

RUN cd ${GOPATH}/src/github.com/tanmaydatta/boggle/cmd && go build -o ../bin/boggle


#FROM alpine
#WORKDIR /app
#COPY --from=build-env /go/src/github.com/tanmaydatta/boggle/bin/boggle /app/
RUN ls /go/src/github.com/tanmaydatta/boggle/bin/
RUN ls /go/src/github.com/tanmaydatta/boggle/
#RUN pwd
EXPOSE 8080
WORKDIR ${GOPATH}/src/github.com/tanmaydatta/boggle
CMD ./bin/boggle
