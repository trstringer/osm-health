package main

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/openservicemesh/osm-health/pkg/ingress"
	"github.com/openservicemesh/osm-health/pkg/kubernetes/pod"
)

func newIngressToPodCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "to-pod DESTINATION_POD",
		Short: "Checks ingress to a given Kubernetes pod",
		Long:  `Checks ingress to a given Kubernetes pod`,
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.Errorf("missing DESTINATION_POD parameter")
			}
			log.Info().Msgf("Checking Ingress to Pod %s", args[0])

			client, err := pod.GetKubeClient()
			if err != nil {
				return err
			}

			dstPod, err := pod.FromString(args[0])
			if err != nil {
				return errors.New("invalid DESTINATION_POD")
			}

			ingress.ToDestinationPod(client, dstPod)

			return nil
		},
		Example: `TODO`,
	}
	return cmd
}
