package config

import "github.com/evalphobia/go-config-loader"

const (
	confType     = "toml"
	basePath     = "."
	discordToken = "discord.token"
	soundsPath   = "paths.sounds"
	jsonDbPath   = "paths.db"
)

type Values struct {
	DiscordToken string
	JsonDbPath   string
	SoundsPath   string
}

func LoadConfigs() Values {
	var conf *config.Config
	conf = config.NewConfig()
	if err := conf.LoadConfigs(basePath, confType); err != nil {
		panic(err)
	}
	return Values{
		DiscordToken: conf.ValueString(discordToken),
		SoundsPath:   conf.ValueString(soundsPath),
		JsonDbPath:   conf.ValueString(jsonDbPath),
	}
}
