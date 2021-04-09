package pod

import (
	"strconv"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

type mongodbAgentVersion struct {
	Version           string `json:"agent.mongodb.com/version"`
	AppMetrics        string `json:"prometheus.io/app-metrics"`
	AppMetricsPath    string `json:"prometheus.io/app-metrics-path"`
	AppMetricsPort    string `json:"prometheus.io/app-metrics-port"`
	AppMetricsProject string `json:"prometheus.io/app-metrics-project"`
	Scrape            string `json:"prometheus.io/scrape"`
}

func PatchPodAnnotation(podNamespace string, lastVersionAchieved int64, memberName string, clientSet kubernetes.Interface) error {
	patcher := NewKubernetesPodPatcher(clientSet)
	mdbAgentVersion := mongodbAgentVersion{
		Version:           strconv.FormatInt(lastVersionAchieved, 10),
		AppMetrics:        "true",
		AppMetricsPath:    "/metrics",
		AppMetricsPort:    "9216",
		AppMetricsProject: "system",
		Scrape:            "true",
	}
	return patchPod(patcher, podNamespace, mdbAgentVersion, memberName)
}

func patchPod(patcher Patcher, podNamespace string, mdbAgentVersion mongodbAgentVersion, memberName string) error {
	payload := []patchValue{{
		Op:    "add",
		Path:  "/metadata/annotations",
		Value: mdbAgentVersion,
	}}

	pod, err := patcher.patchPod(podNamespace, memberName, payload)
	if pod != nil {
		zap.S().Infof("Updated Pod annotation: %v (%s)", pod.Annotations, memberName)
	}
	return err
}
