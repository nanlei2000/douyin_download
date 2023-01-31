package iteminfo

type ItemInfo struct {
	AwemeDetail AwemeDetail `json:"aweme_detail"`
	LogPb       LogPb       `json:"log_pb"`
	StatusCode  int64       `json:"status_code"`
}

type AwemeDetail struct {
	Anchors                 interface{}           `json:"anchors"`
	Author                  Author                `json:"author"`
	AuthorMaskTag           int64                 `json:"author_mask_tag"`
	AuthorUserID            int64                 `json:"author_user_id"`
	AwemeControl            AwemeControl          `json:"aweme_control"`
	AwemeID                 string                `json:"aweme_id"`
	AwemeType               int64                 `json:"aweme_type"`
	ChallengePosition       interface{}           `json:"challenge_position"`
	ChapterList             interface{}           `json:"chapter_list"`
	CollectStat             int64                 `json:"collect_stat"`
	CommentGid              int64                 `json:"comment_gid"`
	CommentList             interface{}           `json:"comment_list"`
	CommentPermissionInfo   CommentPermissionInfo `json:"comment_permission_info"`
	CommerceConfigData      interface{}           `json:"commerce_config_data"`
	CommonBarInfo           string                `json:"common_bar_info"`
	ComponentInfoV2         string                `json:"component_info_v2"`
	CoverLabels             interface{}           `json:"cover_labels"`
	CreateTime              int64                 `json:"create_time"`
	Desc                    string                `json:"desc"`
	DiggLottie              DiggLottie            `json:"digg_lottie"`
	DisableRelationBar      int64                 `json:"disable_relation_bar"`
	DislikeDimensionList    interface{}           `json:"dislike_dimension_list"`
	DuetAggregateInMusicTab bool                  `json:"duet_aggregate_in_music_tab"`
	Duration                int64                 `json:"duration"`
	Geofencing              []interface{}         `json:"geofencing"`
	GeofencingRegions       interface{}           `json:"geofencing_regions"`
	GroupID                 string                `json:"group_id"`
	HybridLabel             interface{}           `json:"hybrid_label"`
	ImageAlbumMusicInfo     ImageAlbumMusicInfo   `json:"image_album_music_info"`
	ImageInfos              interface{}           `json:"image_infos"`
	ImageList               interface{}           `json:"image_list"`
	Images                  interface{}           `json:"images"`
	ImgBitrate              interface{}           `json:"img_bitrate"`
	ImpressionData          ImpressionData        `json:"impression_data"`
	InteractionStickers     interface{}           `json:"interaction_stickers"`
	IsAds                   bool                  `json:"is_ads"`
	IsCollectsSelected      int64                 `json:"is_collects_selected"`
	IsDuetSing              bool                  `json:"is_duet_sing"`
	IsImageBeat             bool                  `json:"is_image_beat"`
	IsLifeItem              bool                  `json:"is_life_item"`
	IsStory                 int64                 `json:"is_story"`
	IsTop                   int64                 `json:"is_top"`
	ItemWarnNotification    ItemWarnNotification  `json:"item_warn_notification"`
	LabelTopText            interface{}           `json:"label_top_text"`
	LongVideo               interface{}           `json:"long_video"`
	Music                   Music                 `json:"music"`
	NicknamePosition        interface{}           `json:"nickname_position"`
	OriginCommentIDS        interface{}           `json:"origin_comment_ids"`
	OriginTextExtra         []interface{}         `json:"origin_text_extra"`
	OriginalImages          interface{}           `json:"original_images"`
	PackedClips             interface{}           `json:"packed_clips"`
	PhotoSearchEntrance     PhotoSearchEntrance   `json:"photo_search_entrance"`
	Position                interface{}           `json:"position"`
	PreviewTitle            string                `json:"preview_title"`
	PreviewVideoStatus      int64                 `json:"preview_video_status"`
	Promotions              []interface{}         `json:"promotions"`
	Rate                    int64                 `json:"rate"`
	Region                  string                `json:"region"`
	RelationLabels          interface{}           `json:"relation_labels"`
	SearchImpr              AwemeDetailSearchImpr `json:"search_impr"`
	SeriesPaidInfo          SeriesPaidInfo        `json:"series_paid_info"`
	ShareInfo               AwemeDetailShareInfo  `json:"share_info"`
	ShareURL                string                `json:"share_url"`
	ShouldOpenAdReport      bool                  `json:"should_open_ad_report"`
	ShowFollowButton        ShowFollowButton      `json:"show_follow_button"`
	SocialTagList           interface{}           `json:"social_tag_list"`
	StandardBarInfoList     interface{}           `json:"standard_bar_info_list"`
	Statistics              Statistics            `json:"statistics"`
	Status                  Status                `json:"status"`
	TextExtra               []TextExtra           `json:"text_extra"`
	UniqidPosition          interface{}           `json:"uniqid_position"`
	UserDigged              int64                 `json:"user_digged"`
	Video                   Video                 `json:"video"`
	VideoLabels             interface{}           `json:"video_labels"`
	VideoTag                []VideoTag            `json:"video_tag"`
	VideoText               []interface{}         `json:"video_text"`
	WannaTag                ShowFollowButton      `json:"wanna_tag"`
}

