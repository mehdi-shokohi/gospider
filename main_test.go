package main

import (
	"fmt"
	"net/url"
	"sync"
	"testing"

	"github.com/jaeles-project/gospider/core"
)



func TestMain(t *testing.T) {
	site, err := url.Parse("http://karo.school")
	if err != nil {
		fmt.Errorf("Failed to parse : %s", err)
		
	}
	var wg sync.WaitGroup
	crawler := core.NewCrawler(site, &core.CrawlerOptions{Quiet: false,MaxDepth:0})
		crawler.Start(true)
		

	// Brute force Sitemap path

		 go core.ParseSiteMap(site, crawler, crawler.C)
		wg.Add(1)

	// Find Robots.txt

		 core.ParseRobots(site, crawler, crawler.C)
		wg.Wait()
}


