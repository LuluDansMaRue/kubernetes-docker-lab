---
layout: single
classes: wide
title: CronJob
sidebar:
  nav: "docs"
---

In this article we will take a look at the CronJob in Kubernetes.

## A special type of Job

In the OS world a CronJob is a task that is command by the CronTab a small process which allow you to define script that need to be executed throughout the time.

Within Kubernetes CronJob is a special kind of Job. Indeed Job are pods that are executed in order to do a task and when they're completed jobs are terminated as well as the pods.

When it come to the CronJob the behavior is the same with a small difference which is that this job will be re executed more through the time.

## Example

Within our example deployment we're giving an example into how to deploy a CronJob with kubernetes by using MiniKube an GCP. take a look

- [Deploy a Cron with Minikube](deployment/cron.md)