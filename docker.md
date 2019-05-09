Docker
Application and OS libraries are put in containers.
Container are ran on host OS then provides access to virtual hardware or real hardware

Docker Images:
* Base Image
* Derived Image

Docker Containers:
If you take an image and run it, then you get the container.
A container is the running state of an image


Basic Docker workflow
1. docker file / text file to build docker image
2. Use run command to run a container, take the state of a container to create a new image, or tag an image
3. push/pull to and from docker registry
4.Save and load command to back images


#### Basic Docker Command

Build a docker image with the give tag from currenty directory
```
$ docker build -t <tag>
```
Prints all local image
```
$docker images
```
Run a docker image: creates and runs a container
* in background
* with defined ENV variable
* with port forwarding for Host to Container
* image name and entry point

```
$docker run
-d
-e <enviroment varialble>
-p <host-port>:<container-port>
<image><entrypoint process>
```

Run a Docker image and open a shell within the container
* with forwarding of local termina;
* image name and shell(or ,,/bin/bash")
```
$docker run
-ti
<image> /bin/sh
```
Prints all containers(without --all prints only running containers)

```
$docker ps --all
```
Terminate container
```
$docker kill <container>
```
remove docker container
```
$docker rm <container>
```
remove local image
```
$docker rmi f <image>
``` 
Build Native Docker Image for Go MicroServices


Running Containerized Go microservices locally

Improving Docker IMage and Docker compose



Docker Hub > search golang > locate `Supported tags and respective Dockerfile links`


After installing docker:
verify the verisons
docker --version
pull image
docker pull golang

Write docker file:
create docker textfile .. Dockerfile


Build docker image from file:

Tag and Push docker image to docker hub:

