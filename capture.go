package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
)

func main()  {
	ctx,cancel := chromedp.NewContext(context.Background(),chromedp.WithDebugf(log.Printf))
	defer cancel()

	url := "https://search.bilibili.com/all?keyword=gin%20%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6&from_source=nav_search_new"
	filename := "url.png"
	var imageBuf []byte
	
	err := chromedp.Run(ctx,ScreenshotTasks(url,&imageBuf))
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename,imageBuf,0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("任务完成")
}

func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {
			*imageBuf, err = page.CaptureScreenshot().WithQuality(90).Do(ctx)
			return err
		}),
	}
}