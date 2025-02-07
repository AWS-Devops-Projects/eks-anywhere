package kubernetes

import (
	"k8s.io/apimachinery/pkg/runtime"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type schemeAdder func(s *runtime.Scheme) error

var schemeAdders = []schemeAdder{
	clusterv1.AddToScheme,
	controlplanev1.AddToScheme,
	anywherev1.AddToScheme,
}

func addToScheme(scheme *runtime.Scheme, schemeAdder ...schemeAdder) error {
	for _, adder := range schemeAdders {
		if err := adder(scheme); err != nil {
			return err
		}
	}

	return nil
}
