package wallet

import (
	"encoding/hex"
	"errors"
	"fmt"

	"crypto/ed25519"
	"crypto/rand"

	"github.com/dchest/blake2b"
	"github.com/rickiey/btcutil/bech32"
)

const (
	SignatureScheme    = "ED25519"
	PUBLIC_KEY_SIZE    = 32
	SEED_SIZE          = 32
	PRIVATE_KEY_SIZE   = 64
	SUI_ADDRESS_LENGTH = 32
	EmptyStr           = ""
	HRP                = "suiprivkey"
)

func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(rand.Reader)
}

func GenerateKeyFrmSeed(seed string) (ed25519.PrivateKey, error) {
	s, err := hex.DecodeString(seed)
	if err != nil {
		return nil, err
	}

	if len(s) != SEED_SIZE {
		return nil, errors.New("invalid seed")
	}

	return ed25519.NewKeyFromSeed(s), nil
}

func NewAddressByPublicKey(publicKey ed25519.PublicKey) string {

	if len(publicKey) != PUBLIC_KEY_SIZE {
		return ""
	}
	k := make([]byte, PUBLIC_KEY_SIZE+1)
	copy(k[1:], publicKey)
	hash := blake2b.New256()
	hash.Write(k)
	h := hash.Sum(nil)
	address := "0x" + hex.EncodeToString(h)[0:64]
	return address
}

func ImportPrivatekey(pv string) (ed25519.PrivateKey, error) {
	_, bdpv, err := bech32.Decode(pv)
	if err != nil {
		return nil, err
	}

	pvv, err := bech32.ConvertBits(bdpv, 5, 8, false)
	if err != nil {
		return nil, err
	}

	return ed25519.NewKeyFromSeed(pvv[1:]), nil
	// return pvv[0:], nil
}

func ExportPrivatekey(pv ed25519.PrivateKey) (string, error) {
	pvv := append([]byte{0}, pv[:32]...)

	bcpv, err := bech32.ConvertBits(pvv, 8, 5, true)
	if err != nil {
		fmt.Println(err)
		return EmptyStr, err
	}
	return bech32.Encode(HRP, bcpv)
}

func Gen(prefix, suffix string) (string, string) {
	addresslen := 66
	for i := 0; i < 999999; i++ {
		pk, pv, _ := GenerateKey()
		addr := NewAddressByPublicKey(pk)
		if addr[(addresslen-len(suffix)):] == suffix && addr[2:(len(prefix))+2] == prefix {
			privateKey, err := ExportPrivatekey(pv)
			if err != nil {
				panic(err)
			}
			return addr, privateKey
		}
	}

	return "", ""
}
