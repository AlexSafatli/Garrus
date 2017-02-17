package championgg

import "github.com/AlexSafatli/DiscordSwissArmyKnife/rest"

type cggAPI struct {
	*rest.Client
	Champion *ChampionService
}

func NewCggAPI(appkey string) (cgg *cggAPI) {
	cgg = new(cggAPI)
	c := rest.NewClient("Discord", appkey, "http://api.champion.gg")
	cgg.Client = c
	cgg.Champion = (*ChampionService)(&cgg.Common)
	return
}
