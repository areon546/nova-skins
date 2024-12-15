package nova

import (
	"github.com/areon546/go-helpers"
)

// helper functions

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func pagesFolder() string {
	return "../pages"

	// reads what files are in the assets folder
	// assets := readDirectory(skinFolder())
	// assetsAsFiles, _ := filterFiles(assets)

	// print("assets", assets)
	// printf("%s", "abba")
	// for _, v := range assetsAsFiles {
	// 	print(v.String())
	// }

}

func skinFolder() string {
	return "../custom_skins/"
}
