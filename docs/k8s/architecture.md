# Architecture

In this article we'll take a look at Kubernetes's architecture. This will enable us to understand how is working Kubernetes.

## Type of architecture

Kubernetes is highly based on the master-slave architecture. A master-slave architecture is a type of architecture where the ```master``` decide the actions that the ```slave``` will do.

Within Kubernetes the ```master``` is name ```master components```. Slaves components are named ```node components```.

There are many components within each of these layers. Below is a simplify presentation of Kubernetes's architecture.

<p align="center">
  <img src="../img/architecture.png" alt="drawing" width="800"/>
</p>
<p align="center"><b>Kubernetes architecture</b></p>

## Components

By understanding Kuberntes's components this will allow us to understand what's happening behind the scene.

### Master components

Master components are a set of components which are running within the Cluster. These master components is pretty much the brain of the Cluster and therefore are critical components to Kubernetes.

#### Kube-api-server

Kube-api-server is an critial components of the Kubernetes architecture. This API is use by every components of the **master components**. It's also use by the kubelet process which is located in the Node.

So what's behind. Kube-api-server is a simple REST API with it's main purposes is to validate, saving Kubernetes's cluster state. It's also use for saving metadata, configuration of your pods & services

As Kube-api-server is stateless. Kube-api-server is tightly coupled to an external database component also located in the master named etcd. See etcd section for more information.

In order to understand how crucial Kube-api-server is important. Please take a look at the schema below:

<p align="center">
  <img src="../img/pod_flow.png" alt="drawing" width="500"/>  
</p>
<p align="center"><b>Create Pod Flow: Source: heptio.com</b></p>

#### Etcd

Etcd is a high performance key value database written in Go. This database is only use by kube-api-server.

For more information about etcd please visit the etcd repo [Etcd repo](https://github.com/etcd-io/etcd) or you could also take a look at this article: [How kubernetes use etcd](https://matthewpalmer.net/kubernetes-app-developer/articles/how-does-kubernetes-use-etcd.html)

#### Controller manager

The controller manager is a daemon, a background process which is located on the master. The controller manager is actually a list of several controller launch in a single process.

What it does is in a few words. Watches for any changes on the state of the cluster by querying ```kube-api-server```. If any changes happened the controller manager will move the current state to the desired state. 

The list of controller use by the controller manager is available [over here](https://github.com/kubernetes/kubernetes/blob/6671d2556f1af67e703c329b1186896d7c6f9f4d/cmd/kube-controller-manager/app/controllermanager.go#L339)

Technically a controller these controllers are a set of control loop that are non terminating that regulates the state of the system.

**Kubernetes's controller definition:**

```
In Kubernetes, a controller is a control loop that watches the shared state of the cluster through the API server and makes changes attempting to move the current state towards the desired state. Examples of controllers that ship with Kubernetes today are the replication controller, endpoints controller, namespace controller, and serviceaccounts controller.
```

#### Scheduler

The scheduler is an other component of the **master component**. It's task is to watch for unscheduled pod. When an unscheduled pod exist, the scheduler will try to find the proper node and will try to to bind it depending on several criteria such as the resources and more...

Binding an unscheduled pod to a node happened through the ```/binding``` subresources. For more information regarding how the scheduling work a fantastic article is available [here](https://kublr.com/blog/implementing-advanced-scheduling-techniques-with-kubernetes/)

#### Cloud controller manager

The cloud controller manager is a set of controller of customize cloud specific control loops. This set of controller allow cloud provider to make their own set of controller while satisfying the cloud controller interfaces. (E.g, a foo cloud provider implement their own node controller)

A list of available controller is available here: [Cloud controller manager](https://github.com/kubernetes/kubernetes/blob/6671d2556f1af67e703c329b1186896d7c6f9f4d/cmd/cloud-controller-manager/app/controllermanager.go#L270)

### Node components

Node components refer to the components running into a Node. A Node is the second layer of Kubernetes (slave in the master-slave architecture) where your pods are running. Pods are controlled and monitored by Kubelet

#### Kubelet

The kubelet is a **critical component** in Kubernetes which is located on each Node. It's the main component of the Node. 

It's purpose is to watch continuously the state of the pods and to send these information back to the kube-api-server (monitoring). Kubelet is also responsible at taking the proper action if something happened to a pod e.g: restarting a pod when a pod is getting shutdown.

TL;DR Kubelet responsability are:

- Run the pods container on the right engine
- Restart the pod if failure happened
- Report status of the node and each pod to kube api-server
- Retrieve metrics from the pods container

<p align="center"> 
  <img src="../img/nodes.svg" alt="drawing" width="400"/>
</p>
<p align="center"><b>Representation of a Node. Kubernetes documentation</b></p>

#### Kube proxy

The Kube proxy is a component which it's purpose is to watch changes in the definition of the service of the pods in order to maintain the desired network configuration by querying kube-api-server. 

This component also expose the pod to the correct back-end which allow you to exposes the pods by manipulating the iptables and assigning IP addresses so that you can easily access to your pods

#### Container runtime

The container runtime is the engine use for running the container pod. E.g: docker, rkt and many more...

## Resources used

[Description of kubernetes architecture](https://elastisys.com/wp-content/uploads/2018/01/kubernetes-ha-setup.pdf?x83281)

[Kubernetes documentation](https://kubernetes.io/docs/concepts/overview/components/)

[Learning kubernetes architetcture](https://medium.com/jorgeacetozi/kubernetes-master-components-etcd-api-server-controller-manager-and-scheduler-3a0179fc8186)

[Scheduler detailed explanation](https://kublr.com/blog/implementing-advanced-scheduling-techniques-with-kubernetes/)

[Creating it's own controller](https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html)
