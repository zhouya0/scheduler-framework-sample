package qos

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/api/v1/pod"
	v1qos "k8s.io/kubernetes/pkg/apis/core/v1/helper/qos"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const Name = "QoSSort"

type QoSSort struct{}

var _ framework.QueueSortPlugin = &QoSSort{}

// Name returns name of the plugin
func (qs *QoSSort) Name() string {
	return Name
}

func (qs *QoSSort) Less(pInfo1, pInfo2 *framework.PodInfo) bool {
	p1 := pod.GetPodPriority(pInfo1.Pod)
	p2 := pod.GetPodPriority(pInfo2.Pod)
	return (p1 > p2) || (p1 == p2 && compQOS(pInfo1.Pod, pInfo2.Pod))
}

func compQOS(p1, p2 *v1.Pod) bool {
	p1QOS, p2QOS := v1qos.GetPodQOS(p1), v1qos.GetPodQOS(p2)
	if p1QOS == v1.PodQOSGuaranteed {
		return true
	} else if p1QOS == v1.PodQOSBurstable {
		return p2QOS != v1.PodQOSGuaranteed
	} else {
		return p2QOS == v1.PodQOSBestEffort
	}
}

// New initializes a new plugin and returns it
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &QoSSort{}, nil
}