type Author struct {
	AvatarThumb                            AvatarThumb     `json:"avatar_thumb"`
	CFList                                 interface{}     `json:"cf_list"`
	CloseFriendType                        int64           `json:"close_friend_type"`
	ContactsStatus                         int64           `json:"contacts_status"`
	ContrailList                           interface{}     `json:"contrail_list"`
	CoverURL                               []AvatarThumb   `json:"cover_url"`
	CreateTime                             int64           `json:"create_time"`
	CustomVerify                           string          `json:"custom_verify"`
	DataLabelList                          interface{}     `json:"data_label_list"`
	EndorsementInfoList                    interface{}     `json:"endorsement_info_list"`
	EnterpriseVerifyReason                 string          `json:"enterprise_verify_reason"`
	FavoritingCount                        int64           `json:"favoriting_count"`
	FollowStatus                           int64           `json:"follow_status"`
	FollowerCount                          int64           `json:"follower_count"`
	FollowerListSecondaryInformationStruct interface{}     `json:"follower_list_secondary_information_struct"`
	FollowerStatus                         int64           `json:"follower_status"`
	FollowingCount                         int64           `json:"following_count"`
	IMRoleIDS                              interface{}     `json:"im_role_ids"`
	IsAdFake                               bool            `json:"is_ad_fake"`
	IsBlockedV2                            bool            `json:"is_blocked_v2"`
	IsBlockingV2                           bool            `json:"is_blocking_v2"`
	IsCF                                   int64           `json:"is_cf"`
	MaxFollowerCount                       int64           `json:"max_follower_count"`
	Nickname                               string          `json:"nickname"`
	NotSeenItemIDList                      interface{}     `json:"not_seen_item_id_list"`
	NotSeenItemIDListV2                    interface{}     `json:"not_seen_item_id_list_v2"`
	OfflineInfoList                        interface{}     `json:"offline_info_list"`
	PersonalTagList                        interface{}     `json:"personal_tag_list"`
	PreventDownload                        bool            `json:"prevent_download"`
	RiskNoticeText                         string          `json:"risk_notice_text"`
	SECUid                                 string          `json:"sec_uid"`
	Secret                                 int64           `json:"secret"`
	ShareInfo                              AuthorShareInfo `json:"share_info"`
	ShortID                                string          `json:"short_id"`
	Signature                              string          `json:"signature"`
	SignatureExtra                         interface{}     `json:"signature_extra"`
	SpecialPeopleLabels                    interface{}     `json:"special_people_labels"`
	Status                                 int64           `json:"status"`
	TextExtra                              interface{}     `json:"text_extra"`
	TotalFavorited                         int64           `json:"total_favorited"`
	Uid                                    string          `json:"uid"`
	UniqueID                               string          `json:"unique_id"`
	UserAge                                int64           `json:"user_age"`
	UserCanceled                           bool            `json:"user_canceled"`
	UserPermissions                        interface{}     `json:"user_permissions"`
	VerificationType                       int64           `json:"verification_type"`
}

type AvatarThumb struct {
	Height  int64    `json:"height"`
	URI     string   `json:"uri"`
	URLList []string `json:"url_list"`
	Width   int64    `json:"width"`
	URLKey  *string  `json:"url_key,omitempty"`
}

type AuthorShareInfo struct {
	ShareDesc        string      `json:"share_desc"`
	ShareDescInfo    string      `json:"share_desc_info"`
	ShareQrcodeURL   AvatarThumb `json:"share_qrcode_url"`
	ShareTitle       string      `json:"share_title"`
	ShareTitleMyself string      `json:"share_title_myself"`
	ShareTitleOther  string      `json:"share_title_other"`
	ShareURL         string      `json:"share_url"`
	ShareWeiboDesc   string      `json:"share_weibo_desc"`
}

type AwemeControl struct {
	CanComment     bool `json:"can_comment"`
	CanForward     bool `json:"can_forward"`
	CanShare       bool `json:"can_share"`
	CanShowComment bool `json:"can_show_comment"`
}

type CommentPermissionInfo struct {
	CanComment              bool  `json:"can_comment"`
	CommentPermissionStatus int64 `json:"comment_permission_status"`
	ItemDetailEntry         bool  `json:"item_detail_entry"`
	PressEntry              bool  `json:"press_entry"`
	ToastGuide              bool  `json:"toast_guide"`
}

