package championgg

import (
	"net/http"

	"github.com/AlexSafatli/DiscordSwissArmyKnife/rest"
)

type ChampionService rest.Service

type ChampionInfo struct {
	Name  string         `json:"name"`
	Roles []ChampionRole `json:"roles"`
}

type ChampionRole struct {
	Key               string             `json:"key,omitempty"`
	NumGames          int64              `json:"games,omitempty"`
	PercentPlayed     int8               `json:"percentPlayed,omitempty"`
	Name              string             `json:"name,omitempty"`
	FinalBuilds       ChampionItems      `json:"items"`
	FirstItems        ChampionItems      `json:"firstItems"`
	Trinkets          []ChampionItemData `json:"trinkets"`
	Matrix            []float32          `json:"championMatrix"`
	Summoners         ChampionSummoners  `json:"summoners"`
	Runes             ChampionRunes      `json:"runes"`
	ExperienceRate    []float32          `json:"experienceRate"`
	ExperienceSample  []float32          `json:"experienceSample"`
	PatchWinRates     [6]float32         `json:"patchWin"`
	PatchPlayRates    [6]float32         `json:"patchPlay"`
	DamageComposition `json:"dmgComposition"`
	Matchups          []Matchup         `json:"matchups"`
	Stats             ChampionStats     `json:"general"`
	Skills            ChampionSkills    `json:"skills"`
	Masteries         ChampionMasteries `json:"masteries"`
}

type ChampionPosition struct {
	Position int64 `json:"position"`
	Change   int64 `json:"change"`
}

type ChampionItems struct {
	MostGames         ChampionItemData `json:"mostGames"`
	HighestWinPercent ChampionItemData `json:"highestWinPercent"`
}

type ChampionSummoners struct {
	MostGames         ChampionSummonerData `json:"mostGames"`
	HighestWinPercent ChampionSummonerData `json:"highestWinPercent"`
}

type ChampionRunes struct {
	MostGames         ChampionRunesData `json:"mostGames"`
	HighestWinPercent ChampionRunesData `json:"highestWinPercent"`
}

type ChampionSkills struct {
	SkillInfo         []Skill           `json:"skillInfo"`
	MostGames         ChampionSkillData `json:"mostGames"`
	HighestWinPercent ChampionSkillData `json:"highestWinPercent"`
}

type ChampionMasteries struct {
	MostGames         ChampionMasteriesData `json:"mostGames"`
	HighestWinPercent ChampionMasteriesData `json:"highestWinPercent"`
}

type ChampionItemData struct {
	Item       `json:"item,omitempty"`
	Items      []Item  `json:"items,omitempty"`
	FirstItems []Item  `json:"firstItems,omitempty"`
	WinPercent float32 `json:"winPercent"`
	Games      int64   `json:"games"`
}

type ChampionSummonerData struct {
	SummonerD  Summoner `json:"summoner1"`
	SummonerF  Summoner `json:"summoner2"`
	WinPercent float32  `json:"winPercent"`
	Games      int64    `json:"games"`
}

type ChampionRunesData struct {
	Runes      []Runes `json:"runes"`
	WinPercent float32 `json:"winPercent"`
	Games      int64   `json:"games"`
}

type ChampionSkillData struct {
	Order      []string `json:"order"`
	WinPercent float32  `json:"winPercent"`
	Games      int64    `json:"games"`
}

type ChampionMasteriesData struct {
	Masteries  []Masteries `json:"masteries"`
	WinPercent float32     `json:"winPercent"`
	Games      int64       `json:"games"`
}

type ChampionStats struct {
	WinRate                NumericalDelta `json:"winPercent"`
	PlayRate               NumericalDelta `json:"playPercent"`
	BanRate                NumericalDelta `json:"banRate"`
	Experience             NumericalDelta `json:"experience"`
	GoldEarned             NumericalDelta `json:"goldEarned"`
	Kills                  NumericalDelta `json:"kills"`
	Deaths                 NumericalDelta `json:"deaths"`
	Assists                NumericalDelta `json:"assists"`
	TotalDamageToChampions NumericalDelta `json:"totalDamageDealtToChampions"`
	LargestSpree           NumericalDelta `json:"largestKillingSpree"`
	MinionsKilled          NumericalDelta `json:"minionsKilled"`
}

type Masteries struct {
	Tree        string        `json:"tree"`
	TotalPoints int8          `json:"total"`
	Allocations []MasteryData `json:"data"`
}

type MasteryData struct {
	ID     int64 `json:"mastery"`
	Points int8  `json:"points"`
}

type Skill struct {
	Name   string `json:"name"`
	Hotkey string `json:"key"`
}

type NumericalDelta struct {
	Change   float32 `json:"change"`
	Position float32 `json:"position"`
}

type Matchup struct {
	Games         int64   `json:"number"`
	Score         float32 `json:"statScore"`
	WinRate       float32 `json:"winRate"`
	WinRateChange float32 `json:"winRateChange"`
	Opponent      string  `json:"key"`
}

type DamageComposition struct {
	TrueDamage     float32 `json:"trueDmg"`
	MagicDamage    float32 `json:"magicDmg"`
	PhysicalDamage float32 `json:"physicalDmg"`
}

type Runes struct {
	ID          int64  `json:"id"`
	Quantity    int8   `json:"number"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Summoner struct {
	Name string `json:"name"`
}

type Item struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Champions []ChampionInfo
type ChampionRoles []ChampionRole

func (c *ChampionService) List() (*Champions, *http.Response, error) {
	r, err := c.Client.NewRequest("GET", "champions", nil, nil)
	if err != nil {
		return nil, nil, err
	}
	champions := &Champions{}
	resp, err := c.Client.Do(r, champions)
	if err != nil {
		return nil, resp, err
	}
	return champions, resp, nil
}

func (c *ChampionService) Get(name string) (*ChampionRoles, *http.Response, error) {
	r, err := c.Client.NewRequest("GET", "champion/"+name, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	champion := &ChampionRoles{}
	resp, err := c.Client.Do(r, champion)
	if err != nil {
		return nil, resp, err
	}
	return champion, resp, nil
}
