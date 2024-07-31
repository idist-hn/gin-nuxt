package jobsProvider

import (
	"github.com/bamzi/jobrunner"
	"idist-core/app/jobs"
	"idist-core/app/providers/loggerProvider"
	"time"
)

func Init() {
	loggerProvider.GetLogger().Info("------------------------------------------------------------")

	jobrunner.Start()
	jobrunner.Every(time.Minute*60, &jobs.CleanRequestLog{})
}
