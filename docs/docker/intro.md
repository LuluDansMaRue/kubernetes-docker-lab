---
layout: single
classes: wide
title: Docker 🐳
sidebar:
  nav: "docs"
---

Docker is a software that allows [Operating system level virtualization](https://en.wikipedia.org/wiki/Container_(virtualization)) thanks to user space instances called **containers**. These containers are isolated processes containing applications that can communicate with each other and with other applications outside of their container thanks to Docker networking.

Thus this allows developers to create multiple test environments on the fly and to share them with others, without too much headaches with a few configurations files.

## Architecture[¶](#architecture)

Docker architecture is centralized on a daemon (background process), which is used by all of Docker's components. This daemon named dockerd managed the containers. Please take a look at the dockerd documentation for further information [dockerd documentation](https://docs.docker.com/engine/reference/commandline/dockerd/)

In order to control these components. Docker use a REST API endpoint that is use to control them. This endpoint is also used by the docker CLI. These request are then send to the daemon which will treat these requests.

An interesting article regarding this topic here: [Communicate with the Docker rest API](https://blog.trifork.com/2013/12/24/docker-from-a-distance-the-remote-api/)

<p align="center"> 
  <img src="../img/docker.png" alt="drawing" width="400"/>
</p>
<p align="center"><b>Docker architecture. Docker documentation</b></p>

## Components[¶](#components)

In order to use Docker, we need to know what components it has.

- Image: A template that is use to run your container (e.g. the ubuntu image).
- Container: A runnable instance of an image with your custom configuraton i.e: a custom ubuntu image customize for running nginx on it
- Data volumes: Virtual volumes that are used to persist your container's data (it's based on your FileSystem) that can be shared to other containers based how you configure it
- Network: When instantiating multiple containers, docker allows you to connect them all together by creating a network system. Docker will manipulate ip tables in order to provide network isolation.

## Using Docker[¶](#using-docker)

There are many ways of using Docker: From creating a simple container to testing a script on the fly with ```docker run ...```, or by creating a complex development environment with ```docker-compose```. Futhermore, you can even create a custom environment on an OS by using ```Dockefile```. Docker allows us to enhance our flexibility to work with other environment.

In order to understand how to use Docker. Take a look at a deployment which show ```docker-compose and a Dockerfile```. [Creating development environment](./example.md)