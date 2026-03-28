package helpers

const (
	PropagationEnableAnnotationKey    = "config-propagator.sanadhis.com/propagate"
	PropagationNamespaceAnnotationKey = "config-propagator.sanadhis.com/target-namespaces"

	ManagedByLabel    = "app.kubernetes.io/managed-by"
	ManagedByValue    = "config-propagator"
	AppNamespaceLabel = "app.kubernetes.io/source-namespace"
)
