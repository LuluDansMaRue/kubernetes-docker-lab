---
layout: single
classes: wide
title: Helm
sidebar:
  nav: "docs"
---

## Context

Kubernetes is a huge platform that's complex. This required as you saw in the previous part of this learning guide an understanding of how is working Kubernetes.

Helm enabled you to interact with Kubernetes with much less headache.

## Helm

Helm is a package manager for Kubernetes. It allow you to download several kind of existing packages called ```chart``` which are representing the type of Deployment that you want to do.

As an example if you want to deploy a MongoDB server into Kubernetes. Helm could have a chart available which you could use out of the box without any headache.

As it's a package manager. It allows you to do several common commands that you could found on package manager such as NPM, brew, Chocolatery, APT

## Impact of Helm in our example

In our example we could see that we'd deploy our service, then our containers by using a bunch of deployment object available within the Kubernetes environment and then an ingress resources.

Nonetheless maintaining all of those files can be paintful. Futhermore upgrading, rolling back one or more of these files can be a tight difficult. And this is where Helm is shining.

As a package manager for Kubernetes. Helm allow you to easily maintain your Kubernetes cluster with only a few commands. So let's take see how can we adapt our configuration to Helm

## Table of contents

* [Architecture](architecture.md)
* Minikube example
* * [Creating our chart for the bobba project](bobba-charts.md)
* * [Creating our chart for the database](bobba-db.md)
* * [Useful actions & Wrap up](actions.md)

## Resources

[Helm documentation](https://helm.sh/docs/)

[Bitmani helm's guide](https://docs.bitnami.com/kubernetes/how-to/deploy-application-kubernetes-helm/)

[CNCF video](https://www.youtube.com/watch?v=vQX5nokoqrQ)
