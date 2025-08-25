package dirs

func content() string {
	return "../www/content/"
}

func Pages() string {
	return content() + "pages/"
}

// Media

func media() string {
	return "../media/"
}

func Skins() string {
	return media() + "custom_skins/"
}

func Assets() string {
	return media() + "assets/"
}

// WWW Raw

func rawMedia() string {
	return "https://raw.githubusercontent.com/areon546/nova-skins/refs/heads/main/media/"
}

// This is the link used to show the resources on the website.
// Could alternatively use a github URL link.
func WwwSkins() string {
	return rawMedia() + "custom_skins/"
}

func WwwAssets() string {
	return rawMedia() + "assets/"
}
