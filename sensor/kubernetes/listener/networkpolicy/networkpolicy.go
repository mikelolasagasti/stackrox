package networkpolicy

import (
	"time"

	pkgV1 "bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/listeners"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/protoconv"
	"bitbucket.org/stack-rox/apollo/sensor/kubernetes/listener/watchlister"
	"k8s.io/api/networking/v1"
	"k8s.io/client-go/rest"
)

var logger = logging.LoggerForModule()

// WatchLister implements the WatchLister interface
type WatchLister struct {
	watchlister.WatchLister
	eventC chan<- *listeners.EventWrap
}

// NewWatchLister implements the watch for network policies
func NewWatchLister(client rest.Interface, eventC chan<- *listeners.EventWrap, resyncPeriod time.Duration) *WatchLister {
	npwl := &WatchLister{
		WatchLister: watchlister.NewWatchLister(client, resyncPeriod),
		eventC:      eventC,
	}
	npwl.SetupWatch("networkpolicies", &v1.NetworkPolicy{}, npwl.resourceChanged)
	return npwl
}

func (npwl *WatchLister) resourceChanged(networkPolicyObj interface{}, action pkgV1.ResourceAction) {
	networkPolicy, ok := networkPolicyObj.(*v1.NetworkPolicy)
	if !ok {
		logger.Errorf("Object %+v is not a valid network policy", networkPolicyObj)
		return
	}

	npwl.eventC <- &listeners.EventWrap{
		SensorEvent: &pkgV1.SensorEvent{
			Id:     string(networkPolicy.UID),
			Action: action,
			Resource: &pkgV1.SensorEvent_NetworkPolicy{
				NetworkPolicy: protoconv.ConvertNetworkPolicy(networkPolicy),
			},
		},
	}
}
