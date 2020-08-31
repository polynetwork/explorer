package service

import (
	"fmt"
	"github.com/polynetwork/explorer/internal/coinmarketcap"
	"github.com/polynetwork/explorer/internal/log"
	"github.com/polynetwork/explorer/internal/model"
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
		err := srv.updateAssetStatisticsByCoinPrice(needUpdatedHistory, coinPrice)
		if err != nil {
			return
		}
	}
	srv.dao.UpdateAssetStatistics(needUpdatedHistory, start)
	//
	latestUpdated := srv.latestUpdated(start, end)
	if latestUpdated == nil {
		return
	}
	err := srv.updateAssetStatisticsByCoinPrice(latestUpdated, coinPrice)
	if err != nil {
		return
	}
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
			assetStatistic.Amount_usd = 0
			assetStatistic.Amount_btc = 0
		} else {
			assetStatistic.Amount_usd = uint64(float64(assetStatistic.Amount) * coinPrice)
			assetStatistic.Amount_btc = uint64(float64(assetStatistic.Amount_usd) / bitcoinPrice)
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
			Amount: 0,
			Amount_usd: 0,
			Amount_btc: 0,
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
				Amount_usd: 0,
				Amount_btc: 0,
				TxNum: assetTxInfo.TxNum,
				LatestUpdate: start,
			}
			assetStatistics = append(assetStatistics, statistic)
		}
	}
	return assetStatistics
}
