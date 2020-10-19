/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package mongo

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/mgo.v2"
	"testing"
	"time"
)

type SliceTagID struct {
	FileID    string `bson:"_id"         json:"fileID,omitempty"`
	Namespace string `bson:"ownerName"   json:"namespace,omitempty"`
	Tag       string `bson:"tag"         json:"tag"`
}

type SliceTagInfo struct {
	SliceTagID `bson:"_id"   `
	Slice      []time.Duration `bson:"slice" json:"slice"`
}

func TestUnionID(t *testing.T) {
	session, err := mgo.DialWithTimeout("localhost:27017", time.Second)
	require.NoError(t, err)
	ctx := session.DB("").C("unionkey")
	doc := &SliceTagInfo{
		SliceTagID: SliceTagID{
			FileID:    uuid.New().String(),
			Namespace: "test",
			Tag:       "diar",
		},
		Slice: []time.Duration{100, 200, 300, 400},
	}
	err = ctx.Insert(doc)
	require.NoError(t, err)
	var findDoc SliceTagInfo
	err = ctx.FindId(doc.SliceTagID).One(&findDoc)
	require.NoError(t, err)
	require.Equal(t, *doc, findDoc)
}
