/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package k8s

import (
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"testing"
	"time"
)

const resyncPeriod = 10 * time.Minute

func TestGetSecret(t *testing.T) {
	{
		fac := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("kube-system"))
		stopper := make(chan struct{})

		fac.Start(stopper)

		time.Sleep(1 * time.Second)
		t.Log(fac.Core().V1().Secrets().Informer().LastSyncResourceVersion())
		{

			ret, err := fac.Core().V1().Secrets().Lister().Secrets("kube-system").List(labels.Everything())
			require.NoError(t, err)
			for _, v := range ret {
				t.Log(v.Name)
			}
		}
		{

			svcs, err := fac.Core().V1().Services().Lister().List(labels.Everything())
			require.NoError(t, err)
			for _, v := range svcs {
				t.Log(v.String())
			}
		}
	}
	{
		secretList, err := clientset.CoreV1().Secrets("kube-system").List(metav1.ListOptions{})
		require.NoError(t, err)
		for _, v := range secretList.Items {
			t.Log(v.Name)
		}
	}
}
