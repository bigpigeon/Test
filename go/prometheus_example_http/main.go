/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"github.com/prometheus/common/expfmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func recordMetrics(opsProcessed prometheus.Counter) {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	reg := prometheus.NewRegistry()
	opsProcessed := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "bigpigeon",
		Subsystem: "test",
		Name:      "processed_count",
		Help:      "The total number of processed events",
		ConstLabels: prometheus.Labels{
			"type": "count",
		},
	})
	recordMetrics(opsProcessed)
	reg.MustRegister(opsProcessed)

	opsPreProcess := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "bigpigeon",
		Subsystem: "test",
		Name:      "preprocess_count",
		Help:      "The total number of processed events",
		ConstLabels: prometheus.Labels{
			"type":   "count",
			"action": "pre",
		},
	})
	recordMetrics(opsPreProcess)
	reg.MustRegister(opsPreProcess)

	opsProcessedSum := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "bigpigeon",
		Subsystem: "test",
		Name:      "processed_seconds_sum",
		Help:      "The total seconds of processed events",
		ConstLabels: prometheus.Labels{
			"type": "sum",
		},
	})
	recordMetrics(opsProcessedSum)
	reg.MustRegister(opsProcessedSum)

	http.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {
		mfs, err := reg.Gather()
		if err != nil {
			panic(err)
		}
		for _, mf := range mfs {
			if _, err := expfmt.MetricFamilyToText(writer, mf); err != nil {
				if err != nil {
					panic(err)
				}
			}
		}
	})
	http.ListenAndServe(":19090", nil)
}