type DiggLottie struct {
	CanBomb  int64  `json:"can_bomb"`
	LottieID string `json:"lottie_id"`
}

type ImageAlbumMusicInfo struct {
	BeginTime int64 `json:"begin_time"`
	EndTime   int64 `json:"end_time"`
	Volume    int64 `json:"volume"`
}

type ImpressionData struct {
	GroupIDListA   []int64     `json:"group_id_list_a"`
	GroupIDListB   []int64     `json:"group_id_list_b"`
	SimilarIDListA interface{} `json:"similar_id_list_a"`
	SimilarIDListB interface{} `json:"similar_id_list_b"`
}

type ItemWarnNotification struct {
	Content string `json:"content"`
	Show    bool   `json:"show"`
	Type    int64  `json:"type"`
}

type Music struct {
	Album                     string        `json:"album"`
	ArtistUserInfos           interface{}   `json:"artist_user_infos"`
	Artists                   []interface{} `json:"artists"`
	AuditionDuration          int64         `json:"audition_duration"`
	Author                    string        `json:"author"`
	AuthorDeleted             bool          `json:"author_deleted"`
	AuthorPosition            interface{}   `json:"author_position"`
	AuthorStatus              int64         `json:"author_status"`
	AvatarLarge               AvatarThumb   `json:"avatar_large"`
	AvatarMedium              AvatarThumb   `json:"avatar_medium"`
	AvatarThumb               AvatarThumb   `json:"avatar_thumb"`
	BindedChallengeID         int64         `json:"binded_challenge_id"`
	CanBackgroundPlay         bool          `json:"can_background_play"`
	CollectStat               int64         `json:"collect_stat"`
	CoverHD                   AvatarThumb   `json:"cover_hd"`
	CoverLarge                AvatarThumb   `json:"cover_large"`
	CoverMedium               AvatarThumb   `json:"cover_medium"`
	CoverThumb                AvatarThumb   `json:"cover_thumb"`
	DmvAutoShow               bool          `json:"dmv_auto_show"`
	DSPStatus                 int64         `json:"dsp_status"`
	Duration                  int64         `json:"duration"`
	EndTime                   int64         `json:"end_time"`
	ExternalSongInfo          []interface{} `json:"external_song_info"`
	Extra                     string        `json:"extra"`
	ID                        int64         `json:"id"`
	IDStr                     string        `json:"id_str"`
	IsAudioURLWithCookie      bool          `json:"is_audio_url_with_cookie"`
	IsCommerceMusic           bool          `json:"is_commerce_music"`
	IsDelVideo                bool          `json:"is_del_video"`
	IsMatchedMetadata         bool          `json:"is_matched_metadata"`
	IsOriginal                bool          `json:"is_original"`
	IsOriginalSound           bool          `json:"is_original_sound"`
	IsPgc                     bool          `json:"is_pgc"`
	IsRestricted              bool          `json:"is_restricted"`
	IsVideoSelfSee            bool          `json:"is_video_self_see"`
	LunaInfo                  LunaInfo      `json:"luna_info"`
	LyricShortPosition        interface{}   `json:"lyric_short_position"`
	Mid                       string        `json:"mid"`
	MusicChartRanks           interface{}   `json:"music_chart_ranks"`
	MusicStatus               int64         `json:"music_status"`
	MusicianUserInfos         interface{}   `json:"musician_user_infos"`
	MuteShare                 bool          `json:"mute_share"`
	OfflineDesc               string        `json:"offline_desc"`
	OwnerHandle               string        `json:"owner_handle"`
	OwnerID                   string        `json:"owner_id"`
	OwnerNickname             string        `json:"owner_nickname"`
	PgcMusicType              int64         `json:"pgc_music_type"`
	PlayURL                   AvatarThumb   `json:"play_url"`
	Position                  interface{}   `json:"position"`
	PreventDownload           bool          `json:"prevent_download"`
	PreventItemDownloadStatus int64         `json:"prevent_item_download_status"`
	// PreviewEndTime            int64         `json:"preview_end_time"`
	// PreviewStartTime          int64           `json:"preview_start_time"`
	ReasonType        int64           `json:"reason_type"`
	Redirect          bool            `json:"redirect"`
	SchemaURL         string          `json:"schema_url"`
	SearchImpr        MusicSearchImpr `json:"search_impr"`
	SECUid            string          `json:"sec_uid"`
	ShootDuration     int64           `json:"shoot_duration"`
	SourcePlatform    int64           `json:"source_platform"`
	StartTime         int64           `json:"start_time"`
	Status            int64           `json:"status"`
	TagList           interface{}     `json:"tag_list"`
	Title             string          `json:"title"`
	UnshelveCountries interface{}     `json:"unshelve_countries"`
	UserCount         int64           `json:"user_count"`
	VideoDuration     int64           `json:"video_duration"`
}

