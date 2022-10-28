package sqlmodels

type DBProduct struct {
	ID            string
	RestaurantID  string
	Category      string
	Name          string
	Description   string
	Image         string
	PriceValue    int
	PriceCurrency string
}
