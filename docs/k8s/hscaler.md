---
layout: single
classes: wide
title: Horizontal pod autoscaler
sidebar:
  nav: "docs"
---

## Autoscaler

Kubernetes provide a horizontal pod autoscaler which allow you to scale your deployment based on metrics such as CPU or custom metrics.

This allow you to scale your Kubernetes environment based on criteria that you'd define.

## Description

As per [Kubernetes's documentation](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#how-does-the-horizontal-pod-autoscaler-work)

> The Horizontal Pod Autoscaler is implemented as a control loop, with a period controlled by the controller manager’s --horizontal-pod-autoscaler-sync-period flag (with a default value of 15 seconds).

> During each period, the controller manager queries the resource utilization against the metrics specified in each HorizontalPodAutoscaler definition. The controller manager obtains the metrics from either the resource metrics API (for per-pod resource metrics), or the custom metrics API (for all other metrics).

TL;DR: The horizontal pod autoscaler is a control loop which is managed by the controller managers which queries the metrics and compare the resources utilization against these metrics's values.

Futhermore I highly suggest you to check the algorithm behind the horizontal pod autoscaler [here](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/#how-does-the-horizontal-pod-autoscaler-work)

## Example

It's actually pretty straightforward. By just running a command we could easily scale our deployment.

Let's do it with our ```bobba-api``` deployment

```shell
kubectl autoscale deployment bobba-api --min=10 --max=18 --cpu-percent=45
```

Wait a bit and you'll see that we have now only 10 pods. 

When the aveage's pod cpu increase above 45%. The autoscaler will add autoscale our deployment and add an other pod.
