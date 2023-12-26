package main

import (
	"flag"
	"log"

	"example.com/util_elasticsearch"
)

var (
	es_host  string
	index_name string
)

/*
-------
same worker for Python like the following
parser = argparse.ArgumentParser(description="Index into Elasticsearch using this script")
parser.add_argument('-e', '--es', dest='es', default="http://localhost:9250", help='host target')
args = parser.parse_args()
go run ./tools/bulk_index_script.go --es_host=http://localhost:9200 --index_name=test
-------
*/
func init() {
	
	// wordPtr := flag.String("word", "foo", "a string")
	// flag.IntVar(es_host, "es_host", "http://localhost:9209", "Host target")
	
	flag.StringVar(&index_name, "index_name", "test_omnisearch_v1", "a string")
	flag.StringVar(&es_host, "es_host", "http://localhost:9209", "Host target")
	
	flag.Parse()
}

/*
func get_es_instance(es_host string) (*elasticsearch.Client) {
	es_client, err := elasticsearch.NewClient(
		elasticsearch.Config{
        	Addresses: []string{es_host},
			Username: "elastic",
			Password: "gsaadmin",
    	},
		// elastic.SetURL(host),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	
	return es_client
}

func validate_es_instance(es_client *elasticsearch.Client) {
	res, err := es_client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
	log.Printf("Elasticsearch Initializing..")
}
*/

func main() {
	// log.Println("main")
	log.Printf("es_host url : [%s], index_name : [%s]", es_host, index_name)
	
	es_client := util_elasticsearch.Get_es_instance(es_host)
	util_elasticsearch.Validate_es_instance(es_client)
}