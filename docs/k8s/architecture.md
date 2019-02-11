# Architecture

In this page we will take a quick look at the internal architecture of Kubernetes. This will allow us to understand the different components of this containarized deployment services.

## Type of architecture

Kubernetes is a master-slave architecture centric service. A master-slave architecture is a kind of architecture where the main process has the control of all the other child process.

Within Kubernetes we have a set of **master components** which will give instruction to the **nodes components**. The **master components** are regarded as the master and the **nodes components** act as the slaves.

Below is big picture of the architecture of Kubernetes

![Kubernetes architecture](../img/architecture.png)

## Components

Understanding kubernetes components will allow you to better understand what happened behind. First we will see the **master components**. Later on the **node component** which are the slaves

### Master components

#### Kube api-server

The API server is one of the most important components in the Kubernetes architecture. Indeed this API server communicate with every component of the **master components**.

Function wise the api-server is a REST api where it's simple functionnality is to validate, update or even send the datas that are send to the api server.

As the API server is tightly coupled to every components of Kubernetes it use an external database named Etcd for saving the state of the cluster.

To see how crucial is the api-server please take a look at this schema:

![pod flow](../img/pod_flow.png)

Create Pod Flow: Source: heptio.com

#### Etcd

Etcd is a high performance key value database written in Go by etcd. It's only used by the api-server. The database is used for saving the cluster state, the configuration of your pods, services and the metdata 

For more information about etcd please visit the etcd repo [Etcd repo](https://github.com/etcd-io/etcd) or you could also take a look at this article: [How kubernetes use etcd](https://matthewpalmer.net/kubernetes-app-developer/articles/how-does-kubernetes-use-etcd.html)

#### Controller manager

The controller manager is a daemon (list a background process) that is executed and use by Kubernetes. It's actually a controller that watches the state of the cluster throughout api server. If any difference happened within the state of the cluster it'll try to move the current state to the desired state. Executed as a single process it's actually a list of controller that you can get the list [over here](https://github.com/kubernetes/kubernetes/blob/7c75723867e9e431da323b8cc410bab928cada17/cmd/kube-controller-manager/app/controllermanager.go#L330)

#### Scheduler

The scheduler is an other component of the **master components**. It's role is to watch for unscheduled pod. When this happened the scheduler will try to find the proper node to bind it depending on several criteria.

The binding happened through the ```/binding``` subresources. For more information regarding how the scheduling work a fantastic article is available [here](https://kublr.com/blog/implementing-advanced-scheduling-techniques-with-kubernetes/)

#### Cloud controller manager

The cloud controller manager is a set of controller that embeds cloud specific control loops. This set of controller allow cloud provider to make their own set of controller while satisfying the cloud controller interfaces. (E.g, a foo cloud provider implement their own node controller)

A list of available controller is available here: [Cloud controller manager](https://github.com/kubernetes/kubernetes/blob/6671d2556f1af67e703c329b1186896d7c6f9f4d/cmd/cloud-controller-manager/app/controllermanager.go#L270)

### Child components

## Resources used. Big kudos for their work.

[Description of kubernetes architecture](https://elastisys.com/wp-content/uploads/2018/01/kubernetes-ha-setup.pdf?x83281)

[Kubernetes documentation](https://kubernetes.io/docs/concepts/overview/components/)

[Learning kubernetes architetcture](https://medium.com/jorgeacetozi/kubernetes-master-components-etcd-api-server-controller-manager-and-scheduler-3a0179fc8186)

[Scheduler detailed explanation](https://kublr.com/blog/implementing-advanced-scheduling-techniques-with-kubernetes/)





