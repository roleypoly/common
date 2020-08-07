package version

import (
	"fmt"
	"io"
)

var (
	GitCommit = "unknown"
	GitBranch = "unknown"
	BuildDate = "unknown"
)

func WriteStartup(serviceName string, writer io.Writer) {
	writer.Write(
		[]byte(fmt.Sprintf(
			"Starting %s service.\n Build %s (%s) at %s",
			serviceName,
			GitCommit,
			GitBranch,
			BuildDate,
		)),
	)
}
