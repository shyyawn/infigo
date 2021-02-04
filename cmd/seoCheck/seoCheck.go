package seoCheck

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	c "github.com/logrusorgru/aurora"
	"github.com/shyyawn/go-to/config"
	log "github.com/shyyawn/go-to/logging"
	"github.com/shyyawn/infigo/pkg/data"
	"github.com/spf13/cobra"
	"net/http"
	"sync"
)

var alias string

var Cmd = &cobra.Command{
	Use:   "seoCheck",
	Short: "Just checks if the seo tags are present on the page and prints them for someone to read.",
	Long:  "Just checks if the seo tags are present on the page and prints them for someone to read.",
	Run:   seoCheck,
}

func seoCheck(cmd *cobra.Command, args []string) {

	var u data.Urls
	configDir := config.AppConfig.GetString("runtime_dir") + "/config"

	u.LoadUrls(configDir + "/seoCheck/urls.yml")

	domain, urls := u.GetUrls(alias)
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		GetMetas(&wg, domain+url)
		//go GetMetas(&wg, domain + url)
	}
	wg.Wait()
}

func GetMetas(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := http.Get(url)
	if err != nil {
		log.Warn("URL Failed: " + url)
		log.Error(err)
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
	detail := fmt.Sprintln(c.BgGreen("Meta Information for " + url))
	title := doc.Find("title")
	detail = detail + fmt.Sprintf("%-10s %s \n", c.Green("Title"), title.Text())
	desc, ok := doc.Find("meta[name='description']").Attr("content")
	if !ok {
		desc = "[[ MISSING ]]"
	}
	detail = detail + fmt.Sprintf("%-10s %s \n", c.Green("Desc"), desc)
	fmt.Println(detail)
}

func init() {
	log.Info("Init SEO Check")
	Cmd.Flags().StringVarP(&alias, "alias", "a", "", "This will get the urls")
	_ = Cmd.MarkFlagRequired("alias")
}
