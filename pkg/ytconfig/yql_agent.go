package ytconfig

import (
	ytv1 "github.com/ytsaurus/yt-k8s-operator/api/v1"
	"github.com/ytsaurus/yt-k8s-operator/pkg/consts"
)

type YQLEmbedded struct {
	MRJobBinary  string `yson:"mr_job_binary"`
	UDFDirectory string `yson:"udf_directory"`
}

type YQLAgent struct {
	YQLEmbedded
	AdditionalClusters map[string]string `yson:"additional_clusters"`
}

type YQLAgentServer struct {
	CommonServer
	User     string   `yson:"user"`
	YQLAgent YQLAgent `yson:"yql_agent"`
}

func getYQLAgentServerCarcass(spec ytv1.YQLAgentSpec) (YQLAgentServer, error) {
	var c YQLAgentServer
	c.RPCPort = consts.YQLAgentRPCPort
	c.MonitoringPort = consts.YQLAgentMonitoringPort

	c.User = "yql_agent"

	c.YQLAgent.UDFDirectory = "/usr/lib/yql"
	c.YQLAgent.MRJobBinary = "/usr/bin/mrjob"

	c.Logging = createLogging(&spec.InstanceSpec, "yql-agent", []ytv1.LoggerSpec{defaultInfoLoggerSpec(), defaultDebugLoggerSpec(), defaultStderrLoggerSpec()})

	return c, nil
}
