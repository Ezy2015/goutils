package eslite

import (
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v3"
)

type ElasticClient struct {
	client *elastic.Client
	bkt    *elastic.BulkService
}

func (es *ElasticClient) Open(host string, port int) error {
	url := fmt.Sprintf("http://%s:%d", host, port)
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return err
	}
	info, code, err := client.Ping("http://127.0.0.1:9200").Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
	es.client = client

	/*
		BulkService will be reset after each Do call.
		In other words, you can reuse BulkService to send many batches.
		You do not have to create a new BulkService for each batch.
	*/
	es.bkt = es.client.Bulk()
	return nil
}

func (es *ElasticClient) Write(index string, id string,
	typ string, v interface{}) error {

	es.bkt.Add(elastic.NewBulkIndexRequest().Index(
		index).Type(typ).Id(id).Doc(v))

	return nil
}

func (es *ElasticClient) Begin() error {
	return nil
}

func (es *ElasticClient) Commit() error {
	log.Println("DOBEFORE bulkRequest:NumberOfActions", es.bkt.NumberOfActions())

	bulkResponse, err := es.bkt.Do()
	if err != nil {
		log.Panic(err)
		return err
	}
	if bulkResponse == nil {
		log.Fatal("expected bulkResponse to be != nil; got nil")
	}
	log.Println("DOAFTER buolkRequest:NumberOfActions", es.bkt.NumberOfActions())
	return nil
}

func (es *ElasticClient) Close() {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := es.client.IndexExists("twitter").Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := es.client.CreateIndex("twitter").Do()
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

}

func (es *ElasticClient) WriteDirect(index string, id string,
	typ string, v interface{}) error {
	_, err := es.client.Index().Index(index).Type(typ).Id(id).BodyJson(v).Do()
	return err
}
