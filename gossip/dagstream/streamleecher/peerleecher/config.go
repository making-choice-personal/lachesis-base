package peerleecher

import (
	"time"

	"github.com/making-choice-personal/lachesis-base/inter/dag"
)

type EpochDownloaderConfig struct {
	RecheckInterval        time.Duration
	DefaultChunkSize       dag.Metric
	ParallelChunksDownload int
}
