## Ingress 🧠

As we saw earlier we'd create our API service based on the NodePort type. However it's not suitable to use each Node's IP as they are ephemeral. In order to resolve this issue we could use an API Object that will manage the external access of the services used by a Cluster.

Here come Ingress, an API Object that allow us to define access to our services from a Single IP. The Ingress controller will be responsible for fullfilling the [Ingress](https://www.webopedia.com/TERM/I/ingress_traffic.html) by spawing a LoadBalancer and other network systems.

If you don't know what a LoadBalancer is. Please check this great article about LoadBalancing [Digital Ocean Loadbalancer explanation](https://www.digitalocean.com/community/tutorials/what-is-load-balancing)

When generating a default Ingress configuration. Ingress come with a ```ephemeral``` IP. Which mean that your single endpoint that manage the services can change at any time.

However with GCP we have the possibility to create **Static IP** which is free if you used it. Let's firstable generate our Static IP

## Generate static IP ⚡

Before allocating the static IP with Google we need to choose which kind of IP we need. Within GCP there are **2 types** of IP:

- Regional: Use for an Instance or a Network Load Balancer
- Global: IP for HTTP(S), SSL Proxy, TCP Proxy Load Balancer

In our case we'll use Global IP as we're reserving an HTTP address.

With the GCloud SDK run the command below which is going to create an IP of type global.

```ssh
gcloud compute addresses create bobba-api-ip --global
```

Now check that your IP has been created with the command

```ssh
gcloud compute addresses describe bobba-api-ip --global
```

You should get a result like below

```yaml
address: <ip>
creationTimestamp: '2019-02-21T02:21:45.465-08:00'
description: ''
id: 'ID'
ipVersion: IPV4
kind: compute#address
name: bobba-api-ip
networkTier: PREMIUM
selfLink: https://www.googleapis.com/compute/v1/projects/<project_name>/global/addresses/bobba-api-ip
status: <status>
users:
- https://www.googleapis.com/compute/v1/projects/<project_name>/global/forwardingRules/k8s-fw-default-bobba-api-ingress--c37fc980ac025b96
```

All right we have our IP Address let's move on and create our Ingress service

## Ingress configuration ⚒️

Like all kubernetes configuration creating an Ingress configuration is straightforward. Below is the configuration file that is location of the file on the project ```gcp/ingress/balancing.yml```. Let's describe it

```yaml
# Version
apiVersion: extensions/v1beta1
# Type of services
kind: Ingress
# Meta data such as name
metadata:
  name: bobba-api-ingress
  # Here we're defining the static IP to use
  annotations:
    # Hey it's the name of our ip address that we created earlier
    kubernetes.io/ingress.global-static-ip-name: bobba-api-ip
spec:
  rules:
    - http:
        paths:
        # Rule to define which service should we target
        # You can add multiple services to target
        - path: /*
          backend:
            serviceName: bobba-api
            # Expose port by the service. E.g if port: 80 in service then use 80 for servicePort
            servicePort: 80
```

Now that we have our configuration we can deploy our ingress configuration to GCP by running this command below:

```shell
kubectl create -f gcp/ingress/balancing.yml
```

You can then check your ingress by running the command 

```shell
kubectl describe ingress
```

> Now it's going to take a bit of time in order for the configuration to be created and spread. Now, take a break, take a frappuccino, bubble tea, watch some stuff on Netflix or just go outside it might be sunny. 🌞

Once you did all of that you should be able to use the Static IP for targeting the API. E.g

```shell
curl <IP>

# you should get a response like below
Hello bubble tea server !!
```

### Save your address somewhere we'll reuse it for building the front image

#### Now let's move on and create our [front image](front.md)