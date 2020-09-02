# hpa-cron
Example of how to scale HPA with a CronJob.

## Prerequisites

* [ko](https://github.com/google/ko)
* [Stackdriver Adapter](https://github.com/GoogleCloudPlatform/k8s-stackdriver/tree/master/custom-metrics-stackdriver-adapter)

## Usage

* `ko apply -f example.yaml`

## Watch

* `watch 'kubectl get hpa ; echo ; kubectl get cronjob ; echo ; kubectl get pods'`
