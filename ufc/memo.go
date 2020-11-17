package ufc

import (
	"fmt"

	"github.com/blocktree/whitecoin-adapter/libs/config"
	"github.com/blocktree/whitecoin-adapter/libs/types"
)

//Decrypt calculates a shared secret by the receivers private key
//and the senders public key, then decrypts the given memo message.
func Decrypt(msg, fromPub, toPub string, nonce uint64, wif string) (string, error) {

	if len(msg) == 0 || len(fromPub) == 0 || len(toPub) == 0 {
		return "", nil
	}

	if len(wif) == 0 {
		return "", fmt.Errorf("wif cannot be empty")
	}

	var buf types.Buffer
	ret := config.FindByID(ChainIDUFC)
	if ret == nil {
		config.Add(config.ChainConfig{
			Name:      "UFC",
			CoreAsset: "UFC",
			Prefix:    "UFC",
			ID:        ChainIDUFC,
		})
	}
	config.SetCurrent(ChainIDUFC)

	from, err := types.NewPublicKeyFromString(fromPub)
	if err != nil {
		return "", fmt.Errorf("NewPublicKeyFromString: %v", err)
	}
	to, err := types.NewPublicKeyFromString(toPub)
	if err != nil {
		return "", fmt.Errorf("NewPublicKeyFromString: %v", err)
	}

	buf.FromString(msg)

	memo := types.Memo{
		From:    *from,
		To:      *to,
		Message: buf,
		Nonce:   types.UInt64(nonce),
	}

	priv, err := types.NewPrivateKeyFromWif(wif)
	if err != nil {
		return "", fmt.Errorf("NewPrivateKeyFromWif: %v", err)
	}

	m, err := memo.Decrypt(priv)
	if err != nil {
		return "", fmt.Errorf("Decrypt: %v", err)
	}

	return m, nil
}
