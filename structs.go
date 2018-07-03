package gohq

import (
	"time"
	"github.com/gorilla/websocket"
)

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

// ScheduleData contains information about the upcoming and ongoing HQ games
type ScheduleData struct {
	Active        bool      `json:"active"`
	AtCapacity    bool      `json:"atCapacity"`
	ShowID        int       `json:"showId"`
	ShowType      string    `json:"showType"`
	StartTime     time.Time `json:"startTime"`
	NextShowTime  time.Time `json:"nextShowTime"`
	NextShowPrize string    `json:"nextShowPrize"`
	Upcoming []struct {
		Time  time.Time `json:"time"`
		Prize string    `json:"prize"`
	} `json:"upcoming"`
	Prize int `json:"prize"`
	Broadcast struct {
		BroadcastID   int           `json:"broadcastId"`
		UserID        int           `json:"userId"`
		Title         string        `json:"title"`
		Status        int           `json:"status"`
		State         string        `json:"state"`
		ChannelID     int           `json:"channelId"`
		Created       time.Time     `json:"created"`
		Started       time.Time     `json:"started"`
		Ended         interface{}   `json:"ended"`
		Permalink     string        `json:"permalink"`
		ThumbnailData interface{}   `json:"thumbnailData"`
		Tags          []interface{} `json:"tags"`
		SocketURL     string        `json:"socketUrl"`
		Streams struct {
			Source      string `json:"source"`
			Passthrough string `json:"passthrough"`
			High        string `json:"high"`
			Medium      string `json:"medium"`
			Low         string `json:"low"`
		} `json:"streams"`
		StreamURL         string `json:"streamUrl"`
		StreamKey         string `json:"streamKey"`
		RelativeTimestamp int    `json:"relativeTimestamp"`
		Links struct {
			Self       string `json:"self"`
			Transcript string `json:"transcript"`
			Viewers    string `json:"viewers"`
		} `json:"links"`
	} `json:"broadcast"`
	GameKey       string `json:"gameKey"`
	BroadcastFull bool   `json:"broadcastFull"`
}

// Game is a type which contains information about websocket connections
type Game struct {
	Conn *websocket.Conn
}

// BroadcastStats gives you information about the current broadcast
type BroadcastStats struct {
	Type          string `json:"type"`
	LikeCount     int    `json:"likeCount"`
	StatusMessage string `json:"statusMessage"`
	ViewerCounts struct {
		Connected int `json:"connected"`
		Playing   int `json:"playing"`
		Watching  int `json:"watching"`
	} `json:"viewerCounts"`
	BroadcastSubscribers []interface{} `json:"broadcastSubscribers"`
	Ts                   time.Time     `json:"ts"`
	Sent                 time.Time     `json:"sent"`
}

// Question gives you information about a new question
type Question struct {
	Type        string `json:"type"`
	TotalTimeMs int    `json:"totalTimeMs"`
	TimeLeftMs  int    `json:"timeLeftMs"`
	QuestionID  int    `json:"questionId"`
	Question    string `json:"question"`
	Category    string `json:"category"`
	Answers []struct {
		AnswerID int    `json:"answerId"`
		Text     string `json:"text"`
	} `json:"answers"`
	QuestionNumber int       `json:"questionNumber"`
	QuestionCount  int       `json:"questionCount"`
	Ts             time.Time `json:"ts"`
	Sent           time.Time `json:"sent"`
}

// QuestionSummary returns post question information
type QuestionSummary struct {
	AdvancingPlayersCount int `json:"advancingPlayersCount"`
	AnswerCounts []struct {
		Answer   string `json:"answer"`
		AnswerID int    `json:"answerId"`
		Correct  bool   `json:"correct"`
		Count    int    `json:"count"`
	} `json:"answerCounts"`
	EliminatedPlayersCount int       `json:"eliminatedPlayersCount"`
	ExtraLivesRemaining    int       `json:"extraLivesRemaining"`
	ID                     string    `json:"id"`
	Question               string    `json:"question"`
	QuestionID             int       `json:"questionId"`
	SavedByExtraLife       bool      `json:"savedByExtraLife"`
	Sent                   time.Time `json:"sent"`
	Ts                     time.Time `json:"ts"`
	Type                   string    `json:"type"`
	YouGotItRight          bool      `json:"youGotItRight"`
	YourAnswerID           int       `json:"yourAnswerId"`
}

// QuestionClosed indicates a question is no longer able to be answered
type QuestionClosed struct {
	Type       string    `json:"type"`
	QuestionID int       `json:"questionId"`
	Ts         time.Time `json:"ts"`
	Sent       time.Time `json:"sent"`
}

// Question finished indicates the question is now over
type QuestionFinished struct {
	Type       string    `json:"type"`
	QuestionID int       `json:"questionId"`
	Ts         time.Time `json:"ts"`
	Sent       time.Time `json:"sent"`
}

// GameStatus indicates status information for the game
type GameStatus struct {
	CardPlaysRemaining  int         `json:"cardPlaysRemaining"`
	Kind                string      `json:"kind"`
	Prize               string      `json:"prize"`
	InTheGame           bool        `json:"inTheGame"`
	Type                string      `json:"type"`
	QuestionCount       int         `json:"questionCount"`
	ExtraLivesRemaining int         `json:"extraLivesRemaining"`
	CurrentState        interface{} `json:"currentState"`
	Cts                 time.Time   `json:"cts"`
	QuestionNumber      int         `json:"questionNumber"`
	ExtraLives          int         `json:"extraLives"`
	Ts                  time.Time   `json:"ts"`
	Sent                time.Time   `json:"sent"`
}

// ChatMessage shows chat messages sent by users
type ChatMessage struct {
	Type   string `json:"type"`
	ItemID string `json:"itemId"`
	UserID int    `json:"userId"`
	Metadata struct {
		UserID      int    `json:"userId"`
		Message     string `json:"message"`
		AvatarURL   string `json:"avatarUrl"`
		Interaction string `json:"interaction"`
		Username    string `json:"username"`
	} `json:"metadata"`
	Ts   time.Time `json:"ts"`
	Sent time.Time `json:"sent"`
}

// Authentication information from logging in?
type Auth struct {
	Auth *Account `json:"auth,omitempty"`
}

// AWSSession gives information for uploading avatars
type AWSSession struct {
	AccessKeyID  string    `json:"accessKeyId"`
	SecretKey    string    `json:"secretKey"`
	SessionToken string    `json:"sessionToken"`
	Expiration   time.Time `json:"expiration"`
}

// A verification session
type Verification struct {
	CallsEnabled   bool      `json:"callsEnabled"`
	Expires        time.Time `json:"expires"`
	Phone          string    `json:"phone"`
	RetrySeconds   int       `json:"retrySeconds"`
	VerificationID string    `json:"verificationId"`
}

// The result to an avatar change
type AvatarChange struct {
	UserID    int       `json:"userId"`
	Username  string    `json:"username"`
	AvatarURL string    `json:"avatarUrl"`
	Created   time.Time `json:"created"`
}
