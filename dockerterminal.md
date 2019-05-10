Setup the Dockerfile

```
native user$ docker build -t cloud-native-go:1.0.0 .
Sending build context to Docker daemon  11.82MB
Step 1/8 : FROM golang
 ---> 7ced090ee82e
Step 2/8 : MAINTAINER Ngan Nguyen
 ---> Using cache
 ---> 2ac167ffbdb2
Step 3/8 : ENV SOURCES /go/src/native
 ---> Running in 87fd69331f15
Removing intermediate container 87fd69331f15
 ---> 3ac0852ee81b
Step 4/8 : COPY . $SOURCES
 ---> 8838c1f9a380
Step 5/8 : RUN cd ${SOURCES} && CGO_ENABLED=0 go install
 ---> Running in ae19490cc181
Removing intermediate container ae19490cc181
 ---> f5f8504599ed
Step 6/8 : ENV PORT 8080
 ---> Running in 457ced0d9559
Removing intermediate container 457ced0d9559
 ---> 9233c186cc38
Step 7/8 : EXPOSE 8080
 ---> Running in 434e16882da4
Removing intermediate container 434e16882da4
 ---> a7d90cfe43b0
Step 8/8 : ENTRYPOINT Cloud-Native-Go
 ---> Running in 1bccd1fc8711
Removing intermediate container 1bccd1fc8711
 ---> 41e097db9c64
Successfully built 41e097db9c64
Successfully tagged cloud-native-go:1.0.0

```

verify docker image, 802MB is too big

```
native user$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
cloud-native-go     1.0.0               41e097db9c64        3 minutes ago       802MB
golang              latest              7ced090ee82e        40 hours ago        774MB

```
update docker file to use alpine

```
native user$ docker pull golang:1.12.5-alpine3.9
1.12.5-alpine3.9: Pulling from library/golang
bdf0201b3a05: Pull complete 
38f114998adb: Pull complete 
21134b1a9e68: Pull complete 
acf7a24e9fe4: Pull complete 
ac62bb6f9bd2: Pull complete 
Digest: sha256:fdb2912263a50a4157808e1952f52a13d9403f30cbfbf3d08c8e6f1f81554620
Status: Downloaded newer image for golang:1.12.5-alpine3.9
```

run docker build after upgrade to alpine
```
native user$ docker build -t cloud-native-go:1.0.0 .
Sending build context to Docker daemon  11.83MB
Step 1/8 : FROM golang:1.12.5-alpine3.9
 ---> 91e7250c8003
Step 2/8 : MAINTAINER Ngan Nguyen
 ---> Running in 3418c454f705
Removing intermediate container 3418c454f705
 ---> eb27814d9bd4
Step 3/8 : ENV SOURCES /go/src/native
 ---> Running in b1d9d61ca372
Removing intermediate container b1d9d61ca372
 ---> 5bc3657ecd1b
Step 4/8 : COPY . $SOURCES
 ---> e01369449a1d
Step 5/8 : RUN cd ${SOURCES} && CGO_ENABLED=0 go install
 ---> Running in 0112a6b597bf
Removing intermediate container 0112a6b597bf
 ---> dfbf20fe9197
Step 6/8 : ENV PORT 8080
 ---> Running in 031f1020327d
Removing intermediate container 031f1020327d
 ---> 3837125b123f
Step 7/8 : EXPOSE 8080
 ---> Running in 25935b4dcd52
Removing intermediate container 25935b4dcd52
 ---> 20aaaa8f2aff
Step 8/8 : ENTRYPOINT Cloud-Native-Go
 ---> Running in 9c5cfb3dad06
Removing intermediate container 9c5cfb3dad06
 ---> 4241c2096709
Successfully built 4241c2096709
Successfully tagged cloud-native-go:1.0.0

```

Forgot to change the tag from 1.0.0 to 1.0.1 which is why none is in the list of images referencing old image

```
native user$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
cloud-native-go     1.0.0               4241c2096709        2 minutes ago       378MB
<none>              <none>              41e097db9c64        10 minutes ago      802MB
golang              latest              7ced090ee82e        40 hours ago        774MB
golang              1.12.5-alpine3.9    91e7250c8003        2 days ago          350MB

```

Tag Docker image to docker username

```
native user$ docker tag cloud-native-go:1.0.0 ngann/cloud-native-go:1.0.0
:native user$ docker images
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
cloud-native-go         1.0.0               4241c2096709        11 minutes ago      378MB
ngann/cloud-native-go   1.0.0               4241c2096709        11 minutes ago      378MB
golang                  latest              7ced090ee82e        40 hours ago        774MB
golang                  1.12.5-alpine3.9    91e7250c8003        2 days ago          350MB
:native user$ 

```

Login and push image to dockerhub
The image should show in docker hub

```
docker login
docker push ngann/cloud-native-go

native user$ docker push ngann/cloud-native-go
The push refers to repository [docker.io/ngann/cloud-native-go]
ba6607939eeb: Pushed 
0d125c03a927: Pushed 
69c38bdeacee: Mounted from library/golang 
84a9917047c2: Mounted from library/golang 
daf1b73f0d7f: Mounted from library/golang 
d0e4fb621c6d: Mounted from library/golang 
a464c54f93a9: Mounted from library/golang 
1.0.0: digest: sha256:53807369a4d1c6c05561c336d20f968c0fa89e64c6a8c86cf5d330b9b086b5c9 size: 1787

```

Run image locally built previously

```
docker run -p 8888:5000 ngann/cloud-native-go

docker run -it -p 8080:8080 cloud-native-go:1.0.0
```
Error , reference incorrect name of entry point which was
Cloud-Native_Go but should be native , name of my directory.
```
/bin/sh: Cloud-Native-Go: not found
```

Created a new docker image and ran the image locally
User should be able to head to 
To see hello message :http://localhost:8080/api/hello
To see list of books: http://localhost:8080/api/books
To see list of echo message: http://localhost:8080/api/books



```
docker build -t cloud-native-go:1.0.2 .
docker run -it -p 8080:8080 cloud-native-go:1.0.2

docker run -it -e "PORT=9090" -p 9090:9090 cloud-native-go:1.0.2
```