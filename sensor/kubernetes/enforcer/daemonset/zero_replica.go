package daemonset

import (
	"fmt"

	"github.com/stackrox/rox/generated/internalapi/central"
	pkgKub "github.com/stackrox/rox/pkg/kubernetes"
	"k8s.io/client-go/kubernetes"
)

// EnforceZeroReplica does nothing but err out, since we can't zero out daemon set replica counts.
func EnforceZeroReplica(client *kubernetes.Clientset, deploymentInfo *central.DeploymentEnforcement) (err error) {
	return fmt.Errorf("scaling to 0 is not supported for %s", pkgKub.DaemonSet)
}
