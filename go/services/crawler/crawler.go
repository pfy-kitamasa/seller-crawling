package crawler

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/panforyou/seller-finding/go/pkg/logging"
)

type SellerCrawlParam struct {
	BaseURL   string
	LastVisit string
	Domains   []string
	MaxPage   *int
}

func SellerCrawling(opts SellerCrawlParam) ([]*Seller, string, error) {
	page := 10
	if opts.MaxPage != nil {
		page = *opts.MaxPage
	}

	c := colly.NewCollector(
		colly.AllowedDomains(opts.Domains...),
	)

	// 詳細リンクをキューに追加
	q, _ := queue.New(
		5,
		&queue.InMemoryQueueStorage{MaxSize: 100000},
	)

	isEmpty := false
	for i := 1; i <= page; i++ {
		if isEmpty {
			break
		}

		lc := c.Clone()
		lc.OnHTML("span.c-page-count__num", func(e *colly.HTMLElement) {
			if e.Index != 0 {
				return
			}
			isEmpty = e.Text == "0"
		})
		lc.OnHTML("a.list-rst__rst-name-target", func(e *colly.HTMLElement) {
			if isEmpty {
				return
			}
			link := e.Attr("href")
			q.AddURL(link)
		})

		if err := lc.Visit(fmt.Sprintf(opts.BaseURL, i)); err != nil {
			logging.Suger().Error(err)
			break
		}
	}

	if q.IsEmpty() {
		return nil, "", nil
	}

	var sellers []*Seller
	c.Init()
	dc := c.Clone()
	dc.OnHTML("div#rst-data-head", func(e *colly.HTMLElement) {
		seller := &Seller{
			URL: path.Join(e.Request.URL.Scheme, e.Request.URL.Host, e.Request.URL.Path),
			Key: e.Request.URL.Path,
		}

		if name := e.ChildText("div.rstinfo-table__name-wrap > span"); name != "" {
			seller.Name = &name
		}

		if tel := e.ChildText("p.rstinfo-table__tel-num-wrap > strong"); tel != "" {
			seller.Tel = &tel
		}

		address := ""
		e.ForEach("p.rstinfo-table__address > span", func(i int, el *colly.HTMLElement) {
			if i == 0 && el.Text != "" {
				seller.Prefecture = &el.Text
			}

			address += el.Text
		})
		if address != "" {
			seller.Address = &address
		}

		openAt := e.ChildText("p.rstinfo-opened-date")
		t, err := time.Parse("2006年1月2日", openAt)
		if err == nil {
			seller.OpenedAt = &t
		}

		// 緯度・経度を取得（google map のデータソースから無理やり取ってくる）
		if mapImg := e.ChildAttr("div.rstinfo-table__map img", "data-original"); mapImg != "" {
			posStart := strings.Index(mapImg, "center=") + 7
			posEnd := strings.Index(mapImg, "&style=")
			pos := strings.Split(mapImg[posStart:posEnd], ",")
			lat, err := strconv.ParseFloat(pos[0], 64)
			if err == nil {
				seller.Latitude = &lat
			}

			lon, err := strconv.ParseFloat(pos[1], 64)
			if err == nil {
				seller.Longitude = &lon
			}
		}

		sellers = append(sellers, seller)
		logging.Suger().Infof("sellers %d/%d, seller: %+v", len(sellers), cap(sellers), seller)
	})

	if err := q.Run(dc); err != nil {
		return nil, "", err
	}

	return sellers, sellers[0].URL, nil
}
