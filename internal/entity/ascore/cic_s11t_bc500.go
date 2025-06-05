package ascore

// ThongTinTongHopKhac - BC500
type ThongTinTongHopKhac struct {
	HDTD              []*HopDongTinDung                `json:"BC510,omitempty"`
	TCTDTraCuu12Thang []*ThongTinhTinDungTraCuu12Thang `json:"BC520,omitempty"`
}

// HopDongTinDung - BC510
type HopDongTinDung struct {
	MaTCTD  string `json:"TTC01"`
	TenTCTD string `json:"TTC02"`
	SoHD    string `json:"BC5101"`
	NgayKy  string `json:"BC5102"`
	NgayKT  string `json:"BC5103"`
}

type ThongTinhTinDungTraCuu12Thang struct {
	MaTCTD     string `json:"TTC01"`
	TenTCTD    string `json:"TTC02"`
	MaSanPham  string `json:"BC5201"`
	NgayTraCuu string `json:"BC5202"`
}
