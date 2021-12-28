package init

import (
	"context"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/panforyou/seller-finding/go/services/crawler"

	"github.com/panforyou/seller-finding/go/pkg/logging"
	"github.com/spf13/cobra"
)

// 初回限定。ローカルで動かす前提

func Cmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "初期データ登録",
		Long:  `初期データ登録`,
		Run: func(cmd *cobra.Command, args []string) {
			logging.Init("init")
			if err := InitializeSeller(ctx); err != nil {
				panic(err)
			}
		},
	}
}

func InitializeSeller(_ context.Context) error {
	/*
		エリア一覧の取得
	*/
	areas := map[string]string{}
	c := colly.NewCollector(
		colly.AllowedDomains("tabelog.com"),
	)

	c.OnHTML("a.area-cat-list__item-target", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		areas[e.Text] = url[8 : len(url)-6]
	})

	if err := c.Visit("https://tabelog.com/matome/area_lst/"); err != nil {
		return err
	}

	/*
		エリアごとにパン屋の取得
	*/
	const template = `https://tabelog.com/%s/rstLst/pan/%d/?Srt=D&SrtT=nod`

	var sellers []*crawler.Seller
	page := 60
	for _, area := range areas {
		results, _, err := crawler.SellerCrawling(
			crawler.SellerCrawlParam{
				BaseURL:   fmt.Sprintf(template, area),
				LastVisit: "",
				Domains:   []string{"tabelog.com"},
				MaxPage:   &page,
			})
		if err != nil {
			return err
		}

		sellers = append(sellers, results...)
	}

	logging.Suger().Infof("seller total: %d \n", len(sellers))

	return nil
}
