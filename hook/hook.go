package hook

import (
	"errors"

	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
)

const (
	LevelFieldName = "level"
)

type MetricHook struct {
	config  *Config
	metrics *types.Linker[logrus.Level, prometheus.Counter]
}

func NewMetricHook(appConfig *compogo.Config, config *Config) *MetricHook {
	hook := &MetricHook{
		config:  config,
		metrics: types.NewLinker[logrus.Level, prometheus.Counter](),
	}

	counterVec := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: compogo.MetricNamePrefix + "log_level_count",
		Help: "log counter of a certain level",
		ConstLabels: map[string]string{
			compogo.MetricAppNameFieldName: appConfig.Name,
		},
	}, []string{LevelFieldName})

	for level := range config.Levels {
		hook.metrics.Add(level, counterVec.With(prometheus.Labels{LevelFieldName: level.String()}))
	}

	return hook
}

func (metric *MetricHook) Levels() []logrus.Level {
	return metric.config.Levels.ToSlice()
}

func (metric *MetricHook) Fire(entry *logrus.Entry) error {
	counter, err := metric.metrics.Get(entry.Level)
	if err != nil && !errors.Is(err, types.DoesNotExistError) {
		return err
	}

	if counter != nil {
		counter.Inc()
	}

	return nil
}
