package nova

import "github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"

var skins = GetCustomSkins(fileIO.ReadDirectory("../custom_skins"))
