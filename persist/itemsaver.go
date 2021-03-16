package persist

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/public"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() (chan engine.Item,error){

	client, err := elastic.NewClient(elastic.SetURL(public.EsUrl),elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	saverChan := make(chan engine.Item)

	go func() {
		itemCount:=0
		for{
			item := <- saverChan
			log.Printf("ItemSaver Got Item #%d : %s\n",itemCount,item)
			//err := Save(item.(engine.Item))
			err := Save(client, item)
			if err != nil {
				log.Printf("ItemSaver err : %v\n",err)
			}
			itemCount++
		}
	}()
	return saverChan,nil
}

func Save(client *elastic.Client,item engine.Item) error  {

	if item.Type == "" {
		return errors.New("Item Type is empty")
	}

	// Index() 是保存数据
	indexService := client.Index().Index(public.EsIndex).Type(public.EsType)
	if item.Id != "" {
		//判断是否存在
		_ , err := client.Get().Index(public.EsIndex).Type(public.EsType).Id(item.Id).Do(context.Background())
		if err == nil {
			return errors.New("已经存在 id: "+item.Id)
		}
		indexService = indexService.Id(item.Id)
	}

	_ ,err := indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	//client.Flush()
	return nil
}
