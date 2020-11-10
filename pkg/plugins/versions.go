package plugins

import jenkinsv1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"

const (
	// HelmPluginName the default name of the helm plugin
	HelmPluginName = "helm"

	// HelmfilePluginName the default name of the helmfile plugin
	HelmfilePluginName = "helmfile"

	// KptPluginName the default name of the kpt plugin
	KptPluginName = "kpt"

	// KubectlPluginName the default name of the kubectl plugin
	KubectlPluginName = "kubectl"

	// HelmVersion the default version of helm to use
	HelmVersion = "3.4.0"

	// HelmfileVersion the default version of helmfile to use
	HelmfileVersion = "0.134.0"

	// KptVersion the default version of kpt to use
	KptVersion = "0.33.0"

	// KubectlVersion the default version of kpt to use
	KubectlVersion = "1.16.15"
)

var (
	// Plugins default plugins
	Plugins = []jenkinsv1.Plugin{
		CreateHelmPlugin(HelmVersion),
		CreateHelmfilePlugin(HelmfileVersion),
		CreateKptPlugin(KptVersion),
		CreateKubectlPlugin(KubectlVersion),
	}
)
