package service

import (
	"fmt"
	"github.com/polynetwork/explorer/internal/coinmarketcap"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
	"math"
	"math/big"
	"time"
)

func (srv *Service) DoStatistic() {
	now := time.Now()
	nowUnix := uint32(now.Unix())
	end := (nowUnix / 60)
	start := end - srv.c.Server.StatisticTimeslot
	if end % srv.c.Server.StatisticTimeslot != 0 {
		return
	}
	log.Infof("do statistic at time: %s", now.Format("2006-01-02 15:04:05"))
	end = end * 60
	start = start * 60
	//
	coins := make([]string, 0)
	for _, item := range srv.tokens {
		coins = append(coins, item.Name)
	}
	coinPrice := srv.getCoinPrice(coins)
	if coinPrice == nil {
		return
	}
	//
	needUpdatedHistory := srv.checkHistory(start)
	if needUpdatedHistory != nil {
		needUpdatedHistory = srv.updatePrecision(needUpdatedHistory)
		err := srv.updateAssetStatisticsByCoinPrice(needUpdatedHistory, coinPrice)
		if err != nil {
			return
		}
		srv.checkAssetStatistics(needUpdatedHistory)
		srv.dao.UpdateAssetStatistics(needUpdatedHistory, start)
	}
	//
	latestUpdated := srv.latestUpdated(start, end)
	if latestUpdated == nil {
		return
	}
	latestUpdated = srv.updatePrecision(latestUpdated)
	err := srv.updateAssetStatisticsByCoinPrice(latestUpdated, coinPrice)
	if err != nil {
		return
	}
	srv.checkAssetStatistics(latestUpdated)
	srv.dao.UpdateAssetStatistics(latestUpdated, end)
}

func (srv *Service) checkHistory(tt uint32) (res []*model.AssetStatistic) {
	assetStatistics, err := srv.dao.SelectAssetStatistic(tt)
	if err != nil {
		log.Errorf("SelectAssetStatistic err: %s", err.Error())
		return nil
	}
	if assetStatistics == nil || len(assetStatistics) == 0 {
		return nil
	}
	for _, assetStatistic := range assetStatistics {
		assetTxInfo, err := srv.dao.SelectAssetHistory(assetStatistic.LatestUpdate, tt, assetStatistic.Name)
		if err != nil {
			log.Errorf("SelectAssetHistory err: %s", err.Error())
			assetStatistic.LatestUpdate = 1
			continue
		}
		if assetTxInfo == nil {
			assetStatistic.TxNum = 0
			assetStatistic.Amount = big.NewInt(0)
			continue
		}
		if assetTxInfo.Name == "" || assetTxInfo.TxNum == 0 {
			assetStatistic.LatestUpdate = 1
			continue
		}
		assetStatistic.TxNum = assetTxInfo.TxNum
		assetStatistic.Amount = assetTxInfo.Amount
	}
	return assetStatistics
}

func (srv *Service) getCoinPrice(coins []string) map[string]float64 {
	//
	var cmcSdk *coinmarketcap.CoinMarketCapSdk
	if srv.c.CoinMarketCap == nil || srv.c.CoinMarketCap.Url == "" {
		cmcSdk = coinmarketcap.DefaultCoinMarketCapSdk()
	} else {
		cmcSdk = coinmarketcap.NewCoinMarketCapSdk(srv.c.CoinMarketCap.Url, srv.c.CoinMarketCap.AppKey)
	}
	listings, err := cmcSdk.ListingsLatest()
	if err != nil {
		log.Errorf("CoinMarketCap ListingsLatest err: %s", err.Error())
		return nil
	}
	//
	coinMaps := make(map[string]string, 0)
	for _, listing := range listings {
		coinMaps[listing.Name] = fmt.Sprintf("%d", listing.ID)
	}
	//
	coinIds := ""
	coinId, ok := coinMaps["Bitcoin"]
	if !ok {
		log.Errorf("There is no coin Bitcoin in CoinMarketCap!")
		return nil
	}
	coinIds += coinId
	//
	for _, coin := range coins {
		coinId, ok := coinMaps[coin]
		if !ok {
			log.Warnf("There is no coin %s in CoinMarketCap!", coin)
			continue
		}
		coinIds += ","
		coinIds += coinId
	}
	//
	quotes, err := cmcSdk.QuotesLatest(coinIds)
	if err != nil {
		log.Errorf("CoinMarketCap QuotesLatest err: %s", err.Error())
		return nil
	}
	//
	coinPrice := make(map[string]float64)
	for _, v := range quotes {
		name := v.Name
		if v.Quote == nil || v.Quote["USD"] == nil {
			log.Errorf(" There is no price for coin %s in CoinMarketCap!", name)
			return nil
		}
		coinPrice[name] = v.Quote["USD"].Price
	}
	return coinPrice
}

