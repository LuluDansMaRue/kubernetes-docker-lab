# Architecture

In this page we will take a quick look at the internal architecture of Kubernetes. This will allow us to understand the different components of this containarized deployment services.

## Type of architecture

Kubernetes is a master-slave architecture centric service. A master-slave architecture is a kind of architecture where the main process has the control of all the other child process.

Within Kubernetes we have a set of **master components** which will give instruction to the **nodes components**. The **master components** are regarded as the master and the **nodes components** act as the slaves

## Components

Understanding kubernetes components will allow you to better understand what happened behind. First we will see the **master components**. Later on the **node component** which are the slaves

### Master components

#### kube apiserver

The kube api server is a CRUD operations that allows operation on the Kubernetes API Object such as adding a configuration, validating it.

This API is the main entry point to the components of Kubernetes. Indeed it allow you the administrator to communicate to your cluster through kubectl.

It's also use internally by the component of the Node (kubelet) to check the node's status

#### etcd

Etcd is an external library that is used by Kubernetes and most particularly by the API server. Indeed kube api server is stateless which mean that nothing is saved on the kube apiserver. In order to store the state of every components of the cluster. API Server is using the etcd database library which will store the cluster data

#### controller-manager

#### scheduler

### Child components

## Resources

[Description of kubernetes architecture](https://elastisys.com/wp-content/uploads/2018/01/kubernetes-ha-setup.pdf?x83281)
[Kubernetes documentation](https://kubernetes.io/docs/concepts/overview/components/)
[Learning kubernetes architetcture](https://medium.com/jorgeacetozi/kubernetes-master-components-etcd-api-server-controller-manager-and-scheduler-3a0179fc8186)





