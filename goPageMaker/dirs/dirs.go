package dirs

func media() string {
	return "../media/"
}

func PagesFolder() string {
	return "../www/content/pages/"
}

func SkinsFolder() string {
	return media() + "custom_skins/"
}

// This is the link used to show the resources on the website.
// Could alternatively use a github URL link.
func WwwSkinsFolder() string {
	return "https://github.com/areon546/nova-skins/blob/main/media/custom_skins/"
}

func AssetsFolder() string {
	return media() + "assets/"
}

func WwwAssetsFolder() string {
	return "https://github.com/areon546/nova-skins/blob/main/media/assets/"
}
