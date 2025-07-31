package models

import "github.com/google/uuid"

type SubDimensiPendidikan struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID          uuid.UUID `gorm:"type:uuid;not null"`
	Year               *int
	KetersediaanPaud   string
	KemudahanAksesPaud string
	ApmPaud            string
	KetersediaanSd     string
	KemudahanAksesSd   string
	ApmSd              string
	KetersediaanSmp    string
	KemudahanAksesSmp  string
	ApmSmp             string
	KetersediaanSma    string
	KemudahanAksesSma  string
	ApmSma             string
	Village            Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiKesehatan struct {
	ID                                         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                                  uuid.UUID `gorm:"type:uuid;not null"`
	Year                                       *int
	KemudahanAksesSaranaKesehatan              string
	KetersediaanSaranaKesehatan                string
	KemudahanAksesFasilitasKesehatan           string
	KetersediaanPosyandu                       string
	JumlahAktivitasPosyandu                    string
	KemudahanAksesPosyandu                     string
	KetersediaanLayananDokter                  string
	HariOperasionalLayananDokter               string
	PenyediaLayananDokter                      string
	PenyediaTransportasiLayananDokter          string
	KetersediaanLayananBidan                   string
	HariOperasionalLayananBidan                string
	PenyediaLayananBidan                       string
	PenyediaTransportasiLayananBidan           string
	KetersediaanLayananTenagaKesehatan         string
	HariOperasionalLayananTenagaKesehatan      string
	PenyediaLayananTenagaKesehatan             string
	PenyediaTransportasiLayananTenagaKesehatan string
	PersentasePesertaJaminanKesehatan          string
	KegiatanSosialisasiJaminanKesehatan        string
	Village                                    Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiUtilitasDasar struct {
	ID                            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                     uuid.UUID `gorm:"type:uuid;not null"`
	Year                          *int
	OperasionalAirMinum           string
	KetersediaanAirMinum          string
	KemudahanAksesAirMinum        string
	KualitasAirMinum              string
	PersentaseRumahTidakLayakHuni string
	Village                       Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiAktivitas struct {
	ID                                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                         uuid.UUID `gorm:"type:uuid;not null"`
	Year                              *int
	KearifanSosial                    string
	KearifanSosialDipertahankan       string
	KegiatanGotongRoyong              string
	FrekuensiGotongRoyong             string
	KeterlibatanWargaGotongRoyong     string
	FrekuensiKegiatanOlahraga         string
	PenyelesaianKonflikSecaraDamai    string
	PeranAparatKeamananMediator       string
	PeranAparatPemerintah             string
	PeranTokohMasyarakat              string
	PeranTokohAgama                   string
	SatuanKeamananLingkungan          string
	AktivitasSatuanKeamananLingkungan string
	Village                           Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiFasilitasMasyarakat struct {
	ID                                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                            uuid.UUID `gorm:"type:uuid;not null"`
	Year                                 *int
	TerdapatTamanBacaanMasyarakat        string
	HariOperasionalTamanBacaanMasyarakat string
	KetersediaanFasilitasOlahraga        string
	KeberadaanRuangPublikTerbuka         string
	Village                              Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiProduksiDesa struct {
	ID                                       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                                uuid.UUID `gorm:"type:uuid;not null"`
	Year                                     *int
	KeragamanAktivitasEkonomi                string
	KeaktifanAktivitasEkonomi                string
	KetersediaanProdukUnggulanDesa           string
	CakupanPasarProdukUnggulan               string
	KetersediaanMerekDagang                  string
	TerdapatKearifanLokalEkonomi             string
	TelahDilakukanKerjaSamaDenganDesaLainnya string
	TelahDilakukanKerjaSamaDenganPihakKetiga string
	Village                                  Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiFasilitasPendukungEkonomi struct {
	ID                                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                         uuid.UUID `gorm:"type:uuid;not null"`
	Year                              *int
	KetersediaanPendidikanNonFormal   string
	KeterlibatanPendidikanNonFormal   string
	KetersediaanPasarRakyat           string
	KemudahanAksesPasarRakyat         string
	KetersediaanToko                  string
	KemudahanAksesToko                string
	KetersediaanRumahMakan            string
	KemudahanAksesRumahMakan          string
	KetersediaanPenginapan            string
	KemudahanAksesPenginapan          string
	KetersediaanLogistik              string
	KemudahanAksesLogistik            string
	TerdapatBumd                      string
	BumdBerbadanHukum                 string
	HariOperasionalLembagaEkonomi     string
	KetersediaanLembagaEkonomiLainnya string
	KetersediaanKud                   string
	KetersediaanUmkm                  string
	LayananPerbankan                  string
	HariOperasionalKeuangan           string
	LayananFasilitasKreditKur         string
	LayananFasilitasKreditKkpE        string
	LayananFasilitasKreditKuk         string
	StatusLayananFasilitasKredit      string
	Village                           Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiPengelolaanLingkungan struct {
	ID                                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                         uuid.UUID `gorm:"type:uuid;not null"`
	Year                              *int
	UpayaMenjagaKelestarianLingkungan string
	RegulasiPelestarianLingkungan     string
	KegiatanPelestarianLingkungan     string
	PemanfaatanEnergiTerbarukan       string
	TempatPembuananganSampah          string
	PengelolaanSampah                 string
	PemanfaatanSampah                 string
	KejadianPencemaranLingkungan      string
	KetersediaanJamban                string
	KeberfungsianJamban               string
	KetersediaanSepticTank            string
	PembuanganAirLimbahCairRumah      string
	Village                           Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiPenanggulanganBencana struct {
	ID                                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                           uuid.UUID `gorm:"type:uuid;not null"`
	Year                                *int
	AspekInformasiKebencanaan           string
	FasilitasMitigasiBencana            string
	AksesMenujuFasilitasMitigasiBencana string
	AktivitasMitigasi                   string
	FasilitasTanggapDaruratBencana      string
	Village                             Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiKondisiAksesJalan struct {
	ID                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID            uuid.UUID `gorm:"type:uuid;not null"`
	Year                 *int
	JenisPermukaanJalan  string
	KualitasJalan        string
	PeneranganJalanUtama string
	OperasionalPju       string
	Village              Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiKemudahanAkses struct {
	ID                           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                    uuid.UUID `gorm:"type:uuid;not null"`
	Year                         *int
	AngkutanPerdesaan            string
	OperasionalAngkutanPerdesaan string
	PelayananListrik             string
	DurasiLayananListrik         string
	AksesTelepon                 string
	AksesInternet                string
	Village                      Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiKelembagaanPelayananDesa struct {
	ID                                     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID                              uuid.UUID `gorm:"type:uuid;not null"`
	Year                                   *int
	LayananDiberikan                       string
	PublikasiInformasiPelayanan            string
	PelayananAdministrasi                  string
	PelayananPengaduan                     string
	PelayananLainnya                       string
	MusyawarahDesa                         string
	MusyawarahDesaDidatangiUnsurMasyarakat string
	Village                                Village `gorm:"foreignKey:VillageID"`
}

type SubDimensiTataKelolaKeuanganDesa struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VillageID             uuid.UUID `gorm:"type:uuid;not null"`
	Year                  *int
	PendapatanAsliDesa    string
	PeningkatanPades      string
	PenyertaanModalDdBumd string
	AsetTanahDesa         string
	AsetKantorDesa        string
	AsetPasarDesa         string
	AsetLainnya           string
	ProduktivitasAsetDesa string
	InventarisasiAsetDesa string
	Village               Village `gorm:"foreignKey:VillageID"`
}
