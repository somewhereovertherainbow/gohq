package gohq

import (
	"net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
)

// Connect to a HQ websocket instance using a game id
func (a *Account) Connect(gID string) (game *Game, err error) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+a.AccessToken)

	c, _, err := websocket.DefaultDialer.Dial("wss://ws-quiz.hype.space/ws/"+gID, headers)
	if err != nil {
		return
	}

	game = &Game{Conn: c}
	return
}

// Close a connection (wrapper)
func (g *Game) Close() error {
	return g.Conn.Close()
}

// Ping a connection
func (g *Game) Ping() error {
	return g.Conn.WriteMessage(websocket.PingMessage, nil)
}

// Subscribe to a game
func (g *Game) Subscribe(gID string) error {
	type Data struct {
		Type        string `json:"type"`
		BroadcastID string `json:"broadcastId"`
	}

	bytes, err := json.Marshal(Data{Type: "subscribe", BroadcastID: gID})
	if err != nil {
		return err
	}

	return g.Conn.WriteMessage(websocket.TextMessage, bytes)
}

// Answer a question
func (g *Game) Answer(aID, qID, gID string) error {
	type Data struct {
		Type        string `json:"type"`
		BroadcastID string `json:"broadcastId"`
		AnswerID    string `json:"answerId"`
		QuestionID  string `json:"questionId"`
	}

	bytes, err := json.Marshal(Data{Type: "answer", BroadcastID: gID, QuestionID: qID, AnswerID: aID})
	if err != nil {
		return err
	}

	return g.Conn.WriteMessage(websocket.TextMessage, bytes)
}

// Use a life on an a question
func (g *Game) Life(qID, gID string) error {
	type Data struct {
		Type        string `json:"type"`
		BroadcastID string `json:"broadcastId"`
		QuestionID  string `json:"questionId"`
	}

	bytes, err := json.Marshal(Data{Type: "useExtraLife", BroadcastID: gID, QuestionID: qID})
	if err != nil {
		return err
	}

	return g.Conn.WriteMessage(websocket.TextMessage, bytes)
}

// Parse broadcast statistics
func (g *Game) ParseBroadcastStats(bytes []byte) (stats *BroadcastStats) {
	json.Unmarshal(bytes, &stats)

	if stats.Type == "broadcastStats" {
		return stats
	}

	return
}

// Parse a question
func (g *Game) ParseQuestion(bytes []byte) (question *Question) {
	json.Unmarshal(bytes, &question)

	if question.Type == "question" && len(question.Answers) != 0 {
		return question
	}

	return
}

// Parse a question summary
func (g *Game) ParseQuestionSummary(bytes []byte) (questionSummary *QuestionSummary) {
	json.Unmarshal(bytes, &questionSummary)

	if questionSummary.Type == "questionSummary" {
		return questionSummary
	}

	return
}

// Parse a question closure
func (g *Game) ParseQuestionClosed(bytes []byte) (questionClosed *QuestionClosed) {
	json.Unmarshal(bytes, &questionClosed)

	if questionClosed.Type == "questionClosed" {
		return questionClosed
	}

	return
}

// Parse when the question finishes
func (g *Game) ParseQuestionFinished(bytes []byte) (questionFinished *QuestionFinished) {
	json.Unmarshal(bytes, &questionFinished)

	if questionFinished.Type == "questionFinished" {
		return questionFinished
	}

	return
}

// Parse a chat message
func (g *Game) ParseChatMessage(bytes []byte) (chatMessage *ChatMessage) {
	json.Unmarshal(bytes, &chatMessage)

	if chatMessage.Type == "interaction" && chatMessage.ItemID == "chat" {
		return chatMessage
	}

	return
}

// Parse a game status
func (g *Game) ParseGameStatus(bytes []byte) (gameStatus *GameStatus) {
	json.Unmarshal(bytes, &gameStatus)

	if gameStatus.Type == "gameStatus" {
		return gameStatus
	}

	return
}
