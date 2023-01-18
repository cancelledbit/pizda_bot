package stickers

import (
	"errors"
	"math/rand"
	"regexp"
	"time"
)

type Sticker struct {
	Name string
	ID   string
}

type StickerList []*Sticker

func (s StickerList) getRandom() *Sticker {
	rand.Seed(time.Now().Unix())
	if len(s) == 0 {
		return nil
	}
	return s[rand.Intn(len(s))]
}

type rulesetMap map[string]*StickerList

func getRuleset() rulesetMap {
	return rulesetMap{
		"(\\s+|^)[Дд][АаAa][!.?]{0,3}$": &StickerList{
			{ID: "CAACAgIAAxkBAAEHOZhjwEsfquuRaCqEBU_U182GOZYuCgACHCAAAjSc8UnGh3clJRWcTC0E", Name: "Pizda"},
			{ID: "CAACAgIAAxkBAAEHUqhjx_U6-jJsimqM31XiMaS1bRarQgAC5QAD5EXyIhs2qo0Jf12kLQQ", Name: "Pizda"},
		},
		"(\\s+|^)[Нн][ЕеEe][Тт][!.?]{0,3}$": &StickerList{
			{ID: "CAACAgIAAxkBAAEHOZpjwEtGZ4pNgR6APnYa97akzvooGgAClyEAAsn48EkRMGHyM7t0sy0E", Name: "Minet"},
			{ID: "CAACAgIAAxkBAAEHUqJjx_Sm-pLiIUutMXQeVMfb1mU8kQACEiIAAjzZ-Ukiy3xs84cpCS0E", Name: "PidoraOtvet"},
		},
		"(\\s+|^)[Яя][!.?]{0,3}$": &StickerList{
			{ID: "CAACAgIAAxkBAAEHOZxjwEtywoFWsE69uoflXkAHLLI7wAACkB8AAjEQ8UnAWWWzeHqdmi0E", Name: "Golovka"},
			{ID: "CAACAgQAAxkBAAEHUrdjx_W8G7yzd0AvOuoMSIkdLrms3wACjgEAAjPl4hUvI3VlUsQgtC0E", Name: "Golovka"},
		},
		"(\\s+|^)[АаAa][Гг][АаAa][!.?]{0,3}$": &StickerList{
			{ID: "CAACAgIAAxkBAAEHOZ5jwEuOAAG3rYZyO5iAV_dPZBnK4nsAAuUlAAIT-PlJxOxD5v5rRUItBA", Name: "Noga"},
			{ID: "CAACAgIAAxkBAAEHUq9jx_VlnRdpGTiN9hvJmuQWFGpl1QAC6gAD5EXyIhtzC4in6JjjLQQ", Name: "Noga"},
		},
		"(\\s+|^)[Чч]([ОоOo0]|[Ее]|[Ёё])[!.?]{0,3}$": &StickerList{
			{ID: "CAACAgIAAxkBAAEHOaBjwEujfGPeM4Yu1pq6Gdncp2l-VQACESUAAlGP8Endvec9raBNri0E", Name: "Plecho"},
		},
		"(\\s+|^)[Мм][Нн][ЕеEe][!.?]{0,3}$": &StickerList{
			{ID: "CAACAgIAAxkBAAEHOaJjwEvAdmTt5FmVHl3PLRXbshur7QACvSIAApjG8Un83RlK-rjfqy0E", Name: "Vgovne"},
		},
	}
}

func GetStickerBy(msg string) (*Sticker, error) {
	for r, sticker := range getRuleset() {
		if rgx, err := regexp.Compile(r); err == nil {
			if rgx.MatchString(msg) {
				return sticker.getRandom(), nil
			}
		}
	}
	return nil, errors.New("no sticker found")
}
