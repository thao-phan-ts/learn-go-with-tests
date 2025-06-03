package ascore

// ThongTinTongHopDuNo - BC300
type ThongTinTongHopDuNo struct {
	TTDUNOHT []*ThongTinDuNoHienTai      `json:"BC310"`
	LSQHTD   []*LopThongTinQuanHeTinDung `json:"BC320"`
}

// ThongTinDuNoHienTai - BC310
type ThongTinDuNoHienTai struct {
	ChiTietNoVay []*ChiTietNoVay `json:"BC3101,omitempty"`
	TTDUNOHT     []*TTDUNOHT     `json:"BC3102,omitempty"`
	NoVAMC       []*NoVAMC       `json:"BC3103,omitempty"`
}

// ChiTietNoVay - BC3101
type ChiTietNoVay struct {
	MaTCTD      string         `json:"TTC01"`
	TenTCTD     string         `json:"TTC02"`
	NgaySL      string         `json:"BC31011"`
	ChiTietDuNo []*ChiTietDuNo `json:"BC31012,omitempty"`
}

// ChiTietDuNo - BC31012
type ChiTietDuNo struct {
	MaLoaiVay     string         `json:"BC310121,omitempty"`
	ChiTietNhomNo *ChiTietNhomNo `json:"BC310122,omitempty"`
}

// ChiTietNhomNo - BC310122
type ChiTietNhomNo struct {
	MaNhomNo string `json:"BC3101221"`
	DuNoVND  string `json:"BC3101222"`
	DuNoUSD  string `json:"BC3101223"`
}

// TTDUNOHT - BC3102
type TTDUNOHT struct {
	MaTCTD           string            `json:"TTC01"`
	TenTCTDPhatHanh  string            `json:"TTC02"`
	DuNoTheTDChiTiet *DuNoTheTDChiTiet `json:"BC3101221,omitempty"`
}

// DuNoTheTDChiTiet - BC3101221
type DuNoTheTDChiTiet struct {
	TenThe           string `json:"TTC01"`
	TenTCTDPhatHanh  string `json:"TTC02"`
	DuNoTheTDChiTiet string `json:"BC31021"`
}

// LopThongTinQuanHeTinDung - BC320
type LopThongTinQuanHeTinDung struct {
	DuNo12Thang              []*DuNo12Thang              `json:"BC3201"`
	NoXau60Thang             []*NoXau60Thang             `json:"BC3202"`
	LichSuChamTTTheTD36Thang []*LichSuChamTTTheTD36Thang `json:"BC3203"`
	NoCanChuYTrong12Thang    []*NoCanChuYTrong12Thang    `json:"BC3204"`
}

// DuNo12Thang - BC3201
type DuNo12Thang struct {
	Thang     string `json:"BC32011"`
	DuNoVay   string `json:"BC32012"`
	DuNoThe   string `json:"BC32013"`
	HanMucThe string `json:"BC32014"`
}

// NoXau60Thang - BC3202
type NoXau60Thang struct {
	MaTCTD                    string `json:"TTC01"`
	TenTCTD                   string `json:"TTC02"`
	NgayPhatSinhNoXauCuoiCung string `json:"BC32021"`
	NhomNo                    string `json:"BC32022"`
	SoTienVND                 string `json:"BC32023"`
	SoTienUSD                 string `json:"BC32024"`
}

// LichSuChamTTTheTD36Thang - BC3203
type LichSuChamTTTheTD36Thang struct {
	MaTCTD   string                 `json:"TTC01"`
	TenTCTD  string                 `json:"TTC02"`
	TTCHAMTT *ThongTinChamThanhToan `json:"BC32031"`
}

// ThongTinChamThanhToan - BC32031
type ThongTinChamThanhToan struct {
	NgayChamThanhToan   string `json:"BC320311"`
	SoTienChamThanhToan string `json:"BC320312"`
	SoNgayChamThanhToan string `json:"BC320313"`
}

// NoCanChuYTrong12Thang - BC3204
type NoCanChuYTrong12Thang struct {
	TCTD     []*ToChucTinDung `json:"BC32041"`
	Thang    string           `json:"BC32042"`
	TongDuNo string           `json:"BC32043"`
}

// ToChucTinDung - BC32041
type ToChucTinDung struct {
	MaTCTD  string `json:"TTC01"`
	TenTCTD string `json:"TTC02"`
	NgaySL  string `json:"BC320411"`
}

// NoVAMC - BC3103
type NoVAMC struct {
	MaTCTD      string `json:"TTC01"`
	TenTCTD     string `json:"TTC02"`
	NoGocConLai string `json:"BC31031"`
	NgaySL      string `json:"BC31032"`
}
