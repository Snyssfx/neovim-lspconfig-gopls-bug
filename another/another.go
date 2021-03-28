package another

import (
	main_package "github.com/Snyssfx/neovim-lspconfig-gopls-bug"
)

func AnotherPrint() {
	println("in another package")
	main_package.Print()
}
