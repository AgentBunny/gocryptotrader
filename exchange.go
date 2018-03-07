package main

import (
	"errors"
	"log"

	"github.com/agentbunny/gocryptotrader/common"
	exchange "github.com/agentbunny/gocryptotrader/exchanges"
	"github.com/agentbunny/gocryptotrader/exchanges/anx"
	"github.com/agentbunny/gocryptotrader/exchanges/binance"
	"github.com/agentbunny/gocryptotrader/exchanges/bitfinex"
	"github.com/agentbunny/gocryptotrader/exchanges/bitflyer"
	"github.com/agentbunny/gocryptotrader/exchanges/bithumb"
	"github.com/agentbunny/gocryptotrader/exchanges/bitstamp"
	"github.com/agentbunny/gocryptotrader/exchanges/bittrex"
	"github.com/agentbunny/gocryptotrader/exchanges/btcc"
	"github.com/agentbunny/gocryptotrader/exchanges/btcmarkets"
	"github.com/agentbunny/gocryptotrader/exchanges/coinut"
	"github.com/agentbunny/gocryptotrader/exchanges/exmo"
	"github.com/agentbunny/gocryptotrader/exchanges/gdax"
	"github.com/agentbunny/gocryptotrader/exchanges/gemini"
	"github.com/agentbunny/gocryptotrader/exchanges/hitbtc"
	"github.com/agentbunny/gocryptotrader/exchanges/huobi"
	"github.com/agentbunny/gocryptotrader/exchanges/itbit"
	"github.com/agentbunny/gocryptotrader/exchanges/kraken"
	"github.com/agentbunny/gocryptotrader/exchanges/lakebtc"
	"github.com/agentbunny/gocryptotrader/exchanges/liqui"
	"github.com/agentbunny/gocryptotrader/exchanges/localbitcoins"
	"github.com/agentbunny/gocryptotrader/exchanges/okcoin"
	"github.com/agentbunny/gocryptotrader/exchanges/okex"
	"github.com/agentbunny/gocryptotrader/exchanges/poloniex"
	"github.com/agentbunny/gocryptotrader/exchanges/wex"
	"github.com/agentbunny/gocryptotrader/exchanges/yobit"
)

// vars related to exchange functions
var (
	ErrNoExchangesLoaded     = errors.New("no exchanges have been loaded")
	ErrExchangeNotFound      = errors.New("exchange not found")
	ErrExchangeAlreadyLoaded = errors.New("exchange already loaded")
	ErrExchangeFailedToLoad  = errors.New("exchange failed to load")
)

// CheckExchangeExists returns true whether or not an exchange has already
// been loaded
func CheckExchangeExists(exchName string) bool {
	for x := range bot.exchanges {
		if common.StringToLower(bot.exchanges[x].GetName()) == common.StringToLower(exchName) {
			return true
		}
	}
	return false
}

// GetExchangeByName returns an exchange given an exchange name
func GetExchangeByName(exchName string) exchange.IBotExchange {
	for x := range bot.exchanges {
		if common.StringToLower(bot.exchanges[x].GetName()) == common.StringToLower(exchName) {
			return bot.exchanges[x]
		}
	}
	return nil
}

// ReloadExchange loads an exchange config by name
func ReloadExchange(name string) error {
	nameLower := common.StringToLower(name)

	if len(bot.exchanges) == 0 {
		return ErrNoExchangesLoaded
	}

	if !CheckExchangeExists(nameLower) {
		return ErrExchangeNotFound
	}

	exchCfg, err := bot.config.GetExchangeConfig(name)
	if err != nil {
		return err
	}

	e := GetExchangeByName(nameLower)
	e.Setup(exchCfg)
	log.Printf("%s exchange reloaded successfully.\n", name)
	return nil
}

