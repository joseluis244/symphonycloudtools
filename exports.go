package symphonycloudtools

import (
	"joseluis244/symphonycloudtools/r2"
)

var R2 *r2.R2
var dev bool = true

func Init(license string) {
	_, r2key, _, err := decryptLicense(license, "MedicareSoft203$")
	if err != nil {
		panic(err)
	}
	r2key["Dev"] = dev
	R2 = r2.Init(r2key)
}
