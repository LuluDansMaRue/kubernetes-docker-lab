# Docker

Docker is a container-based virtualization software that allows users to easily create a development environment without installing any virtualization software (e.g. VirtualBox).

This allows developers to create multiple test environments on the fly and to share them with others, without too much headaches.

# Architecture

Docker architecture is centralized on the daemon, which is used by all of Docker's components.

In order to control all of those components, Docker uses a rest api endpoint that is also used by the docker CLI. This CLI will send requests to the api endpoint which will then be processed by the docker daemon.

An interesting article regarding this topic here: [Communicate with the Docker rest API](https://blog.trifork.com/2013/12/24/docker-from-a-distance-the-remote-api/)

<p align="center"> 
  <img src="../img/docker.png" alt="drawing" width="400"/>
</p>
<p align="center"><b>Docker architecture. Docker documentation</b></p>

# Components

In order to use Docker, we need to know what components it has.

- Image: It is a template that is used to run your container (e.g. the ubuntu image).
- Container: a container is a runnable instance of an image with your custom configuraton i.e: a custom ubuntu image customize for running nginx on it
- Data volumes: they are used to persist the data of your containers
- Network: When instantiating multiple containers, docker allows you to connect them all together by creating a network system. Docker will manipulate ip tables in order to provide network isolation.

# Using Docker

There are many ways of using Docker: from creating a simple container to testing a script on the fly with ```docker run ...```, or creating a complex development environment with ```docker-compose```. Even creating a custom environment on an OS by using ```Dockefile```. Docker allows us to simplify how we work and should provides us an out of the box development environment.

To see how to create a development environment with docker-compose & Dockerfile, please take a look at this section [Creating development environment](./example.md)