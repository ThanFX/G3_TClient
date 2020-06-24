package main

import (
	"fmt"

	tbot "github.com/Syfaro/telegram-bot-api"
)

// Player - структура внутреннего представления игрока
type Player struct {
	TUserID   int64
	MenuState string
	ChanIn    chan UserMsg
	ChunkID   int
}

func createPlayerByID(userID int64) Player {
	p := Player{
		TUserID:   userID,
		MenuState: "base",
		ChanIn:    make(chan UserMsg, 0),
		ChunkID:   Map[getRandInt(0, len(Map)-1)].ID}
	return p
}

// Start - горутина получения всех входящих сообщений от игрока, их обработка и вызов функций ответов
func (p *Player) Start() {
	for {
		userMsg := <-p.ChanIn
		//fmt.Println(userMsg.Msg)
		switch userMsg.Msg {
		case "/start":
			p.sendStateMsg(userMsg.ChatID, "Добро пожаловать в игру!")
		case "/base":
			p.setPlayerMenuState("base")
			p.sendStateMsg(userMsg.ChatID, "С чего начнём?")
		case "/player":
			p.setPlayerMenuState("player")
			p.sendStateMsg(userMsg.ChatID, "Характеристики персонажа:")
		case "/skills":
			p.setPlayerMenuState("skills")
			p.sendStateMsg(userMsg.ChatID, "Навыки и умения:")
		case "/inventory":
			p.setPlayerMenuState("inventory")
			p.sendStateMsg(userMsg.ChatID, "Инвентарь:")
		case "/map":
			p.setPlayerMenuState("map")
			p.sendStateMsg(userMsg.ChatID, p.getMapText())
		case "/show":
			p.setPlayerMenuState("show")
			s := p.getPlayerChunkInfo()
			p.sendStateMsg(userMsg.ChatID, s)
		case "/actions":
			p.setPlayerMenuState("actions")
			p.sendStateMsg(userMsg.ChatID, "Список доступных действий:")
		case "/persons":
			p.setPlayerMenuState("persons")
			p.sendStateMsg(userMsg.ChatID, "Список персонажей:")
		case "/poi":
			p.setPlayerMenuState("poi")
			p.sendStateMsg(userMsg.ChatID, "Интересные места:")
		case "/N":
			if p.movePlayer(0, 1) {
				p.sendCustomMsg("Переход на север выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		case "/NE":
			if p.movePlayer(1, 1) {
				p.sendCustomMsg("Переход на север-восток выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		case "/E":
			if p.movePlayer(1, 0) {
				p.sendCustomMsg("Переход на восток выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		case "/SE":
			if p.movePlayer(1, -1) {
				p.sendCustomMsg("Переход на юго-восток выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		case "/S":
			if p.movePlayer(0, -1) {
				p.sendCustomMsg("Переход на юг выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		case "/SW":
			if p.movePlayer(-1, -1) {
				p.sendCustomMsg("Переход на юго-запад выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		case "/W":
			if p.movePlayer(-1, 0) {
				p.sendCustomMsg("Переход на запад выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		case "/NW":
			if p.movePlayer(-1, 1) {
				p.sendCustomMsg("Переход на северо-запад выполнен")
				p.sendStateMsg(userMsg.ChatID, p.getMapText())
			}
		}
	}
}

func (p *Player) getMapText() string {
	c := Map[p.ChunkID]
	return fmt.Sprintf("Карта местности\nСейчас вы находитесь в локации с ID %d. Координаты (%d, %d)", c.ID, c.X, c.Y)
}

func (p *Player) getPlayerChunkInfo() string {
	c := Map[p.ChunkID]
	var s string
	s = "На локации:\n"
	for _, v := range c.Terrains {
		s1 := fmt.Sprintf("%s занимают %d%% локации\n", v.Name, v.Size*20)
		s += s1
	}
	for _, v := range c.Rivers {
		s1 := fmt.Sprintf("Через локацию протекает река с %s на %s, размер - %d, мост - %t\n", v.From, v.To, v.Size, v.IsBridge)
		s += s1
	}
	return s
}

func (p *Player) movePlayer(x, y int) bool {
	c := Map[p.ChunkID]
	newID := getChunkByCoord(c.X+x, c.Y+y).ID
	if newID != 0 {
		p.ChunkID = newID
		return true
	}
	return false
}

func (p *Player) getPlayerMapDirections() tbot.InlineKeyboardMarkup {
	var k [3][3]tbot.InlineKeyboardButton

	fmt.Println(p.ChunkID)
	nc, err := getNeighborChunkID(p.ChunkID)
	if err != nil {
		fmt.Printf("Пользователь %d находится на несуществующем чанке %d", p.TUserID, p.ChunkID)
		return (tbot.InlineKeyboardMarkup{})
	}
	butNames := [][]string{
		{"↖", "⬆", "↗"},
		{"⬅", "Осмотреться", "➡"},
		{"↙", "⬇", "↘"},
	}
	butVar := [][]string{
		{"/NW", "/N", "/NE"},
		{"/W", "/show", "/E"},
		{"/SW", "/S", "/SE"},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if nc[i][j] != 0 {
				k[i][j] = tbot.NewInlineKeyboardButtonData(butNames[i][j], butVar[i][j])
			} else {
				k[i][j] = tbot.NewInlineKeyboardButtonData("⏹", "/stop")
			}
		}
	}
	keyb := tbot.NewInlineKeyboardMarkup(
		tbot.NewInlineKeyboardRow(k[0][0], k[0][1], k[0][2]),
		tbot.NewInlineKeyboardRow(k[1][0], k[1][1], k[1][2]),
		tbot.NewInlineKeyboardRow(k[2][0], k[2][1], k[2][2]),
		tbot.NewInlineKeyboardRow(tbot.NewInlineKeyboardButtonData("Назад", "/base")),
	)
	return keyb
}

func (p *Player) setPlayerMenuState(state string) {
	p.MenuState = state
}

func (p *Player) getPlayerInlineKeyboard() tbot.InlineKeyboardMarkup {
	var k tbot.InlineKeyboardMarkup

	kPlayer := tbot.NewInlineKeyboardButtonData("Персонаж", "/player")
	kSkills := tbot.NewInlineKeyboardButtonData("Навыки", "/skills")
	kInv := tbot.NewInlineKeyboardButtonData("Инвентарь", "/inventory")
	kMap := tbot.NewInlineKeyboardButtonData("Карта", "/map")
	kBase := tbot.NewInlineKeyboardButtonData("Назад", "/base")
	//kShow := tbot.NewInlineKeyboardButtonData("Осмотреться", "/show")
	kLocShow := tbot.NewInlineKeyboardButtonData("К локации", "/show")
	kLocActions := tbot.NewInlineKeyboardButtonData("Действия", "/actions")
	kLocPersons := tbot.NewInlineKeyboardButtonData("Персонажи", "/persons")
	kLocPoi := tbot.NewInlineKeyboardButtonData("Точки интереса", "/poi")

	switch p.MenuState {
	case "base":
		k = tbot.NewInlineKeyboardMarkup(
			tbot.NewInlineKeyboardRow(kPlayer, kSkills),
			tbot.NewInlineKeyboardRow(kInv, kMap),
		)
	case "player", "skills", "inventory":
		k = tbot.NewInlineKeyboardMarkup(tbot.NewInlineKeyboardRow(kBase))
	case "map":
		k = p.getPlayerMapDirections()
	case "show":
		k = tbot.NewInlineKeyboardMarkup(
			tbot.NewInlineKeyboardRow(kLocActions, kLocPersons),
			tbot.NewInlineKeyboardRow(kLocPoi, kMap),
		)
	case "actions", "persons", "poi":
		k = tbot.NewInlineKeyboardMarkup(tbot.NewInlineKeyboardRow(kLocShow))
	}
	return k
}

func (p *Player) sendStateMsg(chatID int64, text string) {
	replyMsg := tbot.NewMessage(chatID, text)
	//replyMsg.ReplyMarkup = p.getPlayerReplyKeyboard()
	fmt.Println(p.MenuState)
	replyMsg.ReplyMarkup = p.getPlayerInlineKeyboard()
	TBot.Send(replyMsg)
}

func (p *Player) sendCustomMsg(text string) {
	chatID := Users[p.TUserID].ChatID
	replyMsg := tbot.NewMessage(chatID, text)
	TBot.Send(replyMsg)
}

/*
func (p *Player) getPlayerReplyKeyboard() tbot.ReplyKeyboardMarkup {
	var k [][]tbot.KeyboardButton

	kPlayer := tbot.KeyboardButton{Text: "/player"}
	kSkills := tbot.KeyboardButton{Text: "/skills"}
	kInv := tbot.KeyboardButton{Text: "/inventory"}
	kMap := tbot.KeyboardButton{Text: "/map"}
	kBase := tbot.KeyboardButton{Text: "/base"}
	kShow := tbot.KeyboardButton{Text: "/show"}
	kLocActions := tbot.KeyboardButton{Text: "/actions"}
	kLocPersons := tbot.KeyboardButton{Text: "/persons"}
	kLocPoi := tbot.KeyboardButton{Text: "/poi"}

	switch p.MenuState {
	case "base":
		k = [][]tbot.KeyboardButton{
			{kPlayer, kSkills},
			{kInv, kMap},
		}
	case "player", "skills", "inventory":
		k = [][]tbot.KeyboardButton{{kBase}}
	case "map":
		k = [][]tbot.KeyboardButton{{kShow}, {kBase}}
	case "show":
		k = [][]tbot.KeyboardButton{
			{kLocActions, kLocPersons},
			{kLocPoi, kMap},
		}
	case "actions", "persons", "poi":
		k = [][]tbot.KeyboardButton{{kShow}}
	}

	rk := tbot.ReplyKeyboardMarkup{
		Keyboard:        k,
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
		Selective:       false,
	}
	return rk
}
*/
