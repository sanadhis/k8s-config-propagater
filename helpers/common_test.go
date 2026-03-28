package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ownerReferences(t *testing.T) {
	tests := map[string]struct {
		labels   map[string]string
		expected bool
	}{
		"managed by controller": {
			labels: map[string]string{
				ManagedByLabel: ManagedByValue,
			},
			expected: true,
		},
		"not managed by controller": {
			labels:   map[string]string{},
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := IsManagedByPropagationController(test.labels)
			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_propagationEnabled(t *testing.T) {
	tests := map[string]struct {
		annotations map[string]string
		expected    bool
	}{
		"has propagation enabled": {
			annotations: map[string]string{
				PropagationEnableAnnotationKey: "true",
			},
			expected: true,
		},
		"has propagation enabled ignore case": {
			annotations: map[string]string{
				PropagationEnableAnnotationKey: "True",
			},
			expected: true,
		},
		"has propagation disabled": {
			annotations: map[string]string{
				PropagationEnableAnnotationKey: "random-value",
			},
			expected: false,
		},
		"has propagation annotation empty": {
			annotations: map[string]string{},
			expected:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := hasPropagationEnabledAnnotation(test.annotations)
			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_propagationNamespaces(t *testing.T) {
	tests := map[string]struct {
		annotations map[string]string
		expected    int
	}{
		"namespaces not specified": {
			annotations: map[string]string{},
			expected:    0,
		},
		"single namespace target": {
			annotations: map[string]string{
				PropagationNamespaceAnnotationKey: "default",
			},
			expected: 1,
		},
		"multiple namespace target": {
			annotations: map[string]string{
				PropagationNamespaceAnnotationKey: "default,kube-system,kube-public",
			},
			expected: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := len(getPropagationNamespacesFromAnnotations(test.annotations))
			assert.Equal(t, test.expected, result)
		})
	}
}
