package persist

import (
	"Book/Crawler/engine"
	"Book/Models"
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		//itemCount :=0
		for {
			item := <-out
			err := save(client, item)
			if err != nil {
				log.Printf("Item Saver:error saving item %v: %v", item, err)
			}
			//log.Printf("Item Saver:got item #%d:%v",itemCount,item)
			//itemCount++
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}
	indexServer := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexServer.Id(item.Id)
	}
	_, err := indexServer.
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// 查找
func Search(id string) Models.Book {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index("dating_profile").
		Type("book").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual Models.Book
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	return actual
}
