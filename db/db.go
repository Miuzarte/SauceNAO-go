package db

import "reflect"

type IndexId int

const (
	HMAGAZINES     IndexId = 0  // H-Magazines
	HGAMECG        IndexId = 2  // H-Game CG
	DOUJINSHIDB    IndexId = 3  // DoujinshiDB
	PIXIV          IndexId = 5  // pixiv Images
	SEIGA          IndexId = 8  // Nico Nico Seiga
	DANBOORU       IndexId = 9  // Danbooru
	DRAWR          IndexId = 10 // drawr Images
	NIJIE          IndexId = 11 // Nijie Images
	YANDERE        IndexId = 12 // Yande.re
	SHUTTERSTOCK   IndexId = 15 // Shutterstock
	FAKKU          IndexId = 16 // FAKKU
	NHENTAI        IndexId = 18 // H-Misc (nH)
	MARKET2D       IndexId = 19 // 2D-Market
	MEDIBANG       IndexId = 20 // MediBang
	ANIME          IndexId = 21 // Anime
	HANIME         IndexId = 22 // H-Anime
	MOVIES         IndexId = 23 // Movies
	SHOWS          IndexId = 24 // Shows
	GELBOORU       IndexId = 25 // Gelbooru
	KONACHAN       IndexId = 26 // Konachan
	SANKAKU        IndexId = 27 // Sankaku Channel
	ANIMEPICTURES  IndexId = 28 // Anime-Pictures.net
	E621           IndexId = 29 // e621.net
	IDOLCOMPLEX    IndexId = 30 // Idol Complex
	BCY_ILLUST     IndexId = 31 // bcy.net Illust
	BCY_COSPLAY    IndexId = 32 // bcy.net Cosplay
	PORTALGRAPHICS IndexId = 33 // PortalGraphics.net
	DEVIANTART     IndexId = 34 // deviantArt
	PAWOO          IndexId = 35 // Pawoo.net
	MADOKAMI       IndexId = 36 // Madokami (Manga)
	MANGADEX       IndexId = 37 // MangaDex
	EHENTAI        IndexId = 38 // H-Misc (eH)
	ARTSTATION     IndexId = 39 // ArtStation
	FURAFFINITY    IndexId = 40 // FurAffinity
	TWITTER        IndexId = 41 // Twitter
	FURRYNETWORK   IndexId = 42 // Furry Network
	KEMONO         IndexId = 43 // Kemono
	SKEB           IndexId = 44 // Skeb
	ALL            IndexId = 999
)

func (di IndexId) String() string {
	if di >= 0 && int(di) < len(dbIdToName) {
		return dbIdToName[di]
	}
	if di == ALL {
		return "All DBs"
	}
	return "Unknown DB"
}

var dbIdToName = [...]string{
	HMAGAZINES:     "H-Magazines",
	HGAMECG:        "H-Game CG",
	DOUJINSHIDB:    "DoujinshiDB",
	PIXIV:          "pixiv Images",
	SEIGA:          "Nico Nico Seiga",
	DANBOORU:       "Danbooru",
	DRAWR:          "drawr Images",
	NIJIE:          "Nijie Images",
	YANDERE:        "Yande.re",
	SHUTTERSTOCK:   "Shutterstock",
	FAKKU:          "FAKKU",
	NHENTAI:        "H-Misc (nH)",
	MARKET2D:       "2D-Market",
	MEDIBANG:       "MediBang",
	ANIME:          "Anime",
	HANIME:         "H-Anime",
	MOVIES:         "Movies",
	SHOWS:          "Shows",
	GELBOORU:       "Gelbooru",
	KONACHAN:       "Konachan",
	SANKAKU:        "Sankaku Channel",
	ANIMEPICTURES:  "Anime-Pictures.net",
	E621:           "e621.net",
	IDOLCOMPLEX:    "Idol Complex",
	BCY_ILLUST:     "bcy.net Illust",
	BCY_COSPLAY:    "bcy.net Cosplay",
	PORTALGRAPHICS: "PortalGraphics.net",
	DEVIANTART:     "deviantArt",
	PAWOO:          "Pawoo.net",
	MADOKAMI:       "Madokami (Manga)",
	MANGADEX:       "MangaDex",
	EHENTAI:        "H-Misc (eH)",
	ARTSTATION:     "ArtStation",
	FURAFFINITY:    "FurAffinity",
	TWITTER:        "Twitter",
	FURRYNETWORK:   "Furry Network",
	KEMONO:         "Kemono",
	SKEB:           "Skeb",
}

