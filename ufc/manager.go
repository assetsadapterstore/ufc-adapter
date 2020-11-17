/*
 * Copyright 2020 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package ufc

import (
	"fmt"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/whitecoin-adapter/libs/config"
	bt "github.com/blocktree/whitecoin-adapter/libs/types"
	"github.com/blocktree/whitecoin-adapter/whitecoin"
)

const (
	ChainIDUFC = "24938a99198d850bb7d79010c1325fb63fde63e8e477a5443ff5ce50ab867055"
)

type WalletManager struct {
	*whitecoin.WalletManager
}

func NewWalletManager() *WalletManager {
	config.Add(config.ChainConfig{
		Name:      "UFC",
		CoreAsset: "UFC",
		Prefix:    "UFC",
		ID:        ChainIDUFC,
	})
	config.SetCurrent(ChainIDUFC)

	wm := WalletManager{}
	wm.WalletManager = whitecoin.NewWalletManager()
	wm.Config = whitecoin.NewConfig(Symbol)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.DecoderV2 = NewAddressDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.Api = whitecoin.NewWalletClient(wm.Config.ServerAPI, wm.Config.ServerAPI, false)
	return &wm
}

func (wm *WalletManager) GetRequiredFee(ops []bt.Operation, assetID string) ([]bt.AssetAmount, error) {
	resp := make([]bt.AssetAmount, 0)

	if assetID == "1.3.0" {
		//UFC写死1.3.0
		xwcFees := bt.AssetAmount{
			Asset:  bt.AssetIDFromObject(bt.NewAssetID("1.3.0")),
			Amount: bt.Int64(wm.Config.FixFees),
		}

		resp = append(resp, xwcFees)
	}

	if len(resp) == 0 {
		return nil, fmt.Errorf("can not find required fee with asset ID: %s", assetID)
	}

	return resp, nil
}
