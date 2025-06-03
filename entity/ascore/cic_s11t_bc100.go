package ascore

// ThongTinKhachHang - BC100
type ThongTinKhachHang struct {
	TenKhachHang string `json:"TTC04"`
	MaCIC        string `json:"MACIC"`
	DiaChiKH     string `json:"CN003"`
	CCCD         string `json:"CN008"`
	TTXacThuc    bool   `json:"CN100"`
	GiayToKhac   string `json:"CN101"`
}
