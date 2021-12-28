package regularly

import (
	"context"
	"log"

	"github.com/panforyou/seller-finding/go/ent"
	"github.com/panforyou/seller-finding/go/ent/website"
	"github.com/panforyou/seller-finding/go/pkg/logging"
	"github.com/panforyou/seller-finding/go/pkg/rdb"
	"github.com/panforyou/seller-finding/go/services/crawler"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	_ "github.com/lib/pq"
)

func Cmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "crawling",
		Short: "定期的なクローリング",
		Long:  `初期データ登録`,
		Run: func(cmd *cobra.Command, args []string) {
			logging.Init("crawling")
			if err := Crawling(ctx); err != nil {
				log.Fatalf("failed to crawling: %v", err)
			}
		},
	}
}

func Crawling(_ context.Context) error {
	// 環境変数の取得
	var conf crawler.Config
	err := envconfig.Process("", &conf)
	if err != nil {
		return xerrors.Errorf("failed to load env config: %v", err)
	}

	// DB コネクト
	dbc, err := rdb.NewClient(conf.DB.Host, conf.DB.Port, conf.DB.Name, conf.DB.User, conf.DB.Password)
	if err != nil {
		return err
	}
	defer dbc.Close()

	ctx := context.Background()

	// DB マイグレーション&初期データセット
	if err := rdb.MigrationWithSeed(ctx, dbc, func() error {
		err := dbc.Website.Create().
			SetKey("tabelog").
			SetURLTemplate(`https://tabelog.com/rstLst/pan/%s/?Srt=D&SrtT=nod`).
			SetAllowDomains([]string{"tabelog.com"}).
			SetMaxPage(60).
			OnConflictColumns(website.FieldKey).
			UpdateNewValues().
			Exec(ctx)
		return err
	}); err != nil {
		return xerrors.Errorf("failed to migration with seeding: %v", err)
	}

	// クローラー対象のドメイン情報を取得
	targets, err := dbc.Website.Query().Where().All(ctx)
	if err != nil {
		return xerrors.Errorf("failed to upsert initialize data: %v", err)
	}

	for _, target := range targets {
		// クローラー
		sellers, latest, err := crawler.SellerCrawling(
			crawler.SellerCrawlParam{
				BaseURL:   target.URLTemplate,
				LastVisit: target.LatestVisitURL,
				Domains:   target.AllowDomains,
				MaxPage:   &target.MaxPage,
			})
		if err != nil {
			return xerrors.Errorf("failed to SellerCrawling: %v", err)
		}

		// DB へ保存
		bulk := make([]*ent.SellerCreate, len(sellers))

		for i, seller := range sellers {
			bulk[i] = dbc.Seller.Create().
				SetURL(seller.URL).
				SetKey(seller.Key).
				SetNillableName(seller.Name).
				SetNillableTel(seller.Tel).
				SetNillablePrefecture(seller.Prefecture).
				SetNillableAddress(seller.Address).
				SetNillableLongitude(seller.Longitude).
				SetNillableLatitude(seller.Latitude).
				SetNillableOpenedAt(seller.OpenedAt).
				SetDisabled(false)
		}

		if err := rdb.WithTx(ctx, dbc, func(tx *ent.Tx) error {
			// save seller
			if err := tx.Seller.CreateBulk(bulk...).Exec(ctx); err != nil {
				return xerrors.Errorf("failed to save sellers: %v", err)
			}

			// save crawling latest url
			if err := tx.Website.UpdateOne(target).SetLatestVisitURL(latest).Exec(ctx); err != nil {
				return xerrors.Errorf("failed to update website: %v", err)
			}

			return nil
		}); err != nil {
			return xerrors.Errorf("failed to transactions: %v", err)
		}

		logging.Suger().Infof("crawling %s completed: %d saved", target.Key, len(sellers))
	}
	return nil
}
