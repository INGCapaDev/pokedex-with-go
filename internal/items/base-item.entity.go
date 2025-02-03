package items

type BaseItem struct {
	ItemBaseInfo
	Quantity  int     `json:"quantity"`
	SellPrice float64 `json:"sell_price"`
	CanBeSold bool    `json:"can_be_sold"`
}

func (i *BaseItem) GetBaseInfo() ItemBaseInfo {
	return i.ItemBaseInfo
}

func (i *BaseItem) GetQuantity() int {
	return i.Quantity
}

func (i *BaseItem) GetSellPrice() float64 {
	return i.SellPrice
}

func (i *BaseItem) GetCanBeSold() bool {
	return i.CanBeSold
}

func (i *BaseItem) SetQuantity(quantity int) {
	i.Quantity = quantity
}
