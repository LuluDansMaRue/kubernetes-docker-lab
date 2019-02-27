## Multiple cluster ingress

In today's article we're going to take a look at how to LoadBalance the request to 2 clusters. Why 2 clusters ? Imagine that our app is highly popular in the us but also in asia. 

Thus with only 1 cluster in the US taking all the datas it might be suitable to better to redirect the traffic to a cluster in the asia region.

## Creating an other cluster

Let's create an other cluster. This cluster is going to be in asia-east1-c (taiwan). 
Firstable we need to recreate the storage credentials config for the bobba-api

...

Now let's move on and redeploy our front & api...

## Kubemci

get credentials of cluster us-central1-a

```shell
KUBECONFIG=mci/mcikubeconfig gcloud container clusters get-credentials --zone=us-central1-a bubbletea-cluster
```

get credentials of cluster asia-east1-c

```shell
KUBECONFIG=mci/mcikubeconfig gcloud container clusters get-credentials --zone=us-central1-a bobba-cluster-asia	
```

