// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for splunkenterprise metrics.
type MetricsConfig struct {
	SplunkAggregationQueueRatio                 MetricConfig `mapstructure:"splunk.aggregation.queue.ratio"`
	SplunkBucketsSearchableStatus               MetricConfig `mapstructure:"splunk.buckets.searchable.status"`
	SplunkDataIndexesExtendedBucketCount        MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.count"`
	SplunkDataIndexesExtendedBucketEventCount   MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.event.count"`
	SplunkDataIndexesExtendedBucketHotCount     MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.hot.count"`
	SplunkDataIndexesExtendedBucketWarmCount    MetricConfig `mapstructure:"splunk.data.indexes.extended.bucket.warm.count"`
	SplunkDataIndexesExtendedEventCount         MetricConfig `mapstructure:"splunk.data.indexes.extended.event.count"`
	SplunkDataIndexesExtendedRawSize            MetricConfig `mapstructure:"splunk.data.indexes.extended.raw.size"`
	SplunkDataIndexesExtendedTotalSize          MetricConfig `mapstructure:"splunk.data.indexes.extended.total.size"`
	SplunkHealth                                MetricConfig `mapstructure:"splunk.health"`
	SplunkIndexerAvgRate                        MetricConfig `mapstructure:"splunk.indexer.avg.rate"`
	SplunkIndexerCPUTime                        MetricConfig `mapstructure:"splunk.indexer.cpu.time"`
	SplunkIndexerQueueRatio                     MetricConfig `mapstructure:"splunk.indexer.queue.ratio"`
	SplunkIndexerRawWriteTime                   MetricConfig `mapstructure:"splunk.indexer.raw.write.time"`
	SplunkIndexerRollingrestartStatus           MetricConfig `mapstructure:"splunk.indexer.rollingrestart.status"`
	SplunkIndexerThroughput                     MetricConfig `mapstructure:"splunk.indexer.throughput"`
	SplunkIndexesAvgSize                        MetricConfig `mapstructure:"splunk.indexes.avg.size"`
	SplunkIndexesAvgUsage                       MetricConfig `mapstructure:"splunk.indexes.avg.usage"`
	SplunkIndexesBucketCount                    MetricConfig `mapstructure:"splunk.indexes.bucket.count"`
	SplunkIndexesMedianDataAge                  MetricConfig `mapstructure:"splunk.indexes.median.data.age"`
	SplunkIndexesSize                           MetricConfig `mapstructure:"splunk.indexes.size"`
	SplunkIoAvgIops                             MetricConfig `mapstructure:"splunk.io.avg.iops"`
	SplunkKvstoreBackupStatus                   MetricConfig `mapstructure:"splunk.kvstore.backup.status"`
	SplunkKvstoreReplicationStatus              MetricConfig `mapstructure:"splunk.kvstore.replication.status"`
	SplunkKvstoreStatus                         MetricConfig `mapstructure:"splunk.kvstore.status"`
	SplunkLicenseIndexUsage                     MetricConfig `mapstructure:"splunk.license.index.usage"`
	SplunkParseQueueRatio                       MetricConfig `mapstructure:"splunk.parse.queue.ratio"`
	SplunkPipelineSetCount                      MetricConfig `mapstructure:"splunk.pipeline.set.count"`
	SplunkSchedulerAvgExecutionLatency          MetricConfig `mapstructure:"splunk.scheduler.avg.execution.latency"`
	SplunkSchedulerAvgRunTime                   MetricConfig `mapstructure:"splunk.scheduler.avg.run.time"`
	SplunkSchedulerCompletionRatio              MetricConfig `mapstructure:"splunk.scheduler.completion.ratio"`
	SplunkServerIntrospectionQueuesCurrent      MetricConfig `mapstructure:"splunk.server.introspection.queues.current"`
	SplunkServerIntrospectionQueuesCurrentBytes MetricConfig `mapstructure:"splunk.server.introspection.queues.current.bytes"`
	SplunkServerSearchartifactsAdhoc            MetricConfig `mapstructure:"splunk.server.searchartifacts.adhoc"`
	SplunkServerSearchartifactsCompleted        MetricConfig `mapstructure:"splunk.server.searchartifacts.completed"`
	SplunkServerSearchartifactsIncomplete       MetricConfig `mapstructure:"splunk.server.searchartifacts.incomplete"`
	SplunkServerSearchartifactsInvalid          MetricConfig `mapstructure:"splunk.server.searchartifacts.invalid"`
	SplunkServerSearchartifactsJobCacheCount    MetricConfig `mapstructure:"splunk.server.searchartifacts.job.cache.count"`
	SplunkServerSearchartifactsJobCacheSize     MetricConfig `mapstructure:"splunk.server.searchartifacts.job.cache.size"`
	SplunkServerSearchartifactsSavedsearches    MetricConfig `mapstructure:"splunk.server.searchartifacts.savedsearches"`
	SplunkServerSearchartifactsScheduled        MetricConfig `mapstructure:"splunk.server.searchartifacts.scheduled"`
	SplunkTypingQueueRatio                      MetricConfig `mapstructure:"splunk.typing.queue.ratio"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SplunkAggregationQueueRatio: MetricConfig{
			Enabled: false,
		},
		SplunkBucketsSearchableStatus: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedBucketCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedBucketEventCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedBucketHotCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedBucketWarmCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedEventCount: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedRawSize: MetricConfig{
			Enabled: false,
		},
		SplunkDataIndexesExtendedTotalSize: MetricConfig{
			Enabled: false,
		},
		SplunkHealth: MetricConfig{
			Enabled: true,
		},
		SplunkIndexerAvgRate: MetricConfig{
			Enabled: false,
		},
		SplunkIndexerCPUTime: MetricConfig{
			Enabled: false,
		},
		SplunkIndexerQueueRatio: MetricConfig{
			Enabled: false,
		},
		SplunkIndexerRawWriteTime: MetricConfig{
			Enabled: false,
		},
		SplunkIndexerRollingrestartStatus: MetricConfig{
			Enabled: false,
		},
		SplunkIndexerThroughput: MetricConfig{
			Enabled: false,
		},
		SplunkIndexesAvgSize: MetricConfig{
			Enabled: false,
		},
		SplunkIndexesAvgUsage: MetricConfig{
			Enabled: false,
		},
		SplunkIndexesBucketCount: MetricConfig{
			Enabled: false,
		},
		SplunkIndexesMedianDataAge: MetricConfig{
			Enabled: false,
		},
		SplunkIndexesSize: MetricConfig{
			Enabled: false,
		},
		SplunkIoAvgIops: MetricConfig{
			Enabled: false,
		},
		SplunkKvstoreBackupStatus: MetricConfig{
			Enabled: false,
		},
		SplunkKvstoreReplicationStatus: MetricConfig{
			Enabled: false,
		},
		SplunkKvstoreStatus: MetricConfig{
			Enabled: false,
		},
		SplunkLicenseIndexUsage: MetricConfig{
			Enabled: false,
		},
		SplunkParseQueueRatio: MetricConfig{
			Enabled: false,
		},
		SplunkPipelineSetCount: MetricConfig{
			Enabled: false,
		},
		SplunkSchedulerAvgExecutionLatency: MetricConfig{
			Enabled: false,
		},
		SplunkSchedulerAvgRunTime: MetricConfig{
			Enabled: false,
		},
		SplunkSchedulerCompletionRatio: MetricConfig{
			Enabled: false,
		},
		SplunkServerIntrospectionQueuesCurrent: MetricConfig{
			Enabled: false,
		},
		SplunkServerIntrospectionQueuesCurrentBytes: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsAdhoc: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsCompleted: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsIncomplete: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsInvalid: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsJobCacheCount: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsJobCacheSize: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsSavedsearches: MetricConfig{
			Enabled: false,
		},
		SplunkServerSearchartifactsScheduled: MetricConfig{
			Enabled: false,
		},
		SplunkTypingQueueRatio: MetricConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for splunkenterprise metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
