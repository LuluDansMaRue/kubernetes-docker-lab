# Kubernetes Docker experiments

## Introduction

Welcome fan of Bubble tea.
This project's aim to people who want to learn how to deploy a web application with several environments:

- Docker for your development environment
- Kubernetes for your production environment on GCloud


## Table of contents

* [Docker](./docker/intro.md)
* * [Use case example of Docker](./docker/example.md)
* [Kubernetes](./k8s/intro.md)
* * [Architecture of Kubernetes](./k8s/architecture.md)
* * [Type of deployment](./k8s/deployment.md)
* * [Minikube for prototyping](./k8s/minikube.md)
* * [Example of a deployment](./k8s/deployment/example.md)


## Objectives

At the end of this project you should be able to get a good grasp on Docker, Kubernetes and GCloud.

Below are the objectives that we should have reach

```
- Be able to deploy an app on Docker
- Understand the architecture of Kubernetes
- Understand the difference between Kubernetes & Docker Swarm
- Know how to use minikube for local prototyping
- Know how to deploy a web app from A to B with kubernetes within GCloud
- Understand when & how to scale a cluster on multiple pods & on several timezone.
- Be able to add an other service to an existing GCloud product (e.g: GCloud SQL) 
- Be able to understand & implement a deamon, service & side-car (e.g: Cloud SQL Proxy)
- Understand & know how to deploy a load balancer
- Be able to know implement a monitoring process by using stackdriver & grafana.
- Create Cron task on GCloud 
```

## The stack

As describe on the main Readme our web application is divide in 3 parts.

- The back-end API which is made in Go
- The front-end which is small app made with VueJS
- A simple MySQL database

Everything is build up with a Docker :). Take a look at the docker-compose for more information or check out the related Docker page for more information

### Back-end

The back-end contains 4 APIs that are used by the front-end. A postman collection is available within the ```postman``` folder. I invite you to test the APIs :).

### Front-end

WIP

### Database

The database is a simple MySQL database with one table named Bobba.
This table store the information related to the bubble tea which are:

```shell
- Id <int>
- Name <string>
- Price <double>
- Shop <string>
- Flavor <string>
- Calory <double>
```

