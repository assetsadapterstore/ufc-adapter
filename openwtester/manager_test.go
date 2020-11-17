package openwtester

import (
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
)

var (
	testApp        = "assets-adapter"
	configFilePath = filepath.Join("conf")
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"UFC",
	}

	return openw.NewWalletManager(tc)
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO UFC", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "W4pPxx4E4tM1TJWiAyMcGSFMbE8c68URbW")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WJw2FiwwbyP5zyk9YMWi46coeeCDWgi5Mx"
	account := &openwallet.AssetsAccount{Alias: "zbbob111", WalletID: walletID, Required: 1, Symbol: "UFC", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WJw2FiwwbyP5zyk9YMWi46coeeCDWgi5Mx"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WJw2FiwwbyP5zyk9YMWi46coeeCDWgi5Mx"
	accountID := "2cNrcyg8ZQrCDy9BkMA6pSwRncRtipT7FNEPeC8tMTaU"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 1)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WJw2FiwwbyP5zyk9YMWi46coeeCDWgi5Mx"
	accountID := "Fej7cNjwXmr54kno1C8irVRULdufwZhYKVLcf6pLTz3z" //zbbob111 UFC5KvED5xiNiLvwyV8nH5BUH1eTm2uqRrvu5WgwG6CJYkqHiDocD

	// walletID := "W4pPxx4E4tM1TJWiAyMcGSFMbE8c68URbW"
	// accountID := "7z1CyywoB1vC1xnvV4kbvVk1mMjQumV4Vmgp3BbHsHJo" //zbalice999 UFC5wroY2sDyzPMjF9z3EoBugCnDLP1KSoaFXbb3MAm4gWhUAyFb4

	// walletID := "WJw2FiwwbyP5zyk9YMWi46coeeCDWgi5Mx"
	// accountID := "2cNrcyg8ZQrCDy9BkMA6pSwRncRtipT7FNEPeC8tMTaU" //zbcat999 UFC6aWYcRkhqAUPLEkf89eZWwnJBxexSkK6nNwonaB9rQaekNR5WQ UFC7USYmVdEkfzVyePSsawE6VGNFGHsQp8uG2fkRW6MNKKaVWHWFi

	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("address[", i, "] :", w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}
