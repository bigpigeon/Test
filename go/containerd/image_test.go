/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"context"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListImage(t *testing.T) {
	ctx := context.Background()
	client, err := containerd.New("/run/containerd/containerd.sock")
	require.NoError(t, err)
	defer client.Close()
	nss, err := client.NamespaceService().List(ctx)
	for _, ns := range nss {
		t.Log("namespace:", ns)
	}
	ctx = namespaces.WithNamespace(ctx, "example")
	imgs, err := client.ListImages(ctx)
	require.NoError(t, err)
	for _, img := range imgs {
		t.Log("image name:", img.Name())
	}
}

func TestContainer(t *testing.T) {
	ctx := context.Background()
	client, err := containerd.New("/run/containerd/containerd.sock")
	require.NoError(t, err)
	ctx = namespaces.WithNamespace(ctx, "example")
	// pull an image and unpack it into the configured snapshotter
	image, err := client.Pull(ctx, "docker.io/library/redis:latest", containerd.WithPullUnpack)

	// allocate a new RW root filesystem for a container based on the image
	redis, err := client.NewContainer(ctx, "redis-master",
		containerd.WithNewSnapshot("redis-rootfs", image),
		containerd.WithNewSpec(oci.WithImageConfig(image)),
	)
	require.NoError(t, err)
	defer redis.Delete(ctx)
	task, err := redis.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	require.NoError(t, err)
	pid := task.Pid()
	t.Log("pid", pid)
	defer task.Delete(ctx)
}
