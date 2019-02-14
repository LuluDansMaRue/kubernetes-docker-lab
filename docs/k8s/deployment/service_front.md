# Services front

As explained on the services main articles. We need to define services that will group our pods in order to make them accessible.

In our case we want to make them accessible outside of the Cluster. As we saw earlier one of the type that allow us to do that is the ```LoadBalancer```. This type has the same property as the ```NodePort```. However we aren't limited by the ```port``` issue of the NodePort.

Indeed by using the ```NodePort``` we are restricted by the usage of which port to use. Which in production mode is problematic.

# Example (description of the existing yaml file)

We're going to expose our bobba-vue (front) deployment by creating a ```LoadBalancer``` service

```yaml
kind: Service
apiVersion: v1
metadata:
  # Name of your service
  name: bobba-vue
spec:
  # Type of service
  type: LoadBalancer
  selector:
    # this should match the labels define in your pods
    app: bobba-vue
    tier: frontend
  ports:
    # No need to precise a special port. Minikube / GCP will create a load balancer on it's own.. at least on minikube :D
  - protocol: TCP
    port: 80
    targetPort: 8080
```

Pretty simple isn't it ?

Create your service with this command

```shell
kubectl create -f k8s/services/front_service.yml
```

Now we need to check that our services is created and exposed our Nodes by running this command

```shell
kubectl get services
```

As **minikube doesn't get access to GCP*** we need to run the command ```minikube tunnel``` which will create our LoadBalancer

Now if you check the services again you should see an ```externalIP``` for our ```bobba-vue``` service !

![bobba vue](../../img/bobba-vue.png)

# Drawbacks

One drawbacks of this method is that you need to deploy a load balancer for each services which can get quite expensive if you have multiple services
