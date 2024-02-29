package stowiki

type STOWiki struct {
	Version        int64       `json:"version,omitempty"`
	StarshipTraits []Trait     `json:"starship_traits,omitempty"`
	PersonalTraits []Trait     `json:"personal_traits,omitempty"`
	Equipment      []Equipment `json:"equipment,omitempty"`
}

type Trait struct {
	Name             string   `json:"name,omitempty"`
	Type             string   `json:"type,omitempty"`
	Obtained         []string `json:"obtained,omitempty"`
	Desc             string   `json:"desc,omitempty"`
	Image            string   `json:"image,omitempty"`
	Environment      string   `json:"environment,omitempty,omitempty"`
	DisplayType      string   `json:"display_type,omitempty"`
	Availability     string   `json:"availability,omitempty,omitempty"`
	AvailabilityType string   `json:"availability_type,omitempty,omitempty"`
}

type Equipment struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Type        string `json:"type"`
	DisplayType string `json:"display_type"`
	Desc        Desc   `json:"desc"`
	Desc2       string `json:"desc2"`
	Rarity      string `json:"rarity"`
	Image       string `json:"image"`
}

type Desc struct {
	Head    Head `json:"head"`
	Subhead Head `json:"subhead"`
	Text    Text `json:"text"`
}

type Head struct {
	Num1 string `json:"1"`
	Num2 string `json:"2"`
	Num3 string `json:"3"`
	Num4 string `json:"4"`
	Num5 string `json:"5"`
	Num6 string `json:"6"`
	Num7 string `json:"7"`
	Num8 string `json:"8"`
	Num9 string `json:"9"`
}

type Text struct {
	Num1 string      `json:"1"`
	Num2 string      `json:"2"`
	Num3 interface{} `json:"3"`
	Num4 interface{} `json:"4"`
	Num5 interface{} `json:"5"`
	Num6 interface{} `json:"6"`
	Num7 interface{} `json:"7"`
	Num8 interface{} `json:"8"`
	Num9 interface{} `json:"9"`
}
