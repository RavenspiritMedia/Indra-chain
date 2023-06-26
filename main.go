package main

import (
	_ "embed"

	"github.com/RavenspiritMedia/Indra-chain/command/root"
	"github.com/RavenspiritMedia/Indra-chain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
