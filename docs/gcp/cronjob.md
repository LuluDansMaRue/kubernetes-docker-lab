---
layout: single
classes: wide
title: Adding cronjob to your cluster with cloudsql proxy
sidebar:
  nav: "docs"
---

Since we deployed cronjob on our Minikube cluster. You might be tempt to deploy cronjob to your GKE cluster. As we use the cloudsql-proxy to deploy our API it might be a good idea to use the same kind of deployment. Having a pod which is containing 2 containers. The cronjob & the cloudsql-proxy pod as a sidecar.

## Issues

However this solution has an issue. Indeed the job doesn't really support the sidecar feature yet. This mean that once your job is completed the job won't terminate the sidecar container which lead to the fact that the job's pod got never delete...

At this day (02/04/2019) this feature is currently being under consideration. See this issue [issue](https://github.com/kubernetes/kubernetes/issues/25908)

## Solution

In order to use the full capabilities of the cloudsql-proxy feature we will use the strategy list below: 

> Note: There are other kind of strategy that you could use... but this one seems to me the cleanest way to do.

This strategy consist of the following stpes:
- Deploying a ```single pod``` which run the cloudsql-proxy.
- Deploying a service which is not exposed outside of the cluster we'll only exposed it through the kube-dns (Kubernetes use the service name as the domain)
- CronJob will target the DNS instead of the localhost

> ⚠️ Warning: This solution has one drawback. This single pod will be on a node. As you remember node are mortal thus it might be possible that a cronjob might not be able to communicate with the cloudsql-proxy pod as the node has been killed or is being recreated.

## Implementation

### Deployment

Deploying the cloudsql-proxy deployment is pretty straightforward. Create a deployment yml file which you register your cloudsql-proxy container and expose the pod to ```3306```. An already ```cloudsql.yml``` exist in the ```gcp/deployment/cloudsql.yml```

A subtle difference should be note:

```yaml
command: ["/cloud_sql_proxy",
          "-instances=<project>:<zone>:<database_name>=tcp:0.0.0.0:3306",
          "-credential_file=/secrets/cloudsql/credentials.json"]
```

⚠️ In the command value we're exposing the database by bypassing the firewall which should be take with precaution.

### Service

In our case we only need access to the database from the cluster. Thus we don't need to expose the database outside of the cluster (i.e NodePort, Loadbalancer).

Furthermore for security purposes it's not a good idea to expose the pod outside of the world. Thus we're only going to expose our deployment to the kube-dns by creating a ```Headless service```. Below is the configuration which could be found here: ```gcp/deployment/cloudsql.yml```

```yaml
kind: Service
apiVersion: v1
metadata:
  name: cloudsql-proxy-svc
spec:
  clusterIP: None
  selector:
    app: cloudsql
    tier: proxy
  ports:
  - protocol: TCP
    port: 3306
    targetPort: 3306
```

By using clusterIP to none we're only referencing our service to the kube-dns. In return we can communicate to the deployment with the ```name``` of the deployment. For more details on headless service please check the [documentation](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services)

## Configure our cron

Like for our ```bobba-api``` deployment we're going to apply the same configuration to our cronjob with one difference. The configuration is already applicated to the cronjob files available in this folder ```gcp/cron/*.yml```

```yaml
env:
  - name: MYSQL_HOST
    value: cloudsql-proxy-svc
  - name: MYSQL_USER
    valueFrom:
      secretKeyRef:
        name: cloudsql-db-credentials
        key: username
  - name: MYSQL_PASSWORD
    valueFrom:
      secretKeyRef:
        name: cloudsql-db-credentials
        key: password 
```

Instead of passing the ```127.0.0.1``` IP address to the ```MYSQL_HOST``` we're using the ```DNS discovery``` system in order to communicate with our database.

## Let's deploy

As we have every required files let's deploy. First go to the ```gcp/deployment``` folder & run this command

```shell
kubectl apply -f cloudsql.yml
```

*Make sure that the pod is running before processing to the next steps*

Now let's deploy the cron. Move to the folder ```gcp/cron``` and run this command

```shell
kubectl apply -f .
```

Make sure that the cronjob has been deployed to your cluster by running this command

```shell
kubectl get cronjob
```

... Wait for around 10 minutes ⏰

And you should see a cron pod running ! (this should be the insertbobba pod)

If the job run successfully you've just learn how to deploy cronjob with cloudsql-proxy to your GKE cluster.

## Resources

[Headless service in Kubernetes](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services)

[Running CloudSQL Proxy as a standalone](https://github.com/GoogleCloudPlatform/kubernetes-engine-samples/issues/8)
