---
classes: wide
title: Kubernetes Docker experiments 🎋
sidebar:
  nav: "docs"
---

## Introduction[¶](#introduction)

Welcome fan of Bubble tea.

This project is for people who are curious and want to learn how to deploy an application from A to B to several environements such as: 

- Docker for your development environment
- Minikube for emuling your Kubernetes production environment
- Kubernetes for your production environment on GCloud

Throughout this guide we'll learn several thing which are describe in the table of contents below

## Table of contents[¶](#table-of-contents)

* [Introduction to Docker](docker/intro.md)
* * [Example of deployment with Docker](docker/example.md)
* [A quick look at Docker Swarm](k8s/swarm.md)
* [Architecture](k8s/architecture.md)
* [Deployment](k8s/deployment.md)
* [Service](k8s/services.md)
* [Persistent volume](k8s/storage.md)
* [Cron task](k8s/cron.md)
* [Minikube](k8s/minikube.md)
* [Example with Minikube](k8s/deployment/example.md)
* * [Front deployment](k8s/deployment/front.md)
* * [API deployment](k8s/deployment/api.md)
* * [Service front deployment](k8s/deployment/service_front.md)
* * [Service API deployment](k8s/deployment/service_api.md)
* * [Database](k8s/deployment/database.md)
* [Deploying on Google Cloud Platform](gcp/intro.md)
* * [Build the API docker image](gcp/api.md)
* * [Configuring the SQL Database](gcp/sql.md)
* * [Deploying the API](gcp/deployment_api.md)
* * [Configure an Ingress resources](gcp/ingress.md)
* * [Build the Front docker image](gcp/front.md)
* * [Deploy the front image](gcp/deployment_front.md)
* * [Load balance multiple cluster](gcp/mci.md)
* * [Horizontal autoscaler](k8s/hscaler.md)
* * [Debugging](gcp/debug.md)
* [Deploying our Application with Helm](helm/intro.md)
* * [Helm architecture](helm/architecture.md)
* * Helm on Minikube
* * * [Creating our bobba-chart](helm/bobba-charts.md)
* * * [Creating a subchart for the database](helm/bobba-db.md)
* * * [Deploying our Helm on GCP](helm/gcp.md)
* * * [Useful Helm actions](helm/actions.md)

## Objectives[¶](#objectives)

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
- Deploy a load balancer for multiple cluster
- Be able to know implement a monitoring process by using stackdriver & grafana.
- Create Cron task on GCloud 
```

## Go deeper

If you wish to go deeper into learning Kubernetes I recommend you to take a look at this website [Kubernetes by example](http://kubernetesbyexample.com/)

## Enjoy ! 🎎
