FROM golang
MAINTAINER Ngan Nguyen

ENV SOURCES /go/src/native

#location of files from local directory into a docker image
COPY . $SOURCES

#build the go mircroservies
#change into the source directory and build the sources
RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

#name of the executable to build
ENTRYPOINT Cloud-Native-Go


#docker build -t cloud-native-go:1.0.0 .