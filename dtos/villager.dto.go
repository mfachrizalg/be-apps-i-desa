package dtos

type AddVillagerRequest struct {
	NIK              string  `json:"nik" validate:"required"`
	NamaLengkap      string  `json:"nama_lengkap" validate:"required"`
	JenisKelamin     string  `json:"jenis_kelamin" validate:"required"`
	TempatLahir      string  `json:"tempat_lahir" validate:"required"`
	TanggalLahir     string  `json:"tanggal_lahir" validate:"required"`
	Agama            string  `json:"agama" validate:"required"`
	Pendidikan       string  `json:"pendidikan" validate:"required"`
	Pekerjaan        string  `json:"pekerjaan" validate:"required"`
	StatusPerkawinan string  `json:"status_perkawinan" validate:"required"`
	StatusHubungan   string  `json:"status_hubungan" validate:"required"`
	Kewarganegaraan  string  `json:"kewarganegaraan" validate:"required"`
	NomorPaspor      *string `json:"nomor_paspor" validate:"omitempty"`
	NomorKitas       *string `json:"nomor_kitas" validate:"omitempty"`
	NamaAyah         string  `json:"nama_ayah" validate:"required"`
	NamaIbu          string  `json:"nama_ibu" validate:"required"`
	VillageID        string  `json:"village_id" validate:"required"`
	FamilyCardID     string  `json:"family_card_id" validate:"required"`
}

type GetFamilyMember struct {
	Name           string `json:"name"`
	StatusHubungan string `json:"status_hubungan"`
	Age            int    `json:"age"`
	JenisKelamin   string `json:"jenis_kelamin"`
	Pendidikan     string `json:"pendidikan"`
	Pekerjaan      string `json:"pekerjaan"`
}
