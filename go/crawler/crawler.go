package crawler

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/panforyou/seller-finding/go/crawler/pkg/logging"

	"github.com/gocolly/colly"
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

	sellers := []*Seller{}
	for i := 1; i <= page; i++ {

		c := colly.NewCollector(
			colly.AllowedDomains(opts.Domains...),
		)
		dc := c.Clone()

		sc := &SellerCrawler{
			collector:       c,
			detailCollector: dc,

			latestVisit: opts.LastVisit,
			sellers:     []*Seller{},
		}

		c.OnHTML("a.list-rst__rst-name-target", sc.detailPage)
		// c.OnHTML("a.c-pagination__arrow--next", sc.nextPage)
		dc.OnHTML("div#rst-data-head", sc.detailExtract)

		if err := c.Visit(fmt.Sprintf(opts.BaseURL, i)); err != nil {
			return nil, "", err
		}

	}

	if len(sellers) == 0 {
		return nil, "", nil
	}

	return sellers, sellers[0].Url, nil
}

type SellerCrawler struct {
	collector       *colly.Collector
	detailCollector *colly.Collector

	latestVisit string
	sellers     []*Seller
}

// nextPage 次のページへ
//func (s *SellerCrawler) nextPage(e *colly.HTMLElement) {
//	link := e.Attr("href")
//	if link == "" {
//		return
//	}
//
//	// e.Request.Visit(link)
//}

// detailPage 一覧から詳細へ
func (s *SellerCrawler) detailPage(e *colly.HTMLElement) {
	link := e.Attr("href")
	if link == "" || link == s.latestVisit {
		return
	}

	s.detailCollector.Visit(link)
}

// detailExtract 詳細ページから情報を抽出
func (s *SellerCrawler) detailExtract(e *colly.HTMLElement) {
	storeName := e.ChildText("div.rstinfo-table__name-wrap > span")
	tel := e.ChildText("p.rstinfo-table__tel-num-wrap > strong")

	prefecture := ""
	address := ""
	e.ForEach("p.rstinfo-table__address > span", func(i int, el *colly.HTMLElement) {
		if i == 0 {
			prefecture = el.Text
		}

		address += el.Text
	})

	openAt := e.ChildText("p.rstinfo-opened-date")
	t, _ := time.Parse("2006年1月2日", openAt)

	// 緯度・経度を取得（google map のデータソースから無理やり取ってくる）
	mapImg := e.ChildAttr("div.rstinfo-table__map img", "data-original")
	posStart := strings.Index(mapImg, "center=") + 7
	posEnd := strings.Index(mapImg, "&style=")
	pos := strings.Split(mapImg[posStart:posEnd], ",")
	lat, _ := strconv.ParseFloat(pos[0], 64)
	lon, _ := strconv.ParseFloat(pos[1], 64)

	seller := &Seller{
		Url:        path.Join(e.Request.URL.Scheme, e.Request.URL.Host, e.Request.URL.Path),
		Key:        e.Request.URL.Path,
		Name:       storeName,
		Tel:        tel,
		Prefecture: prefecture,
		Address:    address,
		Longitude:  lon,
		Latitude:   lat,
		Disabled:   false,
		OpenedAt:   t,
	}
	s.sellers = append(s.sellers, seller)
	logging.Suger().Infof("seller: %+v", seller)
}
