# Services

A service is a microservice which is a grouping of pods defined by service's label. It allow you to define the policy on how to access to the pods.

Services are cheaps which mean that you can have as many services as you want in your Cluster.

# What does Service provide

As per CoreOS definition of Service:

> Services provide important features that are standardized across the cluster: load-balancing, service discovery between applications, and features to support zero-downtime application deployments.

Moreover pod's lifetime are ephemeral. Which mean that you can't count at using pod's IP in order to connect your stack between them. Service are micro services that support zero-downtime deployments. Meaning that they have an unlimited lifespans

# Type of services

Kubernetes provide many types of services. In this section we'll explore all of them.

- ClusterIP: Exposes the service on a Cluster internal IP. This mean that by choosing this kind of service will make your pods available within the Cluster

- NodePort: This expose the service on each Node's IP at a static port. A ClusterIP will be create to which the NodePort service will be route is automatically created. The kube-proxy will manage which pod to transmit the packet...

- LoadBalancer: This type offer the same functionality as the NodePort. It add an other layer where your services are exposed to an external Load Balancer provider. This mean that the external load balancer will be in charge of handling the traffic to your services 

- ExternalName: As per Kubernetes's documentation: > Maps the service to the contents of the externalName field (e.g. foo.bar.example.com), by returning a CNAME record with its value. No proxying of any kind is set up. This requires version 1.7 or higher of kube-dns.

# Relationship with Kube-proxy

Each Kubernete's Node is running a Kube-proxy (see the architecture page). Each of these kube-proxy is responsible for implementing a form of Virtual IP for the ```Services``` of all type aside from the ```ExternalName```.

As this part can be really lengthy to explained I'll suggest you to read this article which did a tremendous work regarding this subject [Networking services, Kube-proxy](https://medium.com/google-cloud/understanding-kubernetes-networking-services-f0cb48e4cc82)

# Creating services

Creating a services is simple. Please check the example [over here](deployment/service.md)

# Resources

[Basic of networking](https://www.digitalocean.com/community/tutorials/an-introduction-to-networking-terminology-interfaces-and-protocols)
[IP Addresses](https://www.digitalocean.com/community/tutorials/understanding-ip-addresses-subnets-and-cidr-notation-for-networking)
[Netorking introduction in Kubernetes](https://medium.com/google-cloud/understanding-kubernetes-networking-pods-7117dd28727)

[Learning routing tables](https://www.youtube.com/watch?v=g8eP4fhrx3I)

[Kube Proxy, Services networking](https://medium.com/google-cloud/understanding-kubernetes-networking-services-f0cb48e4cc82)

[Which service should I choose](https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0)