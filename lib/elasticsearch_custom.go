package lib

import (
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)


func Get_es_instance(es_host string) (*elasticsearch.Client) {
	es_client, err := elasticsearch.NewClient(
		elasticsearch.Config{
        	Addresses: []string{es_host},
			Username: "elastic",
			Password: "gsaadmin",
    	},
		// elastic.SetURL(host),
	)
	if err != nil {
		log.Fatalf("Error getting response [%s]: %s", es_host, err)
	}
	
	return es_client
}

func Validate_es_instance(es_client *elasticsearch.Client) {
	res, err := es_client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)
	log.Printf("Elasticsearch Initializing..")
}