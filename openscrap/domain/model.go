package domain

type Item struct {
	MtrlAccdKindCd    string `json:"mtrlAccdKindCd" gorm:"column:mtrl_accd_kind_cd"`
	AccdNo            string `json:"accdNo" gorm:"column:accd_no"`
	AccdNm            string `json:"accdNm" gorm:"column:accd_nm"`
	AccdDfsTypeCd     string `json:"accdDfsTypeCd" gorm:"column:accd_dfs_type_cd"`
	SafePlanRsnCd     string `json:"safePlanRsnCd" gorm:"column:safe_plan_rsn_cd"`
	DamageContent     string `json:"damageContent" gorm:"column:damage_content"`
	RcvInstCd         string `json:"rcvInstCd" gorm:"column:rcv_inst_cd"`
	MtrlAccdObjCd     string `json:"mtrlAccdObjCd" gorm:"column:mtrl_accd_obj_cd"`
	AccdSafeprotActYn string `json:"accdSafeprotActYn" gorm:"column:accd_safeprot_act_yn"`
	WcConstStartYmd   string `json:"wcConstStartYmd" gorm:"column:wc_const_start_ymd"`
	ReadCnt           int    `json:"readCnt" gorm:"column:read_cnt"`
	PreventMeasure    string `json:"preventMeasure" gorm:"column:prevent_measure"`
	AccdTypeCd        string `json:"accdTypeCd" gorm:"column:accd_type_cd"`
	WrtInstCd         string `json:"wrtInstCd" gorm:"column:wrt_inst_cd"`
	AddrType          string `json:"addrType" gorm:"column:addr_type"`
	AddrSido          string `json:"addrSido" gorm:"column:addr_sido"`
	FrgDeathCount     int    `json:"frgDeathCount" gorm:"column:frg_death_count"`
	ProStatusCd       string `json:"proStatusCd" gorm:"column:pro_status_cd"`
	AccdReasonType    string `json:"accdReasonType" gorm:"column:accd_reason_type"`
	HmnWorkClassCd    string `json:"hmnWorkClassCd" gorm:"column:hmn_work_class_cd"`
	FrgInjuryCount    int    `json:"frgInjuryCount" gorm:"column:frg_injury_count"`
	FutureTreatPlan   string `json:"futureTreatPlan" gorm:"column:future_treat_plan"`
	AccdDetail        string `json:"accdDetail" gorm:"column:accd_detail"`
	ZipNo             string `json:"zipNo" gorm:"column:zip_no"`
	DspthDate         string `json:"dspthDate" gorm:"column:dspth_date"`
	DamageAmt         int    `json:"damageAmt" gorm:"column:damage_amt"`
	FloorUnder        int    `json:"floorUnder" gorm:"column:floor_under"`
	AdmCd             string `json:"admCd" gorm:"column:adm_cd"`
	HmnAccdObjCd      string `json:"hmnAccdObjCd" gorm:"column:hmn_accd_obj_cd"`
	PublicCd          string `json:"publicCd" gorm:"column:public_cd"`
	ConstNo           string `json:"constNo" gorm:"column:const_no"`
	ConstNm           string `json:"constNm" gorm:"column:const_nm"`
	SysRegDate        string `json:"sysRegDate" gorm:"column:sys_reg_date"`
	AccdAgent         string `json:"accdAgent" gorm:"column:accd_agent"`
	MtrlAccdProcCd    string `json:"mtrlAccdProcCd" gorm:"column:mtrl_accd_proc_cd"`
	AccdKindCd        string `json:"accdKindCd" gorm:"column:accd_kind_cd"`
	FloorSpace        int    `json:"floorSpace" gorm:"column:floor_space"`
	AccdDate          string `json:"accdDate" gorm:"column:accd_date"`
	SafePlanTypeCd    string `json:"safePlanTypeCd" gorm:"column:safe_plan_type_cd"`
	MtrlWorkClassCd   string `json:"mtrlWorkClassCd" gorm:"column:mtrl_work_class_cd"`
	AccdIdvprotActYn  string `json:"accdIdvprotActYn" gorm:"column:accd_idvprot_act_yn"`
	AccdObjCd         string `json:"accdObjCd" gorm:"column:accd_obj_cd"`
	WcConstAmtCd      string `json:"wcConstAmtCd" gorm:"column:wc_const_amt_cd"`
	SysUpdDate        string `json:"sysUpdDate" gorm:"column:sys_upd_date"`
	LclDeathCount     int    `json:"lclDeathCount" gorm:"column:lcl_death_count"`
	TreatContent      string `json:"treatContent" gorm:"column:treat_content"`
	AccdBidRate       string `json:"accdBidRate" gorm:"column:accd_bid_rate"`
	AdtnlInpstCd      string `json:"adtnlInpstCd" gorm:"column:adtnl_inpst_cd"`
	CognPointCd       string `json:"cognPointCd" gorm:"column:cogn_point_cd"`
	SurveyMethodCd    string `json:"surveyMethodCd" gorm:"column:survey_method_cd"`
	HmnAccdKindCd     string `json:"hmnAccdKindCd" gorm:"column:hmn_accd_kind_cd"`
	AddrDong          string `json:"addrDong" gorm:"column:addr_dong"`
	AddrGugun         string `json:"addrGugun" gorm:"column:addr_gugun"`
	LclInjuryCount    int    `json:"lclInjuryCount" gorm:"column:lcl_injury_count"`
	Addr              string `json:"addr" gorm:"column:addr"`
	ConstAmtCd        string `json:"constAmtCd" gorm:"column:const_amt_cd"`
	AccdReason        string `json:"accdReason" gorm:"column:accd_reason"`
	AccdAssReason1    string `json:"accdAssReason1" gorm:"column:accd_ass_reason_1"`
	HmnAccdProcCd     string `json:"hmnAccdProcCd" gorm:"column:hmn_accd_proc_cd"`
	EndYmd            string `json:"endYmd" gorm:"column:end_ymd"`
	NtfcDate          string `json:"ntfcDate" gorm:"column:ntfc_date"`
	AccdDfsRsnCd      string `json:"accdDfsRsnCd" gorm:"column:accd_dfs_rsn_cd"`
	WcConstEndYmd     string `json:"wcConstEndYmd" gorm:"column:wc_const_end_ymd"`
	FloorOver         int    `json:"floorOver" gorm:"column:floor_over"`
	CaseNo            int    `json:"caseNo" gorm:"column:case_no"`
	WorkClassCd       string `json:"workClassCd" gorm:"column:work_class_cd"`
	StartYmd          string `json:"startYmd" gorm:"column:start_ymd"`
	AccdProcCd        string `json:"accdProcCd" gorm:"column:accd_proc_cd"`
	AccdWorkerCnt     string `json:"accdWorkerCnt" gorm:"column:accd_worker_cnt"`
	AddrDetail        string `json:"addrDetail" gorm:"column:addr_detail"`
	AccdPlaceCd       string `json:"accdPlaceCd" gorm:"column:accd_place_cd"`
	AccdReasonDetail  string `json:"accdReasonDetail" gorm:"column:accd_reason_detail"`
	FcltyKindCd       string `json:"fcltyKindCd" gorm:"column:fclty_kind_cd"`
	AccdAssReason2    string `json:"accdAssReason2" gorm:"column:accd_ass_reason_2"`
	AccdProgRate      string `json:"accdProgRate" gorm:"column:accd_prog_rate"`
	ReportReason      string `json:"reportReason" gorm:"column:report_reason"`
	PlacePicture      string `json:"placePicture" gorm:"column:place_picture"`
}

// 테이블 네임
func (m Item) TableName() string {
	return "csi_info_table"
}


type Model struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
			Status     int    `json:"status"`
		} `json:"header"`
		Body struct {
			TotalCount int    `json:"totalCount"`
			Items      []Item `json:"items"`
			NumOfRows  int    `json:"numOfRows"`
			PageNo     int    `json:"pageNo"`
		} `json:"body"`
	} `json:"response"`
}
