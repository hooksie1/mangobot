package p

import "time"

type Message struct {
	Message struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

type Response struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type mangoUser struct {
	UID                                 int           `json:"uid"`
	Username                            string        `json:"username"`
	Userslug                            string        `json:"userslug"`
	Email                               string        `json:"email"`
	EmailConfirmed                      bool          `json:"email:confirmed"`
	Joindate                            int64         `json:"joindate"`
	Lastonline                          int64         `json:"lastonline"`
	Picture                             string        `json:"picture"`
	Fullname                            string        `json:"fullname"`
	Location                            string        `json:"location"`
	Birthday                            string        `json:"birthday"`
	Website                             string        `json:"website"`
	Signature                           string        `json:"signature"`
	Uploadedpicture                     string        `json:"uploadedpicture"`
	Profileviews                        int           `json:"profileviews"`
	Reputation                          int           `json:"reputation"`
	Postcount                           int           `json:"postcount"`
	Topiccount                          int           `json:"topiccount"`
	Lastposttime                        int64         `json:"lastposttime"`
	Banned                              bool          `json:"banned"`
	BannedExpire                        int64         `json:"banned:expire"`
	Status                              string        `json:"status"`
	Flags                               int           `json:"flags"`
	FollowerCount                       int           `json:"followerCount"`
	FollowingCount                      int           `json:"followingCount"`
	CoverURL                            string        `json:"cover:url"`
	CoverPosition                       string        `json:"cover:position"`
	GroupTitle                          string        `json:"groupTitle"`
	GroupTitleArray                     []string      `json:"groupTitleArray"`
	IconText                            string        `json:"icon:text"`
	IconBgColor                         string        `json:"icon:bgColor"`
	JoindateISO                         string        `json:"joindateISO"`
	LastonlineISO                       string        `json:"lastonlineISO"`
	BannedUntil                         int64         `json:"banned_until"`
	BannedUntilReadable                 time.Time     `json:"banned_until_readable"`
	Age                                 int           `json:"age"`
	EmailClass                          string        `json:"emailClass"`
	ModerationNote                      string        `json:"moderationNote"`
	IsBlocked                           bool          `json:"isBlocked"`
	Yourid                              int           `json:"yourid"`
	Theirid                             int           `json:"theirid"`
	IsTargetAdmin                       bool          `json:"isTargetAdmin"`
	IsAdmin                             bool          `json:"isAdmin"`
	IsGlobalModerator                   bool          `json:"isGlobalModerator"`
	IsModerator                         bool          `json:"isModerator"`
	IsAdminOrGlobalModerator            bool          `json:"isAdminOrGlobalModerator"`
	IsAdminOrGlobalModeratorOrModerator bool          `json:"isAdminOrGlobalModeratorOrModerator"`
	IsSelfOrAdminOrGlobalModerator      bool          `json:"isSelfOrAdminOrGlobalModerator"`
	CanEdit                             bool          `json:"canEdit"`
	CanBan                              bool          `json:"canBan"`
	CanChangePassword                   bool          `json:"canChangePassword"`
	IsSelf                              bool          `json:"isSelf"`
	IsFollowing                         bool          `json:"isFollowing"`
	ShowHidden                          bool          `json:"showHidden"`
	Groups                              []interface{} `json:"groups"`
	DisableSignatures                   bool          `json:"disableSignatures"`
	ReputationDisabled                  bool          `json:"reputation:disabled"`
	DownvoteDisabled                    bool          `json:"downvote:disabled"`
	ProfileLinks                        []interface{} `json:"profile_links"`
	Sso                                 []interface{} `json:"sso"`
	WebsiteLink                         string        `json:"websiteLink"`
	WebsiteName                         string        `json:"websiteName"`
	UsernameDisableEdit                 int           `json:"username:disableEdit"`
	EmailDisableEdit                    int           `json:"email:disableEdit"`
	Posts                               []struct {
		Pid          int    `json:"pid"`
		Tid          int    `json:"tid"`
		Content      string `json:"content"`
		UID          int    `json:"uid"`
		Timestamp    int64  `json:"timestamp"`
		Deleted      bool   `json:"deleted"`
		Upvotes      int    `json:"upvotes"`
		Downvotes    int    `json:"downvotes"`
		Votes        int    `json:"votes"`
		TimestampISO string `json:"timestampISO"`
		User         struct {
			UID         int    `json:"uid"`
			Username    string `json:"username"`
			Userslug    string `json:"userslug"`
			Picture     string `json:"picture"`
			IconText    string `json:"icon:text"`
			IconBgColor string `json:"icon:bgColor"`
		} `json:"user"`
		Topic struct {
			UID       int    `json:"uid"`
			Tid       int    `json:"tid"`
			Title     string `json:"title"`
			Cid       int    `json:"cid"`
			Slug      string `json:"slug"`
			Deleted   int    `json:"deleted"`
			Postcount int    `json:"postcount"`
			MainPid   int    `json:"mainPid"`
			TitleRaw  string `json:"titleRaw"`
		} `json:"topic"`
		Category struct {
			Cid       int    `json:"cid"`
			Name      string `json:"name"`
			Icon      string `json:"icon"`
			Slug      string `json:"slug"`
			ParentCid int    `json:"parentCid"`
			BgColor   string `json:"bgColor"`
			Color     string `json:"color"`
		} `json:"category"`
		IsMainPost bool `json:"isMainPost"`
	} `json:"posts"`
	HasPrivateChat string `json:"hasPrivateChat"`
	NextStart      int    `json:"nextStart"`
	Breadcrumbs    []struct {
		Text string `json:"text"`
		URL  string `json:"url,omitempty"`
	} `json:"breadcrumbs"`
	Title             string `json:"title"`
	AllowCoverPicture bool   `json:"allowCoverPicture"`
	Pagination        struct {
		Rel []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"rel"`
		Pages []struct {
			Page      int    `json:"page,omitempty"`
			Active    bool   `json:"active,omitempty"`
			Qs        string `json:"qs,omitempty"`
			Separator bool   `json:"separator,omitempty"`
		} `json:"pages"`
		CurrentPage int `json:"currentPage"`
		PageCount   int `json:"pageCount"`
		Prev        struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"prev"`
		Next struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"next"`
		First struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"first"`
		Last struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"last"`
	} `json:"pagination"`
	SelectedGroup []interface{} `json:"selectedGroup"`
	LoggedIn      bool          `json:"loggedIn"`
	RelativePath  string        `json:"relative_path"`
	Template      struct {
		Name           string `json:"name"`
		AccountProfile bool   `json:"account/profile"`
	} `json:"template"`
	URL       string `json:"url"`
	BodyClass string `json:"bodyClass"`
	Widgets   struct {
		Sidebar []struct {
			HTML string `json:"html"`
		} `json:"sidebar"`
		Footer []struct {
			HTML string `json:"html"`
		} `json:"footer"`
	} `json:"widgets"`
}

type Categories struct {
	Title      string `json:"title"`
	Categories []struct {
			Background        string        `json:"background,omitempty"`
			BgColor           string        `json:"bgColor"`
			Cid               int           `json:"cid"`
			Class             string        `json:"class"`
			Color             string        `json:"color"`
			Description       string        `json:"description"`
			Disabled          int           `json:"disabled"`
			Icon              string        `json:"icon"`
			ImageClass        string        `json:"imageClass"`
			Link              string        `json:"link"`
			Name              string        `json:"name"`
			NumRecentReplies  int           `json:"numRecentReplies"`
			Order             int           `json:"order"`
			PostCount         int           `json:"post_count"`
			Slug              string        `json:"slug"`
			TopicCount        int           `json:"topic_count"`
			ParentCid         int           `json:"parentCid"`
			IsSection         int           `json:"isSection"`
			TotalPostCount    int           `json:"totalPostCount"`
			TotalTopicCount   int           `json:"totalTopicCount"`
			DescriptionParsed string        `json:"descriptionParsed,omitempty"`
			TagWhitelist      []interface{} `json:"tagWhitelist"`
			UnreadClass       string        `json:"unread-class"`
			Children          []interface{} `json:"children"`
			Posts             []struct {
					Pid          int       `json:"pid"`
					Timestamp    int64     `json:"timestamp"`
					Content      string    `json:"content"`
					TimestampISO time.Time `json:"timestampISO"`
					User         struct {
							Username    string `json:"username"`
							Userslug    string `json:"userslug"`
							Picture     string `json:"picture"`
							IconText    string `json:"icon:text"`
							IconBgColor string `json:"icon:bgColor"`
					} `json:"user"`
					Index     int `json:"index"`
					Cid       int `json:"cid"`
					ParentCid int `json:"parentCid"`
					Topic     struct {
							Slug  string `json:"slug"`
							Title string `json:"title"`
					} `json:"topic"`
			} `json:"posts"`
			Teaser struct {
					URL          string    `json:"url"`
					TimestampISO time.Time `json:"timestampISO"`
					Pid          int       `json:"pid"`
					Topic        struct {
							Slug  string `json:"slug"`
							Title string `json:"title"`
					} `json:"topic"`
			} `json:"teaser,omitempty"`
			Image           string `json:"image,omitempty"`
			BackgroundImage string `json:"backgroundImage,omitempty"`
	} `json:"categories"`
	Breadcrumbs []struct {
			Text string `json:"text"`
			URL  string `json:"url,omitempty"`
	} `json:"breadcrumbs"`
	LoggedIn     bool   `json:"loggedIn"`
	RelativePath string `json:"relative_path"`
	Template     struct {
			Name       string `json:"name"`
			Categories bool   `json:"categories"`
	} `json:"template"`
	URL       string `json:"url"`
	BodyClass string `json:"bodyClass"`
	Widgets   struct {
			Footer []struct {
					HTML string `json:"html"`
			} `json:"footer"`
			Sidebar []struct {
					HTML string `json:"html"`
			} `json:"sidebar"`
	} `json:"widgets"`
}

type Users struct {                                                                                                                                                                                                                                                                                                        
	Users []struct {                                                                                                                                                                                                                                                                                                   
			UID                 int         `json:"uid"`                                                                                                                                                                                                                                                               
			Username            string      `json:"username"`                                                                                                                                                                                                                                                          
			Userslug            string      `json:"userslug"`                                                                                                                                                                                                                                                          
			Picture             string      `json:"picture"`                                                                                                                                                                                                                                                           
			Status              string      `json:"status"`                                                                                                                                                                                                                                                            
			Postcount           int         `json:"postcount"`                                                                                                                                                                                                                                                         
			Reputation          int         `json:"reputation"`                                                                                                                                                                                                                                                        
			EmailConfirmed      int         `json:"email:confirmed"`                                                                                                                                                                                                                                                   
			Lastonline          int64       `json:"lastonline"`                                                                                                                                                                                                                                                        
			Flags               interface{} `json:"flags"`                                                                                                                                                                                                                                                             
			Banned              int         `json:"banned"`                                                                                                                                                                                                                                                            
			BannedExpire        int         `json:"banned:expire"`                                                                                                                                                                                                                                                     
			Joindate            int64       `json:"joindate"`                                                                                                                                                                                                                                                          
			IconText            string      `json:"icon:text"`                                                                                                                                                                                                                                                         
			IconBgColor         string      `json:"icon:bgColor"`                                                                                                                                                                                                                                                      
			JoindateISO         time.Time   `json:"joindateISO"`                                                                                                                                                                                                                                                       
			LastonlineISO       time.Time   `json:"lastonlineISO"`                                                                                                                                                                                                                                                     
			BannedUntil         int         `json:"banned_until"`                                                                                                                                                                                                                                                      
			BannedUntilReadable string      `json:"banned_until_readable"`                                                                                                                                                                                                                                             
			Administrator       bool        `json:"administrator"`                                                                                                                                                                                                                                                     
	} `json:"users"`                                                                                                                                                                                                                                                                                                   
	Pagination struct {                                                                                                                                                                                                                                                                                                
			Rel []struct {                                                                                                                                                                                                                                                                                             
					Rel  string `json:"rel"`                                                                                                                                                                                                                                                                           
					Href string `json:"href"`                                                                                                                                                                                                                                                                          
			} `json:"rel"`                                                                                                                                                                                                                                                                                             
			Pages []struct {                                                                                                                                                                                                                                                                                           
					Page      int    `json:"page,omitempty"`                                                                                                                                                                                                                                                           
					Active    bool   `json:"active,omitempty"`                                                                                                                                                                                                                                                         
					Qs        string `json:"qs,omitempty"`                                                                                                                                                                                                                                                             
					Separator bool   `json:"separator,omitempty"`                                                                                                                                                                                                                                                      
			} `json:"pages"`                                                                                                                                                                                                                                                                                           
			CurrentPage int `json:"currentPage"`                                                                                                                                                                                                                                                                       
			PageCount   int `json:"pageCount"`                                                                                                                                                                                                                                                                         
			Prev        struct {                                                                                                                                                                                                                                                                                       
					Page   int    `json:"page"`                                                                                                                                                                                                                                                                        
					Active bool   `json:"active"`                                                                                                                                                                                                                                                                      
					Qs     string `json:"qs"`                             
			} `json:"prev"`                                               
			Next struct {                                                 
					Page   int    `json:"page"`                           
					Active bool   `json:"active"`                         
					Qs     string `json:"qs"`                             
			} `json:"next"`                                               
			First struct {                                                
					Page   int    `json:"page"`                           
					Active bool   `json:"active"`                         
					Qs     string `json:"qs"`                             
			} `json:"first"`                                              
			Last struct {                                                 
					Page   int    `json:"page"`                           
					Active bool   `json:"active"`                         
					Qs     string `json:"qs"`                             
			} `json:"last"`                                               
	} `json:"pagination"`                                                 
	UserCount   int    `json:"userCount"`                                 
	Title       string `json:"title"`                                     
	Breadcrumbs []struct {                                                
			Text string `json:"text"`                                     
			URL  string `json:"url,omitempty"`                            
	} `json:"breadcrumbs"`                                                
	IsAdminOrGlobalMod bool   `json:"isAdminOrGlobalMod"`                                                                                                
	DisplayUserSearch  bool   `json:"displayUserSearch"`                                                                                                 
	SectionJoindate    bool   `json:"section_joindate"`                                                                                                  
	MaximumInvites     int    `json:"maximumInvites"`                     
	InviteOnly         bool   `json:"inviteOnly"`                         
	AdminInviteOnly    bool   `json:"adminInviteOnly"`                    
	ReputationDisabled int    `json:"reputation:disabled"`                                                                                               
	Invites            int    `json:"invites"`                            
	LoggedIn           bool   `json:"loggedIn"`                           
	RelativePath       string `json:"relative_path"`                      
	Template           struct {                                           
			Name  string `json:"name"`                                    
			Users bool   `json:"users"`                                   
	} `json:"template"`                                                   
	URL       string `json:"url"`                                         
	BodyClass string `json:"bodyClass"`                                   
	Widgets   struct {                                                    
			Sidebar []struct {                                            
					HTML string `json:"html"`                             
			} `json:"sidebar"`                                            
			Footer []struct {                                             
					HTML string `json:"html"`                             
			} `json:"footer"`                                             
	} `json:"widgets"`                                                    
}   

type Recent struct {
	NextStart  int `json:"nextStart"`
	TopicCount int `json:"topicCount"`
	Topics     []struct {
		Tid             int       `json:"tid"`
		UID             int       `json:"uid"`
		Cid             int       `json:"cid"`
		MainPid         int       `json:"mainPid"`
		Title           string    `json:"title"`
		Slug            string    `json:"slug"`
		Timestamp       int64     `json:"timestamp"`
		Lastposttime    int64     `json:"lastposttime"`
		Postcount       int       `json:"postcount"`
		Viewcount       int       `json:"viewcount"`
		Locked          int       `json:"locked"`
		Deleted         int       `json:"deleted"`
		Pinned          int       `json:"pinned"`
		TeaserPid       int       `json:"teaserPid"`
		Upvotes         int       `json:"upvotes"`
		Downvotes       int       `json:"downvotes"`
		TitleRaw        string    `json:"titleRaw"`
		TimestampISO    time.Time `json:"timestampISO"`
		LastposttimeISO time.Time `json:"lastposttimeISO"`
		Votes           int       `json:"votes"`
		Category        struct {
			Cid        int         `json:"cid"`
			Name       string      `json:"name"`
			Slug       string      `json:"slug"`
			Icon       string      `json:"icon"`
			Image      interface{} `json:"image"`
			ImageClass string      `json:"imageClass"`
			BgColor    string      `json:"bgColor"`
			Color      string      `json:"color"`
			Disabled   int         `json:"disabled"`
		} `json:"category"`
		User struct {
			UID                 int    `json:"uid"`
			Username            string `json:"username"`
			Userslug            string `json:"userslug"`
			Reputation          int    `json:"reputation"`
			Postcount           int    `json:"postcount"`
			Picture             string `json:"picture"`
			Signature           string `json:"signature"`
			Banned              int    `json:"banned"`
			Status              string `json:"status"`
			IconText            string `json:"icon:text"`
			IconBgColor         string `json:"icon:bgColor"`
			BannedUntilReadable string `json:"banned_until_readable"`
		} `json:"user,omitempty"`
		Teaser struct {
			Pid          int       `json:"pid"`
			UID          int       `json:"uid"`
			Timestamp    int64     `json:"timestamp"`
			Tid          int       `json:"tid"`
			Content      string    `json:"content"`
			TimestampISO time.Time `json:"timestampISO"`
			User         struct {
				UID         int    `json:"uid"`
				Username    string `json:"username"`
				Userslug    string `json:"userslug"`
				Picture     string `json:"picture"`
				IconText    string `json:"icon:text"`
				IconBgColor string `json:"icon:bgColor"`
			} `json:"user"`
			Index int `json:"index"`
		} `json:"teaser"`
		Tags      []interface{} `json:"tags"`
		IsOwner   bool          `json:"isOwner"`
		Ignored   bool          `json:"ignored"`
		Unread    bool          `json:"unread"`
		Bookmark  int           `json:"bookmark"`
		Unreplied bool          `json:"unreplied"`
		Icons     []interface{} `json:"icons"`
		Index     int           `json:"index"`
		Thumb     string        `json:"thumb,omitempty"`
		TeaserUser      struct {
			UID                 int    `json:"uid"`
			Username            string `json:"username"`
			Fullname            string `json:"fullname"`
			Userslug            string `json:"userslug"`
			Reputation          int    `json:"reputation"`
			Postcount           int    `json:"postcount"`
			Picture             string `json:"picture"`
			Signature           string `json:"signature"`
			Banned              int    `json:"banned"`
			Status              string `json:"status"`
			IconText            string `json:"icon:text"`
			IconBgColor         string `json:"icon:bgColor"`
			BannedUntilReadable string `json:"banned_until_readable"`
		} `json:"user,omitempty"`
	} `json:"topics"`
	Tids       []int `json:"tids"`
	CanPost    bool  `json:"canPost"`
	Categories []struct {
		Cid             int           `json:"cid"`
		Name            string        `json:"name"`
		Slug            string        `json:"slug"`
		Icon            string        `json:"icon"`
		Link            string        `json:"link"`
		Color           string        `json:"color"`
		BgColor         string        `json:"bgColor"`
		ParentCid       int           `json:"parentCid"`
		Image           interface{}   `json:"image"`
		ImageClass      string        `json:"imageClass"`
		Selected        bool          `json:"selected"`
		Children        []interface{} `json:"children"`
		Level           string        `json:"level"`
		BackgroundImage string        `json:"backgroundImage,omitempty"`
	} `json:"categories"`
	AllCategoriesURL string        `json:"allCategoriesUrl"`
	SelectedCids     []interface{} `json:"selectedCids"`
	FeedsDisableRSS  int           `json:"feeds:disableRSS"`
	RssFeedURL       string        `json:"rssFeedUrl"`
	Title            string        `json:"title"`
	Filters          []struct {
		Name     string `json:"name"`
		URL      string `json:"url"`
		Selected bool   `json:"selected"`
		Filter   string `json:"filter"`
	} `json:"filters"`
	SelectedFilter struct {
		Name     string `json:"name"`
		URL      string `json:"url"`
		Selected bool   `json:"selected"`
		Filter   string `json:"filter"`
	} `json:"selectedFilter"`
	Terms []struct {
		Name     string `json:"name"`
		URL      string `json:"url"`
		Selected bool   `json:"selected"`
		Term     string `json:"term"`
	} `json:"terms"`
	SelectedTerm struct {
		Name     string `json:"name"`
		URL      string `json:"url"`
		Selected bool   `json:"selected"`
		Term     string `json:"term"`
	} `json:"selectedTerm"`
	Pagination struct {
		Rel []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"rel"`
		Pages []struct {
			Page      int    `json:"page,omitempty"`
			Active    bool   `json:"active,omitempty"`
			Qs        string `json:"qs,omitempty"`
			Separator bool   `json:"separator,omitempty"`
		} `json:"pages"`
		CurrentPage int `json:"currentPage"`
		PageCount   int `json:"pageCount"`
		Prev        struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"prev"`
		Next struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"next"`
		First struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"first"`
		Last struct {
			Page   int    `json:"page"`
			Active bool   `json:"active"`
			Qs     string `json:"qs"`
		} `json:"last"`
	} `json:"pagination"`
	Breadcrumbs []struct {
		Text string `json:"text"`
		URL  string `json:"url,omitempty"`
	} `json:"breadcrumbs"`
	LoggedIn     bool   `json:"loggedIn"`
	RelativePath string `json:"relative_path"`
	Template     struct {
		Name   string `json:"name"`
		Recent bool   `json:"recent"`
	} `json:"template"`
	URL       string `json:"url"`
	BodyClass string `json:"bodyClass"`
	Widgets   struct {
		Footer []struct {
			HTML string `json:"html"`
		} `json:"footer"`
		Sidebar []struct {
			HTML string `json:"html"`
		} `json:"sidebar"`
	} `json:"widgets"`
}

type Roadster struct {
	Name            string    `json:"name"`
	LaunchDateUtc   time.Time `json:"launch_date_utc"`
	LaunchDateUnix  int       `json:"launch_date_unix"`
	LaunchMassKg    int       `json:"launch_mass_kg"`
	LaunchMassLbs   int       `json:"launch_mass_lbs"`
	NoradID         int       `json:"norad_id"`
	EpochJd         float64   `json:"epoch_jd"`
	OrbitType       string    `json:"orbit_type"`
	ApoapsisAu      float64   `json:"apoapsis_au"`
	PeriapsisAu     float64   `json:"periapsis_au"`
	SemiMajorAxisAu float64   `json:"semi_major_axis_au"`
	Eccentricity    float64   `json:"eccentricity"`
	Inclination     float64   `json:"inclination"`
	Longitude       float64   `json:"longitude"`
	PeriapsisArg    float64   `json:"periapsis_arg"`
	PeriodDays      float64   `json:"period_days"`
	SpeedKph        float64   `json:"speed_kph"`
	SpeedMph        float64   `json:"speed_mph"`
	EarthDistanceKm float64   `json:"earth_distance_km"`
	EarthDistanceMi float64   `json:"earth_distance_mi"`
	MarsDistanceKm  float64   `json:"mars_distance_km"`
	MarsDistanceMi  float64   `json:"mars_distance_mi"`
	FlickrImages    []string  `json:"flickr_images"`
	Wikipedia       string    `json:"wikipedia"`
	Details         string    `json:"details"`
}