package reflectdemo

type Application struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	packageData []byte `json:"-"`
}
type MoneyType int

const (
	MoneyTypeUS = MoneyType(iota)
	MoneyTypeCN
)

type Money struct {
	MoneyType MoneyType `json:"money_type" xml:"MoneyType"`
	Number    float64   `json:"number"`
}

type MacApplication struct {
	Application `json:"application"`
	AppleStore  string `json:"apple_store"`
	Favorite    int    `json:"favorite"`
	Money       Money  `json:"money"`
	int
}

func (app Application) GetData() []byte {
	return app.packageData
}
