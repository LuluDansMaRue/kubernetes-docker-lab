---
layout: single
classes: wide
title: Helm on GCP
sidebar:
  nav: "docs"
---

## Helm & GKE

Previously we'd deployed our app in Minikube with Helm. In order to see a real life example let's deploy our app into GCP. For this example we will install a secure Helm. 

## Installing Helm

In order to do that I suggest you to carefully follow the great article of [Jonathan Campos on this subject](https://medium.com/google-cloud/install-secure-helm-in-gke-254d520061f7)

By securely install Helm. We're leveraging the security of our Cluster. Indeed only users with signed keys will be able to access to execute Helm command in our Cluster.

## Deploying our application

With helm installed all of our operations is going to be use like so

```shell
helm <action> --tls
```

With helm we should be able to not modify anything regarding our deployments thus we should be able to use it right-away. First let's analyze our bobba-api

```shell
helm install --tls --debug --dry-run -f bobba-helm-chart/values-api-prod.yaml ./bobba-helm-chart
```

You should see 3 Kubernetes objects:

- A deployment (our bobba-api)
- A service (bobba-api-service)
- An ingress controller

If you didn't get any error let's deploy !

```shell
helm install --tls -f bobba-helm-chart/values-api-prod.yaml ./bobba-helm-chart
```

Let's see the full deployment by using the command

```shell
helm status --tls <release_name>
```

You should get an output like the image below

<p align="center">
  <img src="../img/helm_bobba_prod.png" alt="drawing" width="800"/>
</p>

Repeat the same steps for the front-end project. With the front project you should get these Kubernetes object:

- Deployment
- Service
- Loadbalancer

You should get the same output

<p align="center">
  <img src="../img/helm_front_prod.png" alt="drawing" width="800"/>
</p>

## Resources

[Installing helm on GKE](https://medium.com/google-cloud/installing-helm-in-google-kubernetes-engine-7f07f43c536e)

[Install secure Helm on GKE](https://medium.com/google-cloud/install-secure-helm-in-gke-254d520061f7)

[Ingress documentation](https://helm.sh/docs/)