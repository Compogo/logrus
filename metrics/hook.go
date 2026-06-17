package metrics

import (
	"errors"

	"github.com/Compogo/compogo"
	typesErrors "github.com/Compogo/types/errors"
	"github.com/Compogo/types/linker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
)

// Константы для метрик.
const (
	// LevelFieldName — имя поля в метрике для уровня логирования.
	LevelFieldName = "level"
)

// MetricHook реализует хук Logrus для сбора метрик Prometheus.
// Считает количество логов каждого уровня.
//
// Метрика: compogo_log_level_count{app="myapp", level="error"}
type MetricHook struct {
	config  *Config
	metrics *linker.Linker[logrus.Level, prometheus.Counter]
}

// NewMetricHook создаёт новый хук для метрик.
// Регистрирует счётчик для каждого уровня из конфигурации.
//
// Пример:
//
//	hook := NewMetricHook(appConfig, config)
//	logger.AddHook(hook)
func NewMetricHook(appConfig *compogo.Config, config *Config) *MetricHook {
	hook := &MetricHook{
		config:  config,
		metrics: linker.NewLinker[logrus.Level, prometheus.Counter](),
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

// Levels возвращает уровни, на которых активируется хук.
func (metric *MetricHook) Levels() []logrus.Level {
	return metric.config.Levels.ToSlice()
}

// Fire вызывается при каждом логировании.
// Увеличивает счётчик для соответствующего уровня.
func (metric *MetricHook) Fire(entry *logrus.Entry) error {
	counter, err := metric.metrics.Get(entry.Level)
	if err != nil && !errors.Is(err, typesErrors.DoesNotExistError) {
		return err
	}

	if counter != nil {
		counter.Inc()
	}

	return nil
}
