package gohq

import "time"

// Account information
type Account struct {
	UserID      int    `json:"userId"`
	Username    string `json:"username"`
	Admin       bool   `json:"admin"`
	Tester      bool   `json:"tester"`
	Guest       bool   `json:"guest"`
	AvatarURL   string `json:"avatarUrl"`
	LoginToken  string `json:"loginToken"`
	AccessToken string `json:"accessToken"`
	AuthToken   string `json:"authToken"`
}

// HQ specific errors
type HQError struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
}

// HQ token refresh data
type Tokens struct {
	UserID      int    `json:"userId"`
	Username    string `json:"username"`
	Admin       bool   `json:"admin"`
	Tester      bool   `json:"tester"`
	Guest       bool   `json:"guest"`
	AvatarURL   string `json:"avatarUrl"`
	LoginToken  string `json:"loginToken"`
	AccessToken string `json:"accessToken"`
	AuthToken   string `json:"authToken"`
}

// Me contains profile information
type Me struct {
	UserID    int       `json:"userId"`
	Username  string    `json:"username"`
	AvatarURL string    `json:"avatarUrl"`
	Created   time.Time `json:"created"`
	Broadcasts struct {
		Data []interface{} `json:"data"`
	} `json:"broadcasts"`
	Featured        bool     `json:"featured"`
	Voip            bool     `json:"voip"`
	DeviceTokens    []string `json:"deviceTokens"`
	HasPhone        bool     `json:"hasPhone"`
	PhoneNumber     string   `json:"phoneNumber"`
	ReferralURL     string   `json:"referralUrl"`
	Referred        bool     `json:"referred"`
	ReferringUserID int      `json:"referringUserId"`
	HighScore       int      `json:"highScore"`
	GamesPlayed     int      `json:"gamesPlayed"`
	WinCount        int      `json:"winCount"`
	Blocked         bool     `json:"blocked"`
	BlocksMe        bool     `json:"blocksMe"`
	Preferences struct {
		SharingEnabled bool `json:"sharingEnabled"`
	} `json:"preferences"`
	FriendIds []int  `json:"friendIds"`
	Lives     string `json:"lives"`
	Stk       string `json:"stk"`
	Leaderboard struct {
		TotalCents int    `json:"totalCents"`
		Total      string `json:"total"`
		Unclaimed  string `json:"unclaimed"`
		Wins       int    `json:"wins"`
		Rank       int    `json:"rank"`
		Alltime struct {
			Total string `json:"total"`
			Wins  int    `json:"wins"`
			Rank  int    `json:"rank"`
		} `json:"alltime"`
		Weekly struct {
			Total string `json:"total"`
			Wins  int    `json:"wins"`
			Rank  int    `json:"rank"`
		} `json:"weekly"`
	} `json:"leaderboard"`
}

// CashoutData is data you get after cashing out
type CashoutData struct {
	Data struct {
		PayoutID     int         `json:"payoutId"`
		UserID       int         `json:"userId"`
		Amount       string      `json:"amount"`
		Currency     string      `json:"currency"`
		TargetUserID interface{} `json:"targetUserId"`
		TargetEmail  string      `json:"targetEmail"`
		TargetPhone  interface{} `json:"targetPhone"`
		Status       int         `json:"status"`
		Metadata struct {
			PayoutsConnected bool   `json:"payoutsConnected"`
			Client           string `json:"client"`
		} `json:"metadata"`
		Created  time.Time `json:"created"`
		Modified time.Time `json:"modified"`
	} `json:"data"`
}

// PayoutData is data of all payouts
type PayoutData struct {
	Payouts []struct {
		PayoutID     int         `json:"payoutId"`
		UserID       int         `json:"userId"`
		Amount       string      `json:"amount"`
		Currency     string      `json:"currency"`
		TargetUserID interface{} `json:"targetUserId"`
		TargetEmail  string      `json:"targetEmail"`
		TargetPhone  interface{} `json:"targetPhone"`
		Status       int         `json:"status"`
		Metadata struct {
			PayoutsConnected bool   `json:"payoutsConnected"`
			Client           string `json:"client"`
		} `json:"metadata"`
		Created  time.Time `json:"created"`
		Modified time.Time `json:"modified"`
	} `json:"payouts"`
	Balance struct {
		PrizeTotal        string `json:"prizeTotal"`
		Paid              string `json:"paid"`
		Pending           string `json:"pending"`
		Unpaid            string `json:"unpaid"`
		EligibleForPayout bool   `json:"eligibleForPayout"`
		HasPending        bool   `json:"hasPending"`
		PayoutsConnected  bool   `json:"payoutsConnected"`
		PayoutsEmail      string `json:"payoutsEmail"`
		DocumentRequired  bool   `json:"documentRequired"`
		DocumentStatus    string `json:"documentStatus"`
	} `json:"balance"`
}

// SearchData is data from when a user is searched for
type SearchData struct {
	Data []struct {
		UserID          int         `json:"userId,int"`
		Username        string      `json:"username"`
		AvatarURL       string      `json:"avatarUrl,omitempty"`
		Created         time.Time   `json:"created"`
		Live            bool        `json:"live,omitempty"`
		SubscriberCount int         `json:"subscriberCount,omitempty"`
		LastLive        interface{} `json:"lastLive,omitempty"`
		Featured        bool        `json:"featured,omitempty"`
	} `json:"data"`
	Links struct {
		Next interface{} `json:"next"`
		Prev string      `json:"prev"`
		Self string      `json:"self"`
	} `json:"links"`
}