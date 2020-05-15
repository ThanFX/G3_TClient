package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	tbot "github.com/Syfaro/telegram-bot-api"
	_ "github.com/lib/pq"
)

// Структура игрока во время активной сессии
type TPlayer struct {
	TUser       TUser
	ChatId      int64
	LastMsgTime time.Time
}

var (
	db       *sql.DB
	TPlayers map[int64]*TPlayer
)

func main() {
	var err error
	local := "host=localhost port=5432 user=postgres password=postgres dbname=g3 sslmode=disable"
	db, err = sql.Open("postgres", local)
	if err != nil {
		log.Fatalf("Ошибка соединения с БД: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ошибка работы с БД: %s", err)
	}
	defer db.Close()

	bot, err := tbot.NewBotAPI("839806396:AAGKWntZYsh4z1ippHIcDVWKVRy0P_ECr2o")
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %s", err)
	}
	log.Printf("Инициализирован бот %s", bot.Self.UserName)
	bot.Debug = true

	TPlayers = make(map[int64]*TPlayer)

	// Делаем постоянные запросы с лонг-полингом к серверу Телеграма, ответы складываем в отдельный канал
	u := tbot.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	// Крутим бесконечный цикл с разбором поступивших ответов
	for upd := range updates {
		// Если нет прямого сообщения нам - скип
		if upd.Message == nil {
			continue
		}
		// Если пишет бот - скип
		if upd.Message.From.IsBot {
			fmt.Println("Боты атакуют!!")
			continue
		}
		// Проверяем пользователя, написавшего сообщение. Если такой пользователь уже есть в игровой сессии - обновляем время последнего сообщения
		user_id := int64(upd.Message.From.ID)
		if _, ok := TPlayers[user_id]; ok {
			TPlayers[user_id].LastMsgTime = time.Now().UTC()
			// если же нету - создаём пользователя и в БД и в мапе игровой сессии
		} else {
			// Ищем такого пользователя в БД
			f, err := isTUserExist(user_id)
			if err != nil {
				fmt.Printf("Ошибка поиска пользователя с id %d: %s", user_id, err)
			}
			// Если нет - создаём
			t_user := TUser{
				UserID:       int64(upd.Message.From.ID),
				UserName:     upd.Message.From.UserName,
				FirstName:    upd.Message.From.FirstName,
				LastName:     upd.Message.From.LastName,
				Lang:         upd.Message.From.LanguageCode,
				CreationDate: time.Now().UTC()}
			if !f {
				err = createTUser(t_user)
				if err != nil {
					fmt.Printf("Ошибка создания пользователя с id %d: %s", user_id, err)
				} else {
					TPlayers[user_id] = &TPlayer{
						TUser:       t_user,
						ChatId:      upd.Message.Chat.ID,
						LastMsgTime: time.Now().UTC()}
				}
			}
		}

		log.Printf("[%s] %s", upd.Message.From.UserName, upd.Message.Text)
	}
}
