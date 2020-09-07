/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

func TestEsCreate(t *testing.T) {
	client, err := elastic.NewClient()
	require.NoError(t, err)
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(context.Background())
	require.NoError(t, err)
	t.Logf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	require.NoError(t, err)
	t.Logf("Elasticsearch version %s\n", esversion)

	exists, err := client.IndexExists("twitter").Do(context.Background())
	require.NoError(t, err)
	if !exists {
		// Create a new index.
		mapping := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"_source": {
			"enabled": false
		},
		"properties":{
			"user":{
				"type":"keyword"
			},
			"message":{
				"type":"text",
				"store": true,
				"fielddata": true
			},
			"retweets":{
				"type":"long"
			},
			"tags":{
				"type":"keyword"
			},
			"location":{
				"type":"geo_point"
			},
			"suggest_field":{
				"type":"completion"
			}
		}
		
	}
}
`
		createIndex, err := client.CreateIndex("twitter").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// Index a tweet (using JSON serialization)
	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	put1, err := client.Index().
		Index("twitter").
		Id("1").
		BodyJson(tweet1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	t.Logf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Index a second tweet (by string)
	tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	put2, err := client.Index().
		Index("twitter").
		Id("2").
		BodyString(tweet2).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	t.Logf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	client.Refresh("twitter").Index("twitter").Do(context.Background())
	// Get tweet with specified ID
	get1, err := client.Get().
		Index("twitter").
		Id("1").
		Do(context.Background())
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			panic(fmt.Sprintf("Document not found: %v", err))
		case elastic.IsTimeout(err):
			panic(fmt.Sprintf("Timeout retrieving document: %v", err))
		case elastic.IsConnErr(err):
			panic(fmt.Sprintf("Connection problem: %v", err))
		default:
			// Some other kind of error
			panic(err)
		}
	}
	t.Logf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	t.Logf("Got source %s", get1.Source)

	// Search with a term query
	//termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("twitter"). // search in index "twitter"
		//Query(termQuery). // specify the query
		//Sort("user", true).      // sort by "user" field, ascending
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(Tweet)
		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

}

func TestQuery(t *testing.T) {
	esAddrs := []string{
		"http://192.168.0.61:9200",
		"http://192.168.0.62:9200",
		"http://192.168.0.64:9200",
	}
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
	// Obtain a client. You can also provide your own HTTP client here.\
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := elastic.DialContext(timeout,
		elastic.SetErrorLog(errorlog),
		elastic.SetSniff(false),
		elastic.SetURL(esAddrs...),
	)
	// Trace request and response details like this
	// client, err := elastic.NewClient(elastic.SetTraceLog(log.New(os.Stdout, "", 0)))
	if err != nil {
		// Handle error
		panic(err)
	}
	esversion, err := client.ElasticsearchVersion(esAddrs[0])
	if err != nil {
		// Handle error
		panic(err)
	}
	t.Log("es version ", esversion)

	searchResult, err := client.Search().
		Index("vps"). // search in index "twitter"
		//Query(termQuery).        // specify the query
		//Sort("user", true).      // sort by "user" field, ascending
		From(0).Size(2).         // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	for _, hit := range searchResult.Hits.Hits {
		fmt.Printf("vsp data by %s: %v\n", hit.Id, hit.Fields)
		result, err := client.Get().Index("vps").
			Id(hit.Id).StoredFields("vec").Do(timeout)
		require.NoError(t, err)
		//result.Source
		fmt.Println("get result", result.Fields)
	}
}
