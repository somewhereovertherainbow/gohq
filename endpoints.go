package gohq

import (
	"strconv"
	"net/url"
)

var (
	EndpointBase = "https://api-quiz.hype.space/"

	EndpointUsers    = EndpointBase + "users/"
	EndpointMe       = EndpointUsers + "me/"
	EndpointPayouts  = EndpointMe + "payouts/"
	EndpointShows    = EndpointBase + "shows/"
	EndpointSchedule = EndpointShows + "now?type=hq"
	EndpointFriends = EndpointBase + "friends/"
	EndpointEasterEggs = EndpointBase + "easter-eggs/"
	EndpointMakeItRain = EndpointEasterEggs + "makeItRain/"
	EndpointTokens   = EndpointBase + "tokens/"

	EndpointUser = func(uID int) string { return EndpointUsers + strconv.Itoa(uID) + "/" }
	EndpointFriend = func(uID int) string {return EndpointFriends + strconv.Itoa(uID) + "/"}
	EndpointFriendRequest = func(uID int) string {return EndpointFriend(uID) + "/requests/"}
	EndpointSearchUser = func (query string) string { return EndpointUsers + "?q=" + url.QueryEscape(query) }
)
