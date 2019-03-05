---
layout: single
classes: wide
title: Services 🔑
sidebar:
  nav: "docs"
---

A service is a microservice which is grouping numbers of pods by pods's label. Services allow you to define the access policy to your pods

Services are cheaps which mean that you can have as many services as you want in your Cluster as they're not consuming a lot of hardware resources.

## Usefulness of Service[¶](#usefulness-of-service)

As per CoreOS definition of Service:

> Services provide important features that are standardized across the cluster: load-balancing, service discovery between applications, and features to support zero-downtime application deployments.

Futhermore pod are ephemeral. Which mean that it's not reliable to use the pod's IP in order to connect your stack between them. Service are micro services that support zero-downtime deployments as they have an unlimited lifespans.

## Type of services[¶](#type-of-services)

Kubernetes provide many types of services. In this section we'll explore all of them.

### ClusterIP[¶](#cluserip)

Expose the service on a Cluster internal IP. This mean that by choosing this kind of service will make your pods reachable within the Cluster

### NodePort[¶](#nodeport)

This expose the service on each Node's IP static port. A ClusterIP will be create automatically to which the NodePort service will be created.

E.g: 3 Nodes -> 3 NodeIP created binded to the Cluster.

### LoadBalancer[¶](#loadbalancer)

This type offer the same functionality as the NodePort. It add an other layer where your services are exposed to an external Load Balancer provider. This mean that the external load balancer will be in charge of handling the traffic to your services 

### ExternalName[¶](#externalname)

As per Kubernetes's documentation: 

> Maps the service to the contents of the externalName field (e.g. foo.bar.example.com), by returning a CNAME record with its value. No proxying of any kind is set up. This requires version 1.7 or higher of kube-dns.

## Relationship with Kube-proxy[¶](#relationship-with-kube-proxy)

Each Kubernete's Node is running a Kube-proxy (see the architecture page). Each of these kube-proxy is responsible for implementing a form of Virtual IP for the ```Services``` of all type aside from the ```ExternalName```.

As this part can be really tricky to explained I'll suggest you to read this article which the author Mark Betz, did a tremendous work regarding this subject [Networking services, Kube-proxy](https://medium.com/google-cloud/understanding-kubernetes-networking-services-f0cb48e4cc82)

## Resources[¶](#resources)

[Basic of networking](https://www.digitalocean.com/community/tutorials/an-introduction-to-networking-terminology-interfaces-and-protocols)

[IP Addresses](https://www.digitalocean.com/community/tutorials/understanding-ip-addresses-subnets-and-cidr-notation-for-networking)

[Netorking introduction in Kubernetes](https://medium.com/google-cloud/understanding-kubernetes-networking-pods-7117dd28727)

[Learning routing tables](https://www.youtube.com/watch?v=g8eP4fhrx3I)

[Kube Proxy, Services networking](https://medium.com/google-cloud/understanding-kubernetes-networking-services-f0cb48e4cc82)

[Which service should I choose](https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0)