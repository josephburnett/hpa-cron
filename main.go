package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	gce "cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/monitoring/v3"
)

var (
	namespace = flag.String("namespace", "default", "A namespace for the metric.")
	name      = flag.String("name", "", "A name for the metric.")
	value     = flag.Float64("value", 0.0, "The value to emit.")
)

func main() {
	flag.Parse()
	if err := emit(*namespace, *name, *value); err != nil {
		panic(err)
	}
	log.Printf("Emitted hpacron_%v_%v = %v.", *namespace, *name, *value)
}

func emit(namespace, name string, value float64) error {
	sd, err := monitoring.New(oauth2.NewClient(context.Background(), google.ComputeTokenSource("")))
	if err != nil {
		return err
	}
	now := time.Now().Format(time.RFC3339)
	resourceLabels := getResourceLabelsForNewModel()
	projectName := fmt.Sprintf("projects/%s", resourceLabels["project_id"])
	monitoredResource := "k8s_cluster"
	metricName := fmt.Sprintf("hpacron_%v_%v", namespace, name)
	metricLabels := map[string]string{}
	request := &monitoring.CreateTimeSeriesRequest{
		TimeSeries: []*monitoring.TimeSeries{
			{
				Metric: &monitoring.Metric{
					Type:   "custom.googleapis.com/" + metricName,
					Labels: metricLabels,
				},
				Resource: &monitoring.MonitoredResource{
					Type:   monitoredResource,
					Labels: resourceLabels,
				},
				Points: []*monitoring.Point{{
					Interval: &monitoring.TimeInterval{
						EndTime: now,
					},
					Value: &monitoring.TypedValue{
						DoubleValue: &value,
					},
				}},
			},
		},
	}
	_, err = sd.Projects.TimeSeries.Create(projectName, request).Do()
	return err
}

func getResourceLabelsForNewModel() map[string]string {
	projectId, _ := gce.ProjectID()
	location, _ := gce.InstanceAttributeValue("cluster-location")
	location = strings.TrimSpace(location)
	clusterName, _ := gce.InstanceAttributeValue("cluster-name")
	clusterName = strings.TrimSpace(clusterName)
	return map[string]string{
		"project_id":   projectId,
		"location":     location,
		"cluster_name": clusterName,
	}
}
