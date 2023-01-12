package stickers

import (
	"errors"
	"regexp"
)

const Pizda = "CAACAgIAAxkBAAEHOZhjwEsfquuRaCqEBU_U182GOZYuCgACHCAAAjSc8UnGh3clJRWcTC0E"
const Minet = "CAACAgIAAxkBAAEHOZpjwEtGZ4pNgR6APnYa97akzvooGgAClyEAAsn48EkRMGHyM7t0sy0E"
const Golovka = "CAACAgIAAxkBAAEHOZxjwEtywoFWsE69uoflXkAHLLI7wAACkB8AAjEQ8UnAWWWzeHqdmi0E"
const Noga = "CAACAgIAAxkBAAEHOZ5jwEuOAAG3rYZyO5iAV_dPZBnK4nsAAuUlAAIT-PlJxOxD5v5rRUItBA"
const Plecho = "CAACAgIAAxkBAAEHOaBjwEujfGPeM4Yu1pq6Gdncp2l-VQACESUAAlGP8Endvec9raBNri0E"
const VGovne = "CAACAgIAAxkBAAEHOaJjwEvAdmTt5FmVHl3PLRXbshur7QACvSIAApjG8Un83RlK-rjfqy0E"

type Sticker struct {
	Name string
	ID   string
}

type rulesetMap map[string]*Sticker

func getRuleset() *rulesetMap {
	return &rulesetMap{
		"(\\s+|^)[Дд][АаAa][!.?]{0,3}$": &Sticker{
			ID:   Pizda,
			Name: "Pizda",
		},
		"(\\s+|^)[Нн][ЕеEe][Тт][!.?]{0,3}$": &Sticker{
			ID:   Minet,
			Name: "Minet",
		},
		"(\\s+|^)[Яя][!.?]{0,3}$": &Sticker{
			ID:   Golovka,
			Name: "Golovka",
		},
		"(\\s+|^)[АаAa][Гг][АаAa][!.?]{0,3}$": &Sticker{
			ID:   Noga,
			Name: "Noga",
		},
		"(\\s+|^)[Чч]([ОоOo0]|[Ее]|[Ёё])[!.?]{0,3}$": &Sticker{
			ID:   Plecho,
			Name: "Plecho",
		},
		"(\\s+|^)[Мм][Нн][ЕеEe][!.?]{0,3}$": &Sticker{
			ID:   VGovne,
			Name: "Vgovne",
		},
	}
}

func GetStickerBy(msg string) (*Sticker, error) {
	for r, sticker := range *getRuleset() {
		if rgx, err := regexp.Compile(r); err == nil {
			if rgx.MatchString(msg) {
				return sticker, nil
			}
		}
	}
	return nil, errors.New("no sticker found")
}
