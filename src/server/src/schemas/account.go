package schema

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddItem struct {
	BrandName      string `json:"brand_name"`
	GarmentId      string `json:"garment_id"`
	SizeTypeId     string `json:"size_type_id"`
	SizeTypeItemId string `json:"size_type_item_id"`
	ItemState      string `json:"item_state"`
}

type GetItem struct {
	Id             string `json:"id"`
	BrandName      string `json:"brand_name"`
	GarmentId      string `json:"garment_id"`
	SizeTypeId     string `json:"size_type_id"`
	SizeTypeItemId string `json:"size_type_item_id"`
	ItemState      string `json:"item_state"`
	DtimeCreated   string `json:"dtime_created"`
	DtimeArchived  string `json:"dtime_archived"`
}

type DeleteItem struct {
	Id string `json:"id"`
}

type Token struct {
	TokenType   string `json:"token_type"`
	AuthToken   string `json:"auth_token"`
	GeneratedAt string `json:"generated_at"`
	ExpiresAt   string `json:"expires_at"`
}