func (srv *Service) updateAssetStatisticsByCoinPrice(assetStatistics []*model.AssetStatistic, coinPrices map[string]float64) (err error) {
	bitcoinPrice, ok := coinPrices["Bitcoin"]
	if !ok {
		log.Errorf("There is no coin Bitcoin!")
		return fmt.Errorf("There is no coin Bitcoin!")
	}
	for _, assetStatistic := range assetStatistics {
		coinPrice, ok := coinPrices[assetStatistic.Name]
		if !ok {
			log.Warnf("There is no coin %s!", assetStatistic.Name)
			assetStatistic.Amount_usd = big.NewInt(0)
			assetStatistic.Amount_btc = big.NewInt(0)
		} else {
			amount := new(big.Int).Mul(assetStatistic.Amount, big.NewInt(int64(coinPrice * 10000)))
			assetStatistic.Amount_usd = amount
			amount_btc := new(big.Int).Div(assetStatistic.Amount_usd, big.NewInt(int64(bitcoinPrice)))
			assetStatistic.Amount_btc = amount_btc
		}
	}
	return nil
}

func (srv *Service) latestUpdated(start uint32, end uint32) (res []*model.AssetStatistic) {
	assetAddressNums, err := srv.dao.SelectAssetAddressNum()
	if err != nil {
		log.Errorf("SelectAssetAddressNum err: %s", err.Error())
		return
	}
	assetTxInfos, err := srv.dao.SelectAssetTxInfo(start, end)
	if err != nil {
		log.Errorf("SelectAssetTxInfo err: %s", err.Error())
		return
	}
	assetStatistics := make([]*model.AssetStatistic, 0)
	for _, assetAddressNum := range assetAddressNums {
		statistic := &model.AssetStatistic{
			Name: assetAddressNum.Name,
			Addressnum: assetAddressNum.AddNum,
			Amount: big.NewInt(0),
			Amount_usd: big.NewInt(0),
			Amount_btc: big.NewInt(0),
			TxNum: 0,
			LatestUpdate: start,
		}
		assetStatistics = append(assetStatistics, statistic)
	}
	for _, assetTxInfo := range assetTxInfos {
		var statistic *model.AssetStatistic
		statistic = nil
		for _, item := range assetStatistics {
			if item.Name == assetTxInfo.Name {
				statistic = item
				statistic.Amount = assetTxInfo.Amount
				statistic.TxNum = assetTxInfo.TxNum
				break
			}
		}
		if statistic == nil {
			statistic = &model.AssetStatistic{
				Name: assetTxInfo.Name,
				Addressnum: 0,
				Amount: assetTxInfo.Amount,
				Amount_usd: big.NewInt(0),
				Amount_btc: big.NewInt(0),
				TxNum: assetTxInfo.TxNum,
				LatestUpdate: start,
			}
			assetStatistics = append(assetStatistics, statistic)
		}
	}
	return assetStatistics
}

func (srv *Service) updatePrecision(assetStatistics []*model.AssetStatistic) []*model.AssetStatistic {
	res := make([]*model.AssetStatistic, 0)
	for _, item := range assetStatistics {
		precision := srv.GetTokenPrecision(item.Name)
		if precision == 0 {
			log.Errorf("updatePrecision err, the precision of  token: %s is missing", item.Name)
			continue
		}
		item.Amount = item.Amount.Div(item.Amount, big.NewInt(int64(precision)))
		res = append(res, item)
	}
	return res
}

func (srv *Service) checkAssetStatistics(assetStatistics []*model.AssetStatistic) {
	for _, item := range assetStatistics {
		if item.Amount.Int64() > math.MaxInt64 {
			log.Errorf("checkAssetStatistics err, the amount of  token: %s is too big", item.Name)
			item.Amount.SetInt64(math.MaxInt64)
		}
		if item.Amount_btc.Int64() > math.MaxInt64 {
			log.Errorf("checkAssetStatistics err, the btc amount of  token: %s is too big", item.Name)
			item.Amount_btc.SetInt64(math.MaxInt64)
		}
		if item.Amount_usd.Int64() > math.MaxInt64 {
			log.Errorf("checkAssetStatistics err, the amount of  token: %s is too big", item.Name)
			item.Amount_usd.SetInt64(math.MaxInt64)
		}
	}
}
