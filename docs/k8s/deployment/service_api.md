# Services api

As explained on the services main articles. We need to define services that will group our pods in order to make them accessible.

In our case we want to make them accessible outside of the Cluster. As we saw earlier one of the type that allow us to do that is the ```NodePort```. This type allow us to expose our Node's API throught a ClusterIP which make them available through a single endpoint. This is an interesting kind as Nodes are mortal like pods.

Using NodePort allow the ```kube-proxy``` to handle this issue in an automatic way.

# Example

We're going to expose our bobba-api deployment by creating a service with the NodePort type. Let's do it

```yaml
kind: Service
apiVersion: v1
metadata:
  # This is the name of the service
  name: bobba-api
spec:
  # Type of service
  type: NodePort
  selector:
    # This should match the labels that defined your pods
    app: bobba-api
    tier: backend
  ports:
  # If no IP. Kubernetes will assign a random nodeport
  - nodePort: 31320
    protocol: TCP
    # port expose to the cluster
    port: 80
    # Port of the pods
    targetPort: 8000
```

Pretty simple isn't it ?

Create your service with this command

```shell
kubectl create -f k8s/services/api_service.yml
```

Now we need to check that our services is created and exposed our Nodes by running this command

```shell
kubectl get services
```

You should see the following outputs

![bobba service](../../img/bobba-api.png)

Now you should be able to access to your service. Use the command below in order to retrieve the minikube IP

```shell
minikube ip
```

Then you are able to target the backend with this url with ```<minikube_ip>:port/{route}```

# Drawbacks

While it does what we want it's far from being a safe solution. Moreover we also have theses drawbacks:

- One service per port
- Can only use port from 30000 - 32767
- If NodeIP change we need to deal with this
