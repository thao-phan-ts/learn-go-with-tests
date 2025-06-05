package ascore

type CICS11T struct {
	MaSoPhieu         string                  `json:"MASOPHIEU"`
	TenSP             string                  `json:"TL001"`
	MaSoSP            string                  `json:"TL002"`
	DonViTraCuu       string                  `json:"TL003"`
	DiaChi            string                  `json:"TL004"`
	SDTNguoiTraCuu    string                  `json:"TL005"`
	ThoiGianYeuCau    string                  `json:"TL006"`
	ThoiGianGuiBaoCao string                  `json:"TL007"`
	TenNguoiTraCuu    string                  `json:"TL008"`
	NguoiTraCuu       string                  `json:"TL009"`
	TenKH             string                  `json:"TENKH"`
	KQTTCIC           string                  `json:"TL099"`
	CCCD              string                  `json:"CCCD"`
	GhiChuKhac        string                  `json:"TL100,omitempty"`
	CoDuNoVDB         string                  `json:"TL101,omitempty"`
	TTKH              *ThongTinKhachHang      `json:"BC100,omitempty"`
	TTDiemTinDung     *ThongTinDiemTinDung    `json:"BC200,omitempty"`
	TTTongHopDuNo     *ThongTinTongHopDuNo    `json:"BC300,omitempty"`
	TTTSDB            *ThongTinTaiSanDamBao   `json:"BC400,omitempty"`
	TTTongHopKHAC     *ThongTinTongHopKhac    `json:"BC500,omitempty"`
	TTKhacVeKHVay     *ThongTinVeKhachHangVay `json:"BC600,omitempty"`
}
