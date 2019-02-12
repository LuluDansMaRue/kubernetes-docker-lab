# Introduction

In this article we'll talk about how to deploy an application on Kubernetes by taking a look at the list of options that Kubernetes is offering.

# Type of controller & API Objects use for deploying a Kubernetes app

## Controller

As of Feburary 2019, Kubernetes provide several types of deployments named Controller. 

```
These controller are part of the kube-controller-manager
```

- ReplicaSet
- ReplicationController
- StatefulSets
- DaemonSet
- TTL Controller (see documentation as it's in alpha)
- CronJob

## API Object

API Object are usually a wrapper around those controller. It add additional functionnality to them. Below is the list of available API Object available in Kubernetes:

- Deployments
- Job
- CronJob

# Explanation of each controllers

First of all let's undersand in a few lines what's the purposes of these roles.

## ReplicaSet

ReplicaSet is a type of deployment that aim to surpass the previous ReplicationController. It can be used if you need to maintain a stable amount of replica Pods running at anytime.

This mean that the ReplicaSet controller will monitor the statuses of each pods and create a new one if one of them happened to failed or lost.

Note that ReplicaSet is also ship with additional functionnalities. I suggest you to check the Kubernetes documentation as well as the **Explanation of ReplicaSet article** on the resources section

<p align="center">
  <img src="../img/replicaset_schema.png" alt="drawing" width="450"/>  
  <p align="center"><b>ReplicaSet schema, Jiamin Ning, Medium article.</b></p>
</p>

## ReplicationController

ReplicationController is actually an older version of ReplicaSet. While their purposes are the same they work differently ReplicationController does not support set based selector. It's therefore recommended by Kubernetes to use the **ReplicaSet** instead of the **ReplicationController**

## StatefulSets

StatefulSets is a kind of controller that has the same capabilities as the Deployment. However one major difference is that the StatefulSets has a build-in storage for each replicas. This kind of deployment is useful for statful application an online photoshop application where it need to store multiple files and that each update shouldn't erase associated datas.

Note: You could setup a database with Kubernetes by using this option. However database resources can grow rapidly which a container is at some point limited. However for demo purposes this can be use... (own opinion)

## DaemonSets

A DaemonSet is a kind of controller that will deploy a pod into each Node that exist on the cluster. If a node get removed then the DaemonSet pod will get garbage collected which mean that the DaemonSet scale automatically based on the number of Nodes that are existing.

Similarly if a Node get's create, the DaemonSet will automatically add a pod into this new Node.

This kind of deployment can be useful for many type of application such as **logs collections**, **monitoring tools**

From google's doc:

```
In a simple case, one DaemonSet, covering all nodes, would be used for each type of daemon. A more complex setup might use multiple DaemonSets for a single type of daemon, but with different flags and/or different memory and cpu requests for different hardware types.
```

### API Objects

## Deployment

Deployment is an API Object that leverage the ReplicaSet controller by bringing  additional feature such as:

- Create a deployment to Rollout a ReplicaSet: Creating ReplicaSet in background and monitor the status to see if it succeed or not
- Declare new state of the pods by updating the PodTemplateSpec: This will allow Kubernetes to create a new type of ReplicaSet and to replace the old ones on the fly
- Rollabck to an earlier version of a Deployment: If anything happened you can still rollback your deployment
- Scale up the deployment
- Pause the deployment
- Cleanup older ReplicaSets that we don't need anymore

Thus it's recommended to use the Deployment instead of the 2 aboves directly;

## Job

A Job is a kind of controller that allow you to run a number of pods. It allow you to create several pod, tracks it's completion and set a number of pod that need to be succeed in order for the Job to be considered as complete. In case of a failure the job will recreate an additional pod.

Job can be use to run conccurent pod. (parallelism). A special kind of Job named CronJob exist and is based on the Job controller

## CronJob

CronJob is a special kind of controller that allow you to run a number of pod's replica at a certain amount of time. Like the Jobs the CronJob will attempt to recreate a numbers of pods if one of them fail. I suggest you to take a look at the CronJob section.


# Resources

[Explanation of ReplicaSet](https://medium.com/@jiamin_ning/build-your-first-kubernetes-service-with-replicaset-7c37d9be689c)

[Kubernetes ReplicaSet documentation](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)

[DaemonSet explanation by a GCP Developer](https://medium.com/google-cloud/kubernetes-run-a-pod-per-node-with-daemon-sets-f77ce3f36bf1)

[Job & CronJob explanation](https://codeblog.dotsandbrackets.com/one-off-kubernetes-jobs/)

[TTL Controller](https://kubernetes.io/docs/concepts/workloads/controllers/ttlafterfinished/)

