package another

import (
	main_package "github.com/Snyssfx/neovim-lspconfig-gopls-bug"
)

func AnotherPrint() {
	print("in another package")
	main_package.Print()
}