// UnloadExchange unloads an exchange by
func UnloadExchange(name string) error {
	nameLower := common.StringToLower(name)

	if len(bot.exchanges) == 0 {
		return ErrNoExchangesLoaded
	}

	if !CheckExchangeExists(nameLower) {
		return ErrExchangeNotFound
	}

	exchCfg, err := bot.config.GetExchangeConfig(name)
	if err != nil {
		return err
	}

	exchCfg.Enabled = false
	err = bot.config.UpdateExchangeConfig(exchCfg)
	if err != nil {
		return err
	}

	for x := range bot.exchanges {
		if bot.exchanges[x].GetName() == name {
			bot.exchanges[x].SetEnabled(false)
			bot.exchanges = append(bot.exchanges[:x], bot.exchanges[x+1:]...)
			return nil
		}
	}

	return ErrExchangeNotFound
}

// LoadExchange loads an exchange by name
func LoadExchange(name string) error {
	nameLower := common.StringToLower(name)
	var exch exchange.IBotExchange

	if len(bot.exchanges) > 0 {
		if CheckExchangeExists(nameLower) {
			return ErrExchangeAlreadyLoaded
		}
	}

	switch nameLower {
	case "anx":
		exch = new(anx.ANX)
	case "binance":
		exch = new(binance.Binance)
	case "bitfinex":
		exch = new(bitfinex.Bitfinex)
	case "bitflyer":
		exch = new(bitflyer.Bitflyer)
	case "bithumb":
		exch = new(bithumb.Bithumb)
	case "bitstamp":
		exch = new(bitstamp.Bitstamp)
	case "bittrex":
		exch = new(bittrex.Bittrex)
	case "btcc":
		exch = new(btcc.BTCC)
	case "btc markets":
		exch = new(btcmarkets.BTCMarkets)
	case "coinut":
		exch = new(coinut.COINUT)
	case "exmo":
		exch = new(exmo.EXMO)
	case "gdax":
		exch = new(gdax.GDAX)
	case "gemini":
		exch = new(gemini.Gemini)
	case "hitbtc":
		exch = new(hitbtc.HitBTC)
	case "huobi":
		exch = new(huobi.HUOBI)
	case "itbit":
		exch = new(itbit.ItBit)
	case "kraken":
		exch = new(kraken.Kraken)
	case "lakebtc":
		exch = new(lakebtc.LakeBTC)
	case "liqui":
		exch = new(liqui.Liqui)
	case "localbitcoins":
		exch = new(localbitcoins.LocalBitcoins)
	case "okcoin china":
		exch = new(okcoin.OKCoin)
	case "okcoin international":
		exch = new(okcoin.OKCoin)
	case "okex":
		exch = new(okex.OKEX)
	case "poloniex":
		exch = new(poloniex.Poloniex)
	case "wex":
		exch = new(wex.WEX)
	case "yobit":
		exch = new(yobit.Yobit)
	default:
		return ErrExchangeNotFound
	}

	if exch == nil {
		return ErrExchangeFailedToLoad
	}

	exch.SetDefaults()
	bot.exchanges = append(bot.exchanges, exch)
	exchCfg, err := bot.config.GetExchangeConfig(name)
	if err != nil {
		return err
	}

	exchCfg.Enabled = true
	exch.Setup(exchCfg)
	exch.Start()
	return nil
}

// SetupExchanges sets up the exchanges used by the bot
func SetupExchanges() {
	for _, exch := range bot.config.Exchanges {
		if CheckExchangeExists(exch.Name) {
			e := GetExchangeByName(exch.Name)
			if e == nil {
				log.Println(ErrExchangeNotFound)
				continue
			}

			err := ReloadExchange(exch.Name)
			if err != nil {
				log.Printf("ReloadExchange %s failed: %s", exch.Name, err)
				continue
			}

			if !e.IsEnabled() {
				UnloadExchange(exch.Name)
				continue
			}
			return

		}
		if !exch.Enabled {
			log.Printf("%s: Exchange support: Disabled", exch.Name)
			continue
		} else {
			err := LoadExchange(exch.Name)
			if err != nil {
				log.Printf("LoadExchange %s failed: %s", exch.Name, err)
				continue
			}
		}
		log.Printf(
			"%s: Exchange support: Enabled (Authenticated API support: %s - Verbose mode: %s).\n",
			exch.Name,
			common.IsEnabled(exch.AuthenticatedAPISupport),
			common.IsEnabled(exch.Verbose),
		)
	}
}
