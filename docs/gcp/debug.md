## Debugging your Kubernetes cluster

During any step of our adventure with Minikube & GCP we gave you some sample command regarding how to debug your project. Below is the list of command that can be useful to debug your Kubernetes environment.

### Watch for your deployment

You can use a watcher command which will allow you to see the status of your deployment

```shell
kubectl rollout status deployments.v1.apps/<deployment_name>

# Response type
deployment <deployment_name> successfully rolled out
```

> I guess it should be the same for the replicaSets and other kind of Controller & API Services

### Debug your pod

You can debug in 2 ways. First by looking at the events of the pods itself which can show e.g: if the container has been successfully created or not

```shell
kubectl describe pod <pod_id>

# Look for the events section
```

In the other hand you could also get the logs of your pod

```shell
kubectl logs -p <pod_id>
```

You can also run 

```shell
kubectl get pods

# Response type

NAME                         READY     STATUS    RESTARTS   AGE
bobba-api-7894d6765d-8lm7q   2/2       Running   0          2h
bobba-api-7894d6765d-ch9qn   2/2       Running   0          2h
bobba-vue-54bb88c7bb-466cb   1/1       Running   0          1d
bobba-vue-54bb88c7bb-n4fpd   1/1       Running   0          1d
```

and look for the error such as CrashLoopBackoff

### Debug your services

Like debugging your pod you can query the describe API for debugging your service like so

```shell
kubectl describe service <service_id>

# Response type
Name:                     bobba-api
Namespace:                default
Labels:                   <none>
Annotations:              <none>
Selector:                 app=bobba-api,tier=backend
Type:                     NodePort
IP:                       <ip>
Port:                     <unset>  80/TCP
TargetPort:               8000/TCP
NodePort:                 <unset>  31320/TCP
Endpoints:                <pod ip>:8000,<pod ip>:8000
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

### Connect to your nodes

You can connect to your Nodes cluster by doing this command. This allow you to see the logs of the kubelet and internal Kubernetes components

First you have to find your instances with this command

```shell
gcloud compute instances list
```

```shell
gcloud compute ssh <cluster_instance_pool_name> --zone <zone of your cluster>

# e.g gcloud compute ssh gke-bubbletea-cluster-default-pool-ABCD --zone us-central1-a
```

**Note**: If you don't find your instance make sure that you are using the right project. Find your project by using this command

```shell
gcloud projects list
```

then set the right project 

```shell
gcloud config set project <project_id>
```