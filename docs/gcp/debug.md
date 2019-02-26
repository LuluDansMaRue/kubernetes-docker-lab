## Debugging your Kubernetes cluster

During any step of our adventure with Minikube & GCP we gave you some sample command regarding how to debug your project. Below is the list of command that can be useful to debug your Kubernetes environment.

### Watch for your deployment

You can use a watcher command which will allow you to see the status of your deployment

```shell
kubectl rollout status deployments.v1.apps/<deployment_id>
```

> I guess it should be the same for the replicaSets and other kind of Controller & API Services

### Debug your pod

You can debug in 2 ways. First by looking at the events of the pods itself which can show e.g: if the container has been successfully created or not

```shell
kubectl describe pod <pod_id>
```

In the other hand you could also get the logs of your pod

```shell
kubectl logs -p <pod_id>
```

### Debug your services

Like debugging your pod you can query the describe API for debugging your service like so

```shell
kubectl describe service <service_id>
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