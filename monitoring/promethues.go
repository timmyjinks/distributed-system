package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusService struct {
	reg          *prometheus.Registry
	opsProcessed prometheus.Counter
}

func NewPrometheusService(metricName string, help string) *PrometheusService {
	reg := prometheus.NewRegistry()
	m := &PrometheusService{
		reg: reg,
		opsProcessed: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: metricName,
			Help: "The total number of processed events",
		}),
	}
	return m
}

func (p *PrometheusService) Start(mux *http.ServeMux) {
	mux.Handle("/metrics", promhttp.HandlerFor(p.reg, promhttp.HandlerOpts{}))
}

func (p *PrometheusService) Inc() {
	p.opsProcessed.Inc()
}
