package setup

import (
	"github.com/mongodb/mongodb-kubernetes-operator/pkg/util/envvar"
)

const (
	testNamespaceEnvName      = "TEST_NAMESPACE"
	operatorImageEnvName      = "OPERATOR_IMAGE"
	agentImage                = "AGENT_IMAGE"
	clusterWideEnvName        = "CLUSTER_WIDE"
	versionUpgradeHookEnvName = "VERSION_UPGRADE_HOOK_IMAGE"
	performCleanupEnvName     = "PERFORM_CLEANUP"
)

type testConfig struct {
	namespace               string
	operatorImage           string
	versionUpgradeHookImage string
	clusterWide             bool
	performCleanup          bool
	agentImage              string
}

func loadTestConfigFromEnv() testConfig {
	return testConfig{
		namespace:               envvar.GetEnvOrDefault(testNamespaceEnvName, "default"),
		operatorImage:           envvar.GetEnvOrDefault(operatorImageEnvName, "quay.io/mongodb/community-operator-dev:latest"),
		versionUpgradeHookImage: envvar.GetEnvOrDefault(versionUpgradeHookEnvName, "quay.io/mongodb/mongodb-kubernetes-operator-version-upgrade-post-start-hook:1.0.2"),
		agentImage:              envvar.GetEnvOrDefault(agentImage, "quay.io/mongodb/mongodb-agent:10.27.0.6772-1"), // TODO: better way to decide default agent image.
		clusterWide:             envvar.ReadBool(clusterWideEnvName),
		performCleanup:          envvar.ReadBool(performCleanupEnvName),
	}
}
