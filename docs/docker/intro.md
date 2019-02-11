# Docker

Docker is a container based virtualization software that allow it's user to create a development environment in a easy way without installing a virtualization software i.e VirtualBox.

This allow developer to create multiple test environment on the fly and to share with others without too much headaches.

# Architecture

Docker architecture is centralized on the daemon. The daemon is used by all of the component of Docker. 

In order to control all of these components Docker is using a rest api endpoint that is used by the docker cli. This CLI will send request to the api endpoint which will then be analyze by the docker daemon.

An interesting article regarding this topic can be founded here: [Communicate with the Docker rest API](https://blog.trifork.com/2013/12/24/docker-from-a-distance-the-remote-api/)

<p align="center"> 
  <img src="../img/docker.png" alt="drawing" width="400"/>
</p>
<p align="center"><b>Docker architecture. Docker documentation</b></p>

# Components

In order to use Docker we need to know the several components that has Docker.

- Image: An image is a template that is used to running your container. i.e: the ubuntu image
- Container: a container is a runnable instance of an image with your custom configuraton i.e: a custom ubuntu image customize for running nginx on it
- Data volumes: Data volumes is are used to persisting the data of your containers
- Network: When creating multiple container, docker allow you to connect them together by creating a network system. Docker will manipulate iptable in order to provide network isolation to your docker containers

# Using Docker

There are many ways of using Docker. From creating a simple container to test a script on the fly with ```docker run ...```. Orr to create a complex development environment with the usage of ```docker-compose```. Or to create a custom environment on an OS by using a ```Dockefile```. Docker allow us to simplify how we work and should provide us an out of the box development environment for everyone.

To see how to create a development environment with docker-compose & Dockerfile please take a look at this section [Creating development environment](./example.md)