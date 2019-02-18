## Introduction

Previously we made our deployment on Docker. While docker is good in in development environment there are other platform available in order to deploy your app in production

Some of the most known is Docker Swarm and Kubernetes. Let's see the difference between them

## Docker swarm 🐳

Docker Swarm is a container orchestrator that is integrated to Docker. It aim to simplify the management of multiple containers just like what Kubernetes does.

Kubernetes is also a container orchestrator. However it differ from Docker swarm in some ways.

Indeed Docker swarm has the advantages of being tightly integrated with Docker. This mean that Docker swarm can effectively use the Docker API which lead to an easier usage of Docker swarm than Kubernetes

As docker swarm is tightly coupled to the machine it run. This mean that swarm is platform dependant. As you know sometimes it can get difficult to configure docker so you might encounter those issues with the Swarm mode too.

One important point in which Docker swarm differ from Kubernetes is how it's considering it's deployment system.

As per the new stack article: 

> The applications can be deployed as micro-services or services in a swarm cluster in Docker Swarm. YAML(YAML Ain’t Markup Language) files can be utilized to identify multi-container. Moreover, Docker compose can install the application.

This differ from what Kubernetes does:

Kubernetes can be deployed by using a service or by a deployment types which is creating nodes, and pods... which can be scaled up easily by using a ```yaml``` configuration file

However Docker swarm is not used widely which mean that the community isn't strong. Moreover Docker swarm doesn't provide a way to do host a stateful application while Kubernetes is providing a mode for creating Stateful application

Futhermore Docker swarm doesn't come with a good monitoring service which Kubernetes is providing out of the box

TL;DR

- Docker swarm easy to use
- Fast configuration
- Doesn't support stateful application
- Doesn't provide a service for monitoring
- Too tightly coupled to the OS
- Small communities -> harder to get support

While Docker swarm is an interesting tool we have a database which we need to deploy with Kubernetes for the sake of the exercise. If we wanted to run it outside then Docker swarm could have been a good candidate though not a lot of company support it out of the box.

## Learning Kubernetes 🧠

As said in the main README. Kubernetes is a complex piece of software. Therefore we need to get a good grasp of it's internal before doing anything. Let's check these parts together 👏