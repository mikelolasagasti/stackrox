package queue

import (
	"sync"

	"bitbucket.org/stack-rox/apollo/central/sensorevent/store"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
)

var (
	log = logging.LoggerForModule()

	// Here we configure the actions we actually want to enqueue. New action types should be configured here
	// as unrecognized actions will also be ignored (since we can't deduplicate them).
	handledActions = map[v1.ResourceAction]bool{
		v1.ResourceAction_UNSET_ACTION_RESOURCE: false,
		v1.ResourceAction_PREEXISTING_RESOURCE:  true,
		v1.ResourceAction_CREATE_RESOURCE:       true,
		v1.ResourceAction_REMOVE_RESOURCE:       true,
		v1.ResourceAction_UPDATE_RESOURCE:       true,
		v1.ResourceAction_REFRESH_RESOURCE:      false,
		v1.ResourceAction_DRYRUN_RESOURCE:       false,
	}
)

// EventQueue provides an interface for a queue that stores DeploymentEvents.
type EventQueue interface {
	Push(*v1.SensorEvent) error
	Pull() (*v1.SensorEvent, error)
	Load(clusterID string) error
	Count() int
}

// NewPersistedEventQueue returns a new instance of an EventQueue.
func NewPersistedEventQueue(eventStorage store.Store) EventQueue {
	pen := &persistedEventQueue{
		eventStorage: eventStorage,

		seqIDQueue:   make([]uint64, 0),
		depIDToSeqID: make(map[string]uint64),
	}
	return pen
}

// persistedEventQueue is an implementation of EventQueue that persists items in the queue in the db
// provided.
type persistedEventQueue struct {
	eventStorage store.Store

	mutex        sync.Mutex
	seqIDQueue   []uint64
	depIDToSeqID map[string]uint64
}
