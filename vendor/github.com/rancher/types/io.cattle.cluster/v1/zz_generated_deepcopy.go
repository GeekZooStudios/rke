package v1

import (
	core_v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSConfig) DeepCopyInto(out *AKSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSConfig.
func (in *AKSConfig) DeepCopy() *AKSConfig {
	if in == nil {
		return nil
	}
	out := new(AKSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComponentStatus) DeepCopyInto(out *ClusterComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]core_v1.ComponentCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComponentStatus.
func (in *ClusterComponentStatus) DeepCopy() *ClusterComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNode) DeepCopyInto(out *ClusterNode) {
	*out = *in
	in.Node.DeepCopyInto(&out.Node)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNode.
func (in *ClusterNode) DeepCopy() *ClusterNode {
	if in == nil {
		return nil
	}
	out := new(ClusterNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNodeList) DeepCopyInto(out *ClusterNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNodeList.
func (in *ClusterNodeList) DeepCopy() *ClusterNodeList {
	if in == nil {
		return nil
	}
	out := new(ClusterNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.GKEConfig != nil {
		in, out := &in.GKEConfig, &out.GKEConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(GKEConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AKSConfig != nil {
		in, out := &in.AKSConfig, &out.AKSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(AKSConfig)
			**out = **in
		}
	}
	if in.RKEConfig != nil {
		in, out := &in.RKEConfig, &out.RKEConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(RKEConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		copy(*out, *in)
	}
	if in.ComponentStatuses != nil {
		in, out := &in.ComponentStatuses, &out.ComponentStatuses
		*out = make([]ClusterComponentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(core_v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Allocatable != nil {
		in, out := &in.Allocatable, &out.Allocatable
		*out = make(core_v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDService) DeepCopyInto(out *ETCDService) {
	*out = *in
	out.baseService = in.baseService
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDService.
func (in *ETCDService) DeepCopy() *ETCDService {
	if in == nil {
		return nil
	}
	out := new(ETCDService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKEConfig) DeepCopyInto(out *GKEConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.UpdateConfig = in.UpdateConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKEConfig.
func (in *GKEConfig) DeepCopy() *GKEConfig {
	if in == nil {
		return nil
	}
	out := new(GKEConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAPIService) DeepCopyInto(out *KubeAPIService) {
	*out = *in
	out.baseService = in.baseService
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAPIService.
func (in *KubeAPIService) DeepCopy() *KubeAPIService {
	if in == nil {
		return nil
	}
	out := new(KubeAPIService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerService) DeepCopyInto(out *KubeControllerService) {
	*out = *in
	out.baseService = in.baseService
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerService.
func (in *KubeControllerService) DeepCopy() *KubeControllerService {
	if in == nil {
		return nil
	}
	out := new(KubeControllerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletService) DeepCopyInto(out *KubeletService) {
	*out = *in
	out.baseService = in.baseService
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletService.
func (in *KubeletService) DeepCopy() *KubeletService {
	if in == nil {
		return nil
	}
	out := new(KubeletService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeproxyService) DeepCopyInto(out *KubeproxyService) {
	*out = *in
	out.baseService = in.baseService
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeproxyService.
func (in *KubeproxyService) DeepCopy() *KubeproxyService {
	if in == nil {
		return nil
	}
	out := new(KubeproxyService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEConfig) DeepCopyInto(out *RKEConfig) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]RKEConfigHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Services = in.Services
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEConfig.
func (in *RKEConfig) DeepCopy() *RKEConfig {
	if in == nil {
		return nil
	}
	out := new(RKEConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEConfigHost) DeepCopyInto(out *RKEConfigHost) {
	*out = *in
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEConfigHost.
func (in *RKEConfigHost) DeepCopy() *RKEConfigHost {
	if in == nil {
		return nil
	}
	out := new(RKEConfigHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEConfigServices) DeepCopyInto(out *RKEConfigServices) {
	*out = *in
	out.Etcd = in.Etcd
	out.KubeAPI = in.KubeAPI
	out.KubeController = in.KubeController
	out.Scheduler = in.Scheduler
	out.Kubelet = in.Kubelet
	out.Kubeproxy = in.Kubeproxy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEConfigServices.
func (in *RKEConfigServices) DeepCopy() *RKEConfigServices {
	if in == nil {
		return nil
	}
	out := new(RKEConfigServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerService) DeepCopyInto(out *SchedulerService) {
	*out = *in
	out.baseService = in.baseService
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerService.
func (in *SchedulerService) DeepCopy() *SchedulerService {
	if in == nil {
		return nil
	}
	out := new(SchedulerService)
	in.DeepCopyInto(out)
	return out
}
