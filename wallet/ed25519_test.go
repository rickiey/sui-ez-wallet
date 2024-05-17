package wallet

import (
	"fmt"
	"testing"

	"github.com/rickiey/btcutil/bech32"
	"github.com/stretchr/testify/assert"
)

func Test_gen(t *testing.T) {
	pk, pv, _ := GenerateKey()

	pvv := append([]byte{0}, pv[:32]...)

	fmt.Println(len(pvv))
	fmt.Println(pvv)

	bcpv, err := bech32.ConvertBits(pvv, 8, 5, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	hrp := "suiprivkey"

	expv, err := bech32.Encode(hrp, bcpv)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(NewAddressByPublicKey(pk))
	fmt.Println((expv))
}

func Test_ImExport(t *testing.T) {

	for i := 0; i < 99999; i++ {
		pk, pv, err := GenerateKey()
		assert.Nil(t, err, "GenerateKey")
		pvs, err := ExportPrivatekey(pv)
		assert.Nil(t, err, "ExportPrivatekey")

		pvv, err := ImportPrivatekey(pvs)
		assert.Nil(t, err, "ImportPrivatekey")

		assert.Equal(t, pvv, pv)
		assert.Equal(t, pvv.Public(), pk)
	}

}

func Test_genSubfix(t *testing.T) {

	addresslen := 66
	suffix := "8"
	prefix := ""
	for i := 0; i < 999999; i++ {
		pk, pv, _ := GenerateKey()
		addr := NewAddressByPublicKey(pk)
		if addr[(addresslen-len(suffix)):] == suffix && addr[2:(len(prefix))+2] == prefix {
			privateKey, err := ExportPrivatekey(pv)
			if err != nil {
				panic(err)
			}
			fmt.Println(addr, privateKey)
		}
	}

}