var dbIdToType = [...]reflect.Type{
	HMAGAZINES:     reflect.TypeFor[*ResultDataHMagazines](),
	HGAMECG:        reflect.TypeFor[*ResultDataHGameCg](),
	DOUJINSHIDB:    reflect.TypeFor[*ResultDataDoujinshiDb](),
	PIXIV:          reflect.TypeFor[*ResultDataPixiv](),
	SEIGA:          reflect.TypeFor[*ResultDataSeiga](),
	DANBOORU:       reflect.TypeFor[*ResultDataDanbooru](),
	DRAWR:          reflect.TypeFor[*ResultDataDrawr](),
	NIJIE:          reflect.TypeFor[*ResultDataNijie](),
	YANDERE:        reflect.TypeFor[*ResultDataYandere](),
	SHUTTERSTOCK:   reflect.TypeFor[*ResultDataShutterstock](),
	FAKKU:          reflect.TypeFor[*ResultDataFakku](),
	NHENTAI:        reflect.TypeFor[*ResultDataNHentai](),
	MARKET2D:       reflect.TypeFor[*ResultDataMarket2d](),
	MEDIBANG:       reflect.TypeFor[*ResultDataMediBang](),
	ANIME:          reflect.TypeFor[*ResultDataAnime](),
	HANIME:         reflect.TypeFor[*ResultDataHAnime](),
	MOVIES:         reflect.TypeFor[*ResultDataMovies](),
	SHOWS:          reflect.TypeFor[*ResultDataShows](),
	GELBOORU:       reflect.TypeFor[*ResultDataGelbooru](),
	KONACHAN:       reflect.TypeFor[*ResultDataKonachan](),
	SANKAKU:        reflect.TypeFor[*ResultDataSankaku](),
	ANIMEPICTURES:  reflect.TypeFor[*ResultDataAnimePictures](),
	E621:           reflect.TypeFor[*ResultDataE621](),
	IDOLCOMPLEX:    reflect.TypeFor[*ResultDataIdolComplex](),
	BCY_ILLUST:     reflect.TypeFor[*ResultDataBcyIllust](),
	BCY_COSPLAY:    reflect.TypeFor[*ResultDataBcyCosplay](),
	PORTALGRAPHICS: reflect.TypeFor[*ResultDataPortalGraphics](),
	DEVIANTART:     reflect.TypeFor[*ResultDataDeviantArt](),
	PAWOO:          reflect.TypeFor[*ResultDataPawoo](),
	MADOKAMI:       reflect.TypeFor[*ResultDataMadokami](),
	MANGADEX:       reflect.TypeFor[*ResultDataMangaDex](),
	EHENTAI:        reflect.TypeFor[*ResultDataEHentai](),
	ARTSTATION:     reflect.TypeFor[*ResultDataArtStation](),
	FURAFFINITY:    reflect.TypeFor[*ResultDataFurAffinity](),
	TWITTER:        reflect.TypeFor[*ResultDataTwitter](),
	FURRYNETWORK:   reflect.TypeFor[*ResultDataFurryNetwork](),
	KEMONO:         reflect.TypeFor[*ResultDataKemono](),
	SKEB:           reflect.TypeFor[*ResultDataSkeb](),
}
