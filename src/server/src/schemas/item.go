package schema

type AddItem struct {
	BrandName      string `json:"brand_name"`
	GarmentId      string `json:"garment_id"`
	SizeTypeId     string `json:"size_type_id"`
	SizeTypeItemId string `json:"size_type_item_id"`
}

type GetItem struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	BrandName      string `json:"brand_name"`
	GarmentId      string `json:"garment_id"`
	SizeTypeId     string `json:"size_type_id"`
	SizeTypeItemId string `json:"size_type_item_id"`
	ItemState      string `json:"item_state"`
}
