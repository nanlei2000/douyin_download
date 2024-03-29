package model

// Generated by https://quicktype.io

type Show struct {
	Visible               Visible               `json:"visible"`
	CreatedAt             string                `json:"created_at"`
	ID                    int64                 `json:"id"`
	Idstr                 string                `json:"idstr"`
	Mid                   string                `json:"mid"`
	Mblogid               string                `json:"mblogid"`
	User                  User                  `json:"user"`
	CanEdit               bool                  `json:"can_edit"`
	TextRaw               string                `json:"text_raw"`
	Text                  string                `json:"text"`
	TextLength            int64                 `json:"textLength"`
	Source                string                `json:"source"`
	Favorited             bool                  `json:"favorited"`
	Cardid                string                `json:"cardid"`
	PicIDS                []string              `json:"pic_ids"`
	PicFocusPoint         []PicFocusPoint       `json:"pic_focus_point"`
	Geo                   interface{}           `json:"geo"`
	PicNum                int64                 `json:"pic_num"`
	PicInfos              map[string]PicInfo    `json:"pic_infos"`
	IsPaid                bool                  `json:"is_paid"`
	PicBgNew              string                `json:"pic_bg_new"`
	MblogVipType          int64                 `json:"mblog_vip_type"`
	NumberDisplayStrategy NumberDisplayStrategy `json:"number_display_strategy"`
	RepostsCount          int64                 `json:"reposts_count"`
	CommentsCount         int64                 `json:"comments_count"`
	AttitudesCount        int64                 `json:"attitudes_count"`
	AttitudesStatus       int64                 `json:"attitudes_status"`
	IsLongText            bool                  `json:"isLongText"`
	Mlevel                int64                 `json:"mlevel"`
	ContentAuth           int64                 `json:"content_auth"`
	IsShowBulletin        int64                 `json:"is_show_bulletin"`
	CommentManageInfo     CommentManageInfo     `json:"comment_manage_info"`
	ShareRepostType       int64                 `json:"share_repost_type"`
	Title                 Title                 `json:"title"`
	Mblogtype             int64                 `json:"mblogtype"`
	ShowFeedRepost        bool                  `json:"showFeedRepost"`
	ShowFeedComment       bool                  `json:"showFeedComment"`
	PictureViewerSign     bool                  `json:"pictureViewerSign"`
	ShowPictureViewer     bool                  `json:"showPictureViewer"`
	RCList                []interface{}         `json:"rcList"`
	RegionName            string                `json:"region_name"`
	CustomIcons           []interface{}         `json:"customIcons"`
	Ok                    int64                 `json:"ok"`
}

type CommentManageInfo struct {
	CommentPermissionType int64 `json:"comment_permission_type"`
	ApprovalCommentType   int64 `json:"approval_comment_type"`
	CommentSortType       int64 `json:"comment_sort_type"`
}

type NumberDisplayStrategy struct {
	ApplyScenarioFlag    int64  `json:"apply_scenario_flag"`
	DisplayTextMinNumber int64  `json:"display_text_min_number"`
	DisplayText          string `json:"display_text"`
}

type PicFocusPoint struct {
	FocusPoint FocusPoint `json:"focus_point"`
	PicID      string     `json:"pic_id"`
}

type FocusPoint struct {
	Left   float64 `json:"left"`
	Top    float64 `json:"top"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type PicInfo struct {
	Thumbnail  Bmiddle     `json:"thumbnail"`
	Bmiddle    Bmiddle     `json:"bmiddle"`
	Large      Bmiddle     `json:"large"`
	Original   Bmiddle     `json:"original"`
	Largest    Bmiddle     `json:"largest"`
	Mw2000     Bmiddle     `json:"mw2000"`
	FocusPoint *FocusPoint `json:"focus_point,omitempty"`
	ObjectID   string      `json:"object_id"`
	PicID      string      `json:"pic_id"`
	PhotoTag   int64       `json:"photo_tag"`
	Type       string      `json:"type"`
	PicStatus  int64       `json:"pic_status"`
}

type Bmiddle struct {
	URL     string      `json:"url"`
	Width   int64       `json:"width"`
	Height  int64       `json:"height"`
	CutType int64       `json:"cut_type"`
	Type    interface{} `json:"type"`
}

type Title struct {
	Text      string `json:"text"`
	BaseColor int64  `json:"base_color"`
	IconURL   string `json:"icon_url"`
}

type User struct {
	ID              int64      `json:"id"`
	Idstr           string     `json:"idstr"`
	PCNew           int64      `json:"pc_new"`
	ScreenName      string     `json:"screen_name"`
	ProfileImageURL string     `json:"profile_image_url"`
	ProfileURL      string     `json:"profile_url"`
	Verified        bool       `json:"verified"`
	VerifiedType    int64      `json:"verified_type"`
	Domain          string     `json:"domain"`
	Weihao          string     `json:"weihao"`
	VerifiedTypeEXT int64      `json:"verified_type_ext"`
	AvatarLarge     string     `json:"avatar_large"`
	AvatarHD        string     `json:"avatar_hd"`
	FollowMe        bool       `json:"follow_me"`
	Following       bool       `json:"following"`
	Mbrank          int64      `json:"mbrank"`
	Mbtype          int64      `json:"mbtype"`
	PlanetVideo     bool       `json:"planet_video"`
	IconList        []IconList `json:"icon_list"`
}

type IconList struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	Mbrank int64 `json:"mbrank"`
	Mbtype int64 `json:"mbtype"`
}

type Visible struct {
	Type   int64 `json:"type"`
	ListID int64 `json:"list_id"`
}
