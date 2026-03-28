package helpers

import (
	"context"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func IsManagedByPropagationController(labels map[string]string) bool {
	labelValue := labels[ManagedByLabel]
	return labelValue == ManagedByValue
}

func hasPropagationEnabledAnnotation(annotations map[string]string) bool {
	annotationValue := annotations[PropagationEnableAnnotationKey]
	return strings.EqualFold(annotationValue, "true")
}

func getPropagationNamespacesFromAnnotations(annotations map[string]string) []string {
	if annotationValue, ok := annotations[PropagationNamespaceAnnotationKey]; !ok {
		return nil
	} else {
		return strings.Split(annotationValue, ",")
	}
}

func getObjectIfExists(ctx context.Context, client ctrlclient.Client,
	objKey ctrlclient.ObjectKey, obj ctrlclient.Object) (bool, error) {
	if err := client.Get(ctx, objKey, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func propagatedResourceLabels(sourceNamespace string) map[string]string {
	return map[string]string{
		ManagedByLabel:    ManagedByValue,
		AppNamespaceLabel: sourceNamespace,
	}
}
