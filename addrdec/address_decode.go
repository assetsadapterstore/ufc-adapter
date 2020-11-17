package addrdec

import (
	"fmt"
	"strings"

	"github.com/blocktree/openwallet/v2/openwallet"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	UFCPublicKeyPrefix       = "PUB_"
	UFCPublicKeyK1Prefix     = "PUB_K1_"
	UFCPublicKeyR1Prefix     = "PUB_R1_"
	UFCPublicKeyPrefixCompat = "UFC"

	//UFC stuff
	UFC_mainnetPublic = addressEncoder.AddressType{"ufc", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(UFCPublicKeyPrefixCompat), nil}
	// UFC_mainnetPrivateWIF           = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	// UFC_mainnetPrivateWIFCompressed = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	openwallet.AddressDecoderV2Base
	IsTestNet bool
}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2() *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	return &decoder
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string, opts ...interface{}) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, UFCPublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(UFCPublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, UFCPublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(UFCPublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, UFCPublicKeyPrefixCompat) { // "UFC"
		pubKeyMaterial = pubKey[len(UFCPublicKeyPrefixCompat):] // strip "UFC"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", UFCPublicKeyK1Prefix, UFCPublicKeyR1Prefix, UFCPublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(UFC_mainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, UFC_mainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, UFC_mainnetPublic.ChecksumType))
	return string(UFC_mainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", UFC_mainnetPublic.Alphabet), nil
}
