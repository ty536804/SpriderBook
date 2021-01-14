package persist

import (
	"Book/Models"
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		//itemCount :=0
		for {
			item := <-out
			save(item)
			//log.Printf("Item Saver:got item #%d:%v",itemCount,item)
			//itemCount++
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().
		Index("dating_profile").
		Type("book").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
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
