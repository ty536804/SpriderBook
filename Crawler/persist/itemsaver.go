package persist

import (
	"context"
	"fmt"
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

func save(item interface{}) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Index().
		Index("dating_profile").
		Type("book").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}
