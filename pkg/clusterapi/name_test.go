package clusterapi_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/aws/eks-anywhere/pkg/clusterapi"
)

func TestIncrementName(t *testing.T) {
	tests := []struct {
		name    string
		oldName string
		want    string
		wantErr string
	}{
		{
			name:    "valid",
			oldName: "cluster-1",
			want:    "cluster-2",
			wantErr: "",
		},
		{
			name:    "invalid format",
			oldName: "cluster-1a",
			want:    "",
			wantErr: "invalid format of name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			got, err := clusterapi.IncrementName(tt.oldName)
			if tt.wantErr == "" {
				g.Expect(err).To(Succeed())
				g.Expect(got).To(Equal(tt.want))
			} else {
				g.Expect(err).To(MatchError(ContainSubstring(tt.wantErr)))
			}
		})
	}
}

func TestObjectName(t *testing.T) {
	tests := []struct {
		name    string
		base    string
		version int
		want    string
	}{
		{
			name:    "cluster-1",
			base:    "cluster",
			version: 1,
			want:    "cluster-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			g.Expect(clusterapi.ObjectName(tt.base, tt.version)).To(Equal(tt.want))
		})
	}
}

func TestDefaultObjectName(t *testing.T) {
	tests := []struct {
		name string
		base string
		want string
	}{
		{
			name: "cluster-1",
			base: "cluster",
			want: "cluster-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			g.Expect(clusterapi.DefaultObjectName(tt.base)).To(Equal(tt.want))
		})
	}
}

func TestKubeadmControlPlaneName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test cluster",
			want: "test-cluster",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newApiBuilerTest(t)
			g.Expect(clusterapi.KubeadmControlPlaneName(g.clusterSpec)).To(Equal(tt.want))
		})
	}
}

func TestMachineDeploymentName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "wng 1",
			want: "wng-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newApiBuilerTest(t)
			g.Expect(clusterapi.MachineDeploymentName(*g.workerNodeGroupConfig)).To(Equal(tt.want))
		})
	}
}
