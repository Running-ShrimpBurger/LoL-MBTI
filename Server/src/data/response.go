package data

type Answer struct {
	Line     string `json:"line"`
	Response []int  `json:"response"`
}

type Result struct {
	Name      string `json:"name"`
	GoodChamp string `json:"good_champ"`
	BadChamp  string `json:"bad_champ"`
}
