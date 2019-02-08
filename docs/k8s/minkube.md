# Minikube

Minikube is a local kubernetes environment which allow you to test your container within a local cluster

# Troubleshootings

## Stuck at connection timed out

If you get an ```ssh error...: operation timed out ```
Locate the ```hyperkit.pid``` and delete it.

Below is an example on OSX:

```shell
rm -rf ~/.minikube/machines/minikube/hyperkit.pid
```