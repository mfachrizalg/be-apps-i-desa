package dtos

type FamilyCardsResponse struct {
	FamilyCards []FamilyCardResponse `json:"family_cards"`
}

type FamilyCardResponse struct {
	NIK          string `json:"nik"`
	Name         string `json:"name"`
	TotalMembers int    `json:"total_members"`
	Address      string `json:"address"`
}
