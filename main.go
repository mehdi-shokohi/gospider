package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/elazarl/goproxy"

	"github.com/jaeles-project/gospider/core"
)

func main() {
	site, err := url.Parse("https://karo.school")
	if err != nil {
		fmt.Errorf("Failed to parse : %s", err)

	}
	go func() {
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = false
		proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
		// proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)

		proxy.OnRequest().DoFunc(func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
			
			fmt.Println("<PROXY>",r.URL,r.URL.Query())
			return r,nil
		})
		log.Fatal(http.ListenAndServe(":8080", proxy))
	}()
	scanContext := context.Background()
	ctx, cancel := context.WithCancel(scanContext)
	crawler := core.NewCrawler(ctx, site, &core.CrawlerOptions{Proxy:"http://localhost:8080", Quiet:true, MaxDepth: 0, JsonOutput: false})
	crawler.Start(true)
	// Brute force Sitemap path

	go core.ParseSiteMap(site, crawler, crawler.C)

	// Find Robots.txt

	go core.ParseRobots(site, crawler, crawler.C)

	<-time.After(time.Second * 100)
	cancel()

	<-ctx.Done()
	fmt.Println("finished")
}
