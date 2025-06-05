package ascore

// ThongTinTaiSanDamBao - BC400
type ThongTinTaiSanDamBao struct {
	TTChiTietTSDB *ThongTinChiTietTaiSanDamBao `json:"BC4001"`
}

// ThongTinChiTietTaiSanDamBao - BC4001
type ThongTinChiTietTaiSanDamBao struct {
	MaTCTD       string          `json:"TTC01"`
	TenTCTD      string          `json:"TTC02"`
	NgaySLDuNo   string          `json:"BC4101,omitempty"`
	NgaySLTaiSan string          `json:"BC4102,omitempty"`
	DanhSachTSDB []*DanhSachTSDB `json:"BC4103,omitempty"`
}

// DanhSachTSDB - BC4103
type DanhSachTSDB struct {
	MaTSDB       string `json:"BC41031"`
	LoaiTSDB     string `json:"BC41032"`
	ChuSoHuu     string `json:"BC41033"`
	GiaTriTSDB   string `json:"BC41034"`
	NgayTheChap  string `json:"BC41035"`
	NgayGiaiChap string `json:"BC41036"`
	MotaTSDB     string `json:"BC41037"`
}
