package stickers

import (
	"errors"
	"regexp"
)

const Pizda = "CAACAgQAAxkBAAEGp1tjijXtr2D5U6k8sOiH8u0Dm4hnbgACkwEAAjPl4hVchh0Om9EItisE"
const Minet = "CAACAgQAAxkBAAEGp11jijZCAnIzORNZ3toEYW7a3hLrjwACkQEAAjPl4hUAATXKzK6Do7IrBA"
const Golovka = "CAACAgQAAxkBAAEGp2FjijZsOO-4qlhah9KXbJjK4DZu7wACjgEAAjPl4hUvI3VlUsQgtCsE"
const Noga = "CAACAgQAAxkBAAEGqW9jiy53_U2rAq1zXnh_CWtAFMfgNQACkgEAAjPl4hWqAbtFfYNHTCsE"
const Plecho = "CAACAgQAAxkBAAEGqXFjiy6iVSF_Ef85Feh4TZvQl2NBswACjwEAAjPl4hU0OKNGHbQi-SsE"
const VGovne = "CAACAgQAAxkBAAEGqXNjiy7Rqv7xzKhY2f485s3R1Y02dgAClAEAAjPl4hWDe5gGCZpsOSsE"

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
