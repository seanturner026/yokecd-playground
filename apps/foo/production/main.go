package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/seanturner026/yokecd-playground/pkg/k8s/deployment"
	"github.com/seanturner026/yokecd-playground/pkg/k8s/service"
	"github.com/yokecd/yoke/pkg/flight"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		release   = flight.Release()
		namespace = flight.Namespace()
		labels    = map[string]string{"app": release}
	)

	resources := []flight.Resource{
		deployment.Create(deployment.Config{
			Name:      release,
			Namespace: namespace,
			Labels:    labels,
			Replicas:  2,
		}),
		service.Create(service.Config{
			Name:       release,
			Namespace:  namespace,
			Labels:     labels,
			Port:       80,
			TargetPort: 80,
		}),
	}

	return json.NewEncoder(os.Stdout).Encode(resources)
}
