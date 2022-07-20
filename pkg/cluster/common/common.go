package common

const (
	ApplicationLabelKey   = "cloudnative.music.netease.com/application"
	ApplicationIDLabelKey = "cloudnative.music.netease.com/application-id"
	ClusterLabelKey       = "cloudnative.music.netease.com/cluster"
	ClusterIDLabelKey     = "cloudnative.music.netease.com/cluster-id"
	EnvironmentLabelKey   = "cloudnative.music.netease.com/environment"
	PipelinerunIDLabelKey = "cloudnative.music.netease.com/pipelinerun-id"
	RegionLabelKey        = "cloudnative.music.netease.com/region"
	RegionIDLabelKey      = "cloudnative.music.netease.com/region-id"
	OperatorAnnotationKey = "cloudnative.music.netease.com/operator"
	TemplateKey           = "cloudnative.music.netease.com/template"
	RestartTimeKey        = "cloudnative.music.netease.com/user-restart-time"
)

// status of cluster
const (
	StatusEmpty    = ""
	StatusFreeing  = "Freeing"
	StatusFreed    = "Freed"
	StatusDeleting = "Deleting"
	StatusCreating = "Creating"
)
