package seoCheck

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	c "github.com/logrusorgru/aurora"
	"github.com/shyyawn/go-to/config"
	log "github.com/shyyawn/go-to/logging"
	"github.com/shyyawn/infigo/pkg/data"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "seoCheck",
	Short: "Just checks if the seo tags are present on the page and prints them for someone to read.",
	Long:  "Just checks if the seo tags are present on the page and prints them for someone to read.",
	Run:   seoCheck,
}

func seoCheck(cmd *cobra.Command, args []string) {
	log.Info("SEO Check")
	var u data.Urls
	configDir := config.AppConfig.GetString("runtime_dir") + "/config"
	log.Info(configDir)
	u.LoadUrls(configDir + "/seoCheck/urls.yml")
	log.Info(u)
	domain, urls := u.GetUrls("one-carmudi")
	for _, url := range urls {
		GetMetas(domain + url)
	}
}

func GetMetas(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// @todo: create new function in log package
		log.Fatal(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.BgGreen("Meta Information for " + url))
	title := doc.Find("title")
	fmt.Printf("%-10s %s \n", c.Green("Title"), title.Text())
	desc, ok := doc.Find("meta[name='description']").Attr("content")
	if !ok {
		desc = "[[ MISSING ]]"
	}
	fmt.Printf("%-10s %s \n", c.Green("Desc"), desc)
	fmt.Println("")
}

func init() {
	log.Info("Init SEO Check")
}
