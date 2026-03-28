package helpers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func GetAllNamespaces(ctx context.Context, client ctrlclient.Client) ([]string, error) {
	namespaces := &corev1.NamespaceList{}

	if err := client.List(ctx, namespaces); err != nil {
		return nil, err
	}

	namespaceNames := make([]string, 0, len(namespaces.Items))
	for _, ns := range namespaces.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}

	return namespaceNames, nil
}

func VerifyNamespaceExists(ctx context.Context, client ctrlclient.Client, namespace string) (bool, error) {
	ns := &corev1.Namespace{}
	return getObjectIfExists(ctx, client, ctrlclient.ObjectKey{Name: namespace}, ns)
}
