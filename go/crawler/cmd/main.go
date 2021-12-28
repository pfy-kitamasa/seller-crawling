package main

import (
	"context"
	"log"

	"github.com/panforyou/seller-finding/go/crawler"
	"github.com/panforyou/seller-finding/go/crawler/repository"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	// 環境変数の取得
	var c crawler.Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// DB コネクト
	ctx := context.Background()
	db, err := crawler.NewDB(c.DB.Host, c.DB.Port, c.DB.Name, c.DB.User, c.DB.Password, c.Env)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// クローラー対象のドメイン情報を取得
	domain, err := db.GetDomainByID(ctx, "tabelog")
	if err != nil {
		log.Fatal(err.Error())
	}

	// クローラー
	sellers, latest, err := crawler.SellerCrawling(domain.Url, domain.LatestVisitUrl, domain.AllowDomains)
	if err != nil {
		log.Fatal(err.Error())
	}

	// DB へ保存
	for _, seller := range sellers {
		_, err := db.CreateSeller(ctx, repository.CreateSellerParams{
			Url:        seller.Url,
			Key:        seller.Key,
			Name:       seller.Name,
			Tel:        seller.Tel,
			Prefecture: seller.Prefecture,
			Address:    seller.Address,
			Longitude:  seller.Longitude,
			Latitude:   seller.Latitude,
			OpenedAt:   seller.OpenedAt,
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if err := db.UpdateDomainLastedVisitURL(ctx, repository.UpdateDomainLastedVisitURLParams{
		Key:            domain.Key,
		LatestVisitUrl: latest,
	}); err != nil {
		log.Fatal(err.Error())
	}
}
