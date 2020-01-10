/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package prometheus_test

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestHistogram(t *testing.T) {

	var histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "bigpigeon",
		Subsystem: "test",
		Name:      "histogram_test_bucket",
		Help:      "test histogram bucket",
		Buckets:   prometheus.ExponentialBuckets(0.001, 2, 14),
	})
	err := prometheus.Register(histogram)
	require.NoError(t, err)
	histogram.Observe(0.0005)
	histogram.Observe(0.0006)
	histogram.Observe(0.02)
	histogram.Observe(0.03)
	histogram.Observe(0.4)
	histogram.Observe(0.5)
	histogram.Observe(1.2)
	histogram.Observe(1.3)
	histogram.Observe(1.4)
	histogram.Observe(2.2)

	mfs, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)
	for _, mf := range mfs {
		if mf.Help != nil && *mf.Help == "test histogram bucket" {
			if _, err := expfmt.MetricFamilyToText(os.Stdout, mf); err != nil {
				t.Error(err)
			}
		}
	}
	histogram.Observe(0.00051)
	mfs, err = prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)
	for _, mf := range mfs {
		if mf.Help != nil && *mf.Help == "test histogram bucket" {
			if _, err := expfmt.MetricFamilyToText(os.Stdout, mf); err != nil {
				t.Error(err)
			}
		}
	}

}
