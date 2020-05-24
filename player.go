package main

import (
	"fmt"

	tbot "github.com/Syfaro/telegram-bot-api"
	uuid "github.com/satori/go.uuid"
)

type Player struct {
	TUserID   int64
	MenuState string
	ChanIn    chan UserMsg
	ChunkID   uuid.UUID
}

func createPlayerById(user_id int64) Player {
	p := Player{
		TUserID:   user_id,
		MenuState: "base",
		ChanIn:    make(chan UserMsg, 0),
		ChunkID:   Map[getRandInt(0, len(Map)-1)].ID}
	return p
}

func (p *Player) Start() {
	for {
		user_msg := <-p.ChanIn
		fmt.Println(user_msg.Msg)
		switch user_msg.Msg {
		case "/start":
			p.sendMsg(user_msg.ChatId, "Добро пожаловать в игру!")
		case "/base":
			p.setPlayerMenuState("base")
			p.sendMsg(user_msg.ChatId, "С чего начнём?")
		case "/player":
			p.setPlayerMenuState("player")
			p.sendMsg(user_msg.ChatId, "Характеристики персонажа:")
		case "/skills":
			p.setPlayerMenuState("skills")
			p.sendMsg(user_msg.ChatId, "Навыки и умения:")
		case "/inventory":
			p.setPlayerMenuState("inventory")
			p.sendMsg(user_msg.ChatId, "Инвентарь:")
		case "/map":
			p.setPlayerMenuState("map")
			p.sendMsg(user_msg.ChatId, "Карта местности:")
		case "/show":
			p.setPlayerMenuState("show")
			p.getPlayerChunkInfo()
			p.sendMsg(user_msg.ChatId, "Вы осмотрелись и увидели много интересного:")
		case "/actions":
			p.setPlayerMenuState("actions")
			p.sendMsg(user_msg.ChatId, "Список доступных действий:")
		case "/persons":
			p.setPlayerMenuState("persons")
			p.sendMsg(user_msg.ChatId, "Список персонажей:")
		case "/poi":
			p.setPlayerMenuState("poi")
			p.sendMsg(user_msg.ChatId, "Интересные места:")
		}
	}
}

func (p *Player) getPlayerChunkInfo() {
	ct := GetChunkTerrainsInfo(p.ChunkID)
	cm := GetChunckAreasMastery(p.ChunkID)
	fmt.Println(ct)
	fmt.Println(cm)
}

func (p *Player) setPlayerMenuState(state string) {
	p.MenuState = state
}

func (p *Player) getPlayerInlineKeyboard() tbot.InlineKeyboardMarkup {
	var k tbot.InlineKeyboardMarkup
	switch p.MenuState {
	case "base":
		k = tbot.NewInlineKeyboardMarkup(
			tbot.NewInlineKeyboardRow(
				tbot.NewInlineKeyboardButtonData("Персонаж", "/player"),
				tbot.NewInlineKeyboardButtonData("Навыки", "/skills"),
			),
			tbot.NewInlineKeyboardRow(
				tbot.NewInlineKeyboardButtonData("Инвентарь", "/inventory"),
				tbot.NewInlineKeyboardButtonData("Карта", "/map"),
			),
		)
	case "player", "skills", "inventory":
		k = tbot.NewInlineKeyboardMarkup(tbot.NewInlineKeyboardRow(tbot.NewInlineKeyboardButtonData("Назад", "/base")))
	case "map":
		k = tbot.NewInlineKeyboardMarkup(
			tbot.NewInlineKeyboardRow(
				tbot.NewInlineKeyboardButtonData("Осмотреться", "/show"),
			),
			tbot.NewInlineKeyboardRow(
				tbot.NewInlineKeyboardButtonData("Назад", "/base"),
			),
		)
	case "show":
		k = tbot.NewInlineKeyboardMarkup(
			tbot.NewInlineKeyboardRow(
				tbot.NewInlineKeyboardButtonData("Действия", "/actions"),
				tbot.NewInlineKeyboardButtonData("Персонажи", "/persons"),
			),
			tbot.NewInlineKeyboardRow(
				tbot.NewInlineKeyboardButtonData("Точки интереса", "/poi"),
				tbot.NewInlineKeyboardButtonData("Карта", "/map"),
			),
		)
	case "actions", "persons", "poi":
		k = tbot.NewInlineKeyboardMarkup(tbot.NewInlineKeyboardRow(tbot.NewInlineKeyboardButtonData("Назад", "/show")))
	}
	return k
}

func (p *Player) sendMsg(chat_id int64, text string) {
	reply_msg := tbot.NewMessage(chat_id, text)
	//fmt.Println(reply_msg)
	reply_msg.ReplyMarkup = p.getPlayerInlineKeyboard()
	//fmt.Println(reply_msg)
	TBot.Send(reply_msg)
}
