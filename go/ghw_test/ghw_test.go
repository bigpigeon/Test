/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package ghw_test

import (
	"github.com/jaypipes/ghw"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGHW(t *testing.T) {
	cpu, err := ghw.CPU(ghw.WithChroot("/host"))
	require.NoError(t, err)
	t.Log(cpu.String())
	netDev, err := ghw.Network()
	require.NoError(t, err)
	t.Log(netDev.String())
	for _, nic := range netDev.NICs {
		t.Logf(" %v\n", nic.String())

		enabledCaps := make([]int, 0)
		for x, cap := range nic.Capabilities {
			if cap.IsEnabled {
				enabledCaps = append(enabledCaps, x)
			}
		}

	}
	block, err := ghw.Block()
	require.NoError(t, err)

	t.Logf("%v\n", block.YAMLString())

	for _, disk := range block.Disks {
		t.Logf(" %v\n", disk)
		for _, part := range disk.Partitions {
			t.Logf("  %v\n", part)
		}
	}
	baseboard, err := ghw.Baseboard()
	require.NoError(t, err)

	t.Log(baseboard.String())

	chassis, err := ghw.Chassis()
	require.NoError(t, err)

	t.Logf("%v\n", chassis)
	product, err := ghw.Product()
	require.NoError(t, err)

	t.Logf("%v\n", product)
}
