# hpa-cron
Example of how to scale HPA with a CronJob.

## Prerequisites

* [ko](https://github.com/google/ko)
* [Stackdriver Adapter](https://github.com/GoogleCloudPlatform/k8s-stackdriver/tree/master/custom-metrics-stackdriver-adapter)

## Usage

**`ko apply -f example.yaml`**

HPA will scale on an hourly schedule **and** CPU utilization (whichever is greater).

`watch 'kubectl get hpa ; echo ; kubectl get cronjob ; echo ; kubectl get pods'`

==>

```
Every 2.0s: kubectl get hpa ; echo ; kubectl get cronjob ; echo ; kubectl get pods                                                                                                                                                                             josephburnett.waw.corp.google.com: Wed Sep  2 20:51:54 2020

NAME    REFERENCE          TARGETS             MINPODS   MAXPODS   REPLICAS   AGE
nginx   Deployment/nginx   1/1 (avg), 0%/60%   1         20        10         9m41s

NAME        SCHEDULE     SUSPEND   ACTIVE   LAST SCHEDULE   AGE
scale-in    1 * * * *    False     0        <none>          9m42s
scale-out   50 * * * *   False     0        116s            9m42s

NAME                         READY   STATUS      RESTARTS   AGE
nginx-5754944d6c-4tww4       1/1     Running     0          66s
nginx-5754944d6c-7jhw7       1/1     Running     0          97s
nginx-5754944d6c-9fxwk       1/1     Running     0          82s
nginx-5754944d6c-gtksp       1/1     Running     0          97s
nginx-5754944d6c-hj5nn       1/1     Running     0          66s
nginx-5754944d6c-lxcq6       1/1     Running     0          82s
nginx-5754944d6c-mkx2f       1/1     Running     0          82s
nginx-5754944d6c-pmfc2       1/1     Running     0          97s
nginx-5754944d6c-sl878       1/1     Running     0          82s
nginx-5754944d6c-w8vbh       1/1     Running     0          9m42s
scale-out-1599072600-v4wjx   0/1     Completed   0          109s
```

This is what it looks like over time:

![](example.png)
