package nova

import "github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"

var Skins = GetCustomSkins(fileIO.ReadDirectory("../custom_skins"))
