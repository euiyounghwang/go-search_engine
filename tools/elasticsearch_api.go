package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	// elasticsearch "github.com/elastic/go-elasticsearch/v7"

	"github.com/dustin/go-humanize"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

var (
    IndexNameEmptyStringError = errors.New("index name cannot be empty string")
    IndexAlreadyExistsError   = errors.New("elasticsearch index already exists")
)


func Set_Env(initial_str string, replace_str string) (string) {
	transform_str := ""
	if initial_str == "" {
		transform_str = replace_str
	}
	log.Println("Set_Env : ", transform_str)
	return replace_str
}

type ElasticDocs struct {
	SomeStr string
	SomeInt int
	SomeBool bool
	Timestamp int64
}

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Published time.Time `json:"published"`
	Author    Author    `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	indexName  string
	numWorkers int
	flushBytes int
	numItems   int
)

func init() {
	// https://github.com/elastic/go-elasticsearch/blob/main/_examples/bulk/indexer.go
	// You can configure the settings with command line flags:
	//
	//     go run indexer.go --workers=8 --count=100000 --flush=1000000
	//
	log.Println("init..")
	// flag.StringVar(&indexName, "index", "test-bulk-example", "Index name")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	flag.IntVar(&numItems, "count", 10000, "Number of documents to generate")
	flag.Parse()
	
	rand.Seed(time.Now().UnixNano())
}

/*
go run ./tools/elasticsearch_api.go
*/

func main() {
	
	host := Set_Env(os.Getenv("ES_HOST"), "http://localhost:9209")
    index := Set_Env(os.Getenv("ES_INDEX"), "go_test_omnisearch_v1")
	
	if index == "" {
        log.Fatal(IndexNameEmptyStringError)
    }

	es_client, err := elasticsearch.NewClient(
		elasticsearch.Config{
        	Addresses: []string{host},
			Username: "elastic",
			Password: "gsaadmin",
    	},
		// elastic.SetURL(host),
	)
	
	res, err := es_client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
		
	log.Println("Elasticsearch Initializing..")
	
	try_delete_index := func() {
		log.Println("inner func")
		response, err := es_client.Indices.Exists([]string{index})
		log.Println(response)
		if response.StatusCode != 404 {
			// log.Fatal(IndexAlreadyExistsError)
			es_client.Indices.Delete([]string{index})
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	
	try_delete_index()
	
    response, err := es_client.Indices.Create(index)
    if err != nil {
        log.Fatal(err)
    }

    // if response.IsError() {
    //     log.Fatal(err)
    // }
	
	if response.StatusCode == 200 {
		log.Println("Elasticsearch Indices Created..")
	}
	
	// Declare an empty slice for the Elasticsearch document struct objects
	docs := []ElasticDocs{}
	// Get the type of the 'docs' struct slice
	log.Println("docs TYPE:", reflect.TypeOf(docs))
	
	// New ElasticDocs struct instances
	newDoc1 := ElasticDocs{SomeStr: "Hello, world!", SomeInt: 42, SomeBool: true, Timestamp: 0.0}
	newDoc2 := ElasticDocs{SomeStr: "Hello, world2!", SomeInt: 7654, SomeBool: false, Timestamp: 0.0}
	
	// Append the new Elasticsearch document struct objects to the slice
	docs = append(docs, newDoc1)
	docs = append(docs, newDoc2)
	
	log.Println(docs)
	
	// res, _ = es_client.Index(index, esutil.NewJSONReader(&newDoc1))
	// fmt.Println(res)
	
	res, _ = es_client.Index(index, esutil.NewJSONReader(newDoc1))
	fmt.Println(res)
	defer res.Body.Close()
	
	// Deserialize the response into a map.
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	} else {
		log.Printf("\nIndexRequest() RESPONSE:")
		// Print the response status and indexed document version.
		fmt.Println("Status:", res.Status())
		fmt.Println("Result:", resMap["_index"])
		// fmt.Println("Version:", int(resMap["_version"].(float64)))
		fmt.Println("resMap:", resMap)
		fmt.Println("\n")
	}
	
	// Generate the articles collection
	
	var (
		articles        []*Article
		countSuccessful uint64

		// res *esapi.Response
		// err error
	)

	names := []string{"Alice", "John", "Mary"}
	for i := 1; i <= numItems; i++ {
		articles = append(articles, &Article{
			ID:        i,
			Title:     strings.Join([]string{"Title", strconv.Itoa(i)}, " "),
			Body:      "Lorem ipsum dolor sit amet...",
			Published: time.Now().Round(time.Second).UTC().AddDate(0, 0, i),
			Author: Author{
				FirstName: names[rand.Intn(len(names))],
				LastName:  "Smith",
			},
		})
	}
	log.Printf("→ Generated %s articles", humanize.Comma(int64(len(articles))))
	
	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//
	// Create the BulkIndexer
	//
	// NOTE: For optimal performance, consider using a third-party JSON decoding package.
	//       See an example in the "benchmarks" folder.
	//
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         index,        // The default index name
		Client:        es_client,               // The Elasticsearch client
		NumWorkers:    numWorkers,       // The number of worker goroutines
		FlushBytes:    int(flushBytes),  // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	
	start := time.Now().UTC()
	
	// Loop over the collection
	//
	for _, a := range articles {
		// Prepare the data payload: encode article to JSON
		//
		data, err := json.Marshal(a)
		// log.Println((data))
		if err != nil {
			log.Fatalf("Cannot encode article %d: %s", a.ID, err)
		}

		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		//
		// Add an item to the BulkIndexer
		//
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",

				// DocumentID is the (optional) document ID
				DocumentID: strconv.Itoa(a.ID),

				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),

				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
		if err := bi.Close(context.Background()); err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	
		biStats := bi.Stats()
	
		// Report the results: number of indexed docs, number of errors, duration, indexing rate
		//
		log.Println(strings.Repeat("▔", 65))
	
		dur := time.Since(start)
	
		if biStats.NumFailed > 0 {
			log.Fatalf(
				"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
				humanize.Comma(int64(biStats.NumFlushed)),
				humanize.Comma(int64(biStats.NumFailed)),
				dur.Truncate(time.Millisecond),
				humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
			)
		} else {
			log.Printf(
				"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
				humanize.Comma(int64(biStats.NumFlushed)),
				dur.Truncate(time.Millisecond),
				humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
			)
		}
	}

}