type LunaInfo struct {
	IsLunaUser bool `json:"is_luna_user"`
}

type MusicSearchImpr struct {
	EntityID string `json:"entity_id"`
}

type PhotoSearchEntrance struct {
	EcomType int64 `json:"ecom_type"`
}

type AwemeDetailSearchImpr struct {
	EntityID   string `json:"entity_id"`
	EntityType string `json:"entity_type"`
}

type SeriesPaidInfo struct {
	ItemPrice        int64 `json:"item_price"`
	SeriesPaidStatus int64 `json:"series_paid_status"`
}

type AwemeDetailShareInfo struct {
	ShareDesc     string `json:"share_desc"`
	ShareDescInfo string `json:"share_desc_info"`
	ShareLinkDesc string `json:"share_link_desc"`
	ShareURL      string `json:"share_url"`
}

type ShowFollowButton struct {
}

type Statistics struct {
	AdmireCount  int64  `json:"admire_count"`
	AwemeID      string `json:"aweme_id"`
	CollectCount int64  `json:"collect_count"`
	CommentCount int64  `json:"comment_count"`
	DiggCount    int64  `json:"digg_count"`
	PlayCount    int64  `json:"play_count"`
	ShareCount   int64  `json:"share_count"`
}

type Status struct {
	AllowShare        bool         `json:"allow_share"`
	AwemeID           string       `json:"aweme_id"`
	InReviewing       bool         `json:"in_reviewing"`
	IsDelete          bool         `json:"is_delete"`
	IsProhibited      bool         `json:"is_prohibited"`
	ListenVideoStatus int64        `json:"listen_video_status"`
	PartSee           int64        `json:"part_see"`
	PrivateStatus     int64        `json:"private_status"`
	ReviewResult      ReviewResult `json:"review_result"`
}

type ReviewResult struct {
	ReviewStatus int64 `json:"review_status"`
}

type TextExtra struct {
	End         int64  `json:"end"`
	HashtagID   string `json:"hashtag_id"`
	HashtagName string `json:"hashtag_name"`
	IsCommerce  bool   `json:"is_commerce"`
	Start       int64  `json:"start"`
	Type        int64  `json:"type"`
}

type Video struct {
	BigThumbs          []BigThumb  `json:"big_thumbs"`
	BitRate            []BitRate   `json:"bit_rate"`
	Cover              AvatarThumb `json:"cover"`
	CoverOriginalScale AvatarThumb `json:"cover_original_scale"`
	Duration           int64       `json:"duration"`
	DynamicCover       AvatarThumb `json:"dynamic_cover"`
	Height             int64       `json:"height"`
	IsH265             int64       `json:"is_h265"`
	IsLongVideo        int64       `json:"is_long_video"`
	IsSourceHDR        int64       `json:"is_source_HDR"`
	Meta               string      `json:"meta"`
	OriginCover        AvatarThumb `json:"origin_cover"`
	PlayAddr           PlayAddr    `json:"play_addr"`
	Ratio              string      `json:"ratio"`
	VideoModel         string      `json:"video_model"`
	Width              int64       `json:"width"`
}

type BigThumb struct {
	Duration float64 `json:"duration"`
	Fext     string  `json:"fext"`
	ImgNum   int64   `json:"img_num"`
	ImgURL   string  `json:"img_url"`
	ImgXLen  int64   `json:"img_x_len"`
	ImgXSize int64   `json:"img_x_size"`
	ImgYLen  int64   `json:"img_y_len"`
	ImgYSize int64   `json:"img_y_size"`
	Interval int64   `json:"interval"`
	URI      string  `json:"uri"`
}

type BitRate struct {
	FPS         int64    `json:"FPS"`
	HDRBit      string   `json:"HDR_bit"`
	HDRType     string   `json:"HDR_type"`
	BitRate     int64    `json:"bit_rate"`
	GearName    string   `json:"gear_name"`
	IsBytevc1   int64    `json:"is_bytevc1"`
	IsH265      int64    `json:"is_h265"`
	PlayAddr    PlayAddr `json:"play_addr"`
	QualityType int64    `json:"quality_type"`
}

type PlayAddr struct {
	DataSize int64    `json:"data_size"`
	FileCS   string   `json:"file_cs"`
	FileHash string   `json:"file_hash"`
	Height   int64    `json:"height"`
	URI      string   `json:"uri"`
	URLKey   string   `json:"url_key"`
	URLList  []string `json:"url_list"`
	Width    int64    `json:"width"`
}

type VideoTag struct {
	Level   int64  `json:"level"`
	TagID   int64  `json:"tag_id"`
	TagName string `json:"tag_name"`
}

type LogPb struct {
	ImprID string `json:"impr_id"`
}
