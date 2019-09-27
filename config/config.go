package config

import "github.com/evalphobia/go-config-loader"

const (
	confType     = "toml"
	basePath     = "."
	discordToken = "discord.token"
)

type Values struct {
	DiscordToken string
}

func LoadConfigs() Values {
	var conf *config.Config
	conf = config.NewConfig()
	if err := conf.LoadConfigs(basePath, confType); err != nil {
		panic(err)
	}
	return Values{
		DiscordToken: conf.ValueString(discordToken),
	}
}
