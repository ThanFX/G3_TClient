package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	tbot "github.com/Syfaro/telegram-bot-api"
	_ "github.com/lib/pq"
)

// User - структура игрока во время активной сессии
type User struct {
	TUser       TUser
	ChatID      int64
	LastMsgTime time.Time
	Player      Player
}

// UserMsg - структура хранения пришедшего сообщения для каждого игрока
type UserMsg struct {
	ChatID int64
	Msg    string
}

// Users - мапа всех активных игроков в текущей сессии. Ключ - UserID
var (
	db    *sql.DB
	Users map[int64]*User
	TBot  *tbot.BotAPI
)

func getRandInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min+1) + min
}

func main() {
	var err error
	local := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", local)
	if err != nil {
		log.Fatalf("Ошибка соединения с БД: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ошибка работы с БД: %s", err)
	}
	defer db.Close()

	// Читаем и инициализируем карту
	ReadMapFromDB()

	TBot, err = tbot.NewBotAPI("839806396:AAGKWntZYsh4z1ippHIcDVWKVRy0P_ECr2o")
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %s", err)
	}
	log.Printf("Инициализирован бот %s", TBot.Self.UserName)
	//TBot.Debug = true

	Users = make(map[int64]*User)

	// Делаем постоянные запросы с лонг-полингом к серверу Телеграма, ответы складываем в отдельный канал
	u := tbot.NewUpdate(0)
	u.Timeout = 60
	updates, err := TBot.GetUpdatesChan(u)

	// Крутим бесконечный цикл с разбором поступивших ответов
	for upd := range updates {
		// Если нет прямого сообщения нам - скип
		if upd.Message == nil && upd.CallbackQuery.Message == nil {
			continue
		}

		var usr *tbot.User
		var um UserMsg
		// Сообщение может придти напрямую от пользователя, а может от самого бота через инлайн-кнопки. Проверяем оба вариант, собираем данные пользователя, чат и сообщение
		if upd.Message == nil {
			usr = upd.CallbackQuery.From
			um = UserMsg{
				ChatID: upd.CallbackQuery.Message.Chat.ID,
				Msg:    upd.CallbackQuery.Data}
		} else {
			usr = upd.Message.From
			um = UserMsg{
				ChatID: upd.Message.Chat.ID,
				Msg:    upd.Message.Text}
		}

		//fmt.Println(um)

		// Если пишет бот - скип
		if usr.IsBot {
			fmt.Println("Боты атакуют!!")
			continue
		}

		// Проверяем пользователя, написавшего сообщение. Если такой пользователь уже есть в игровой сессии - обновляем время последнего сообщения
		userID := int64(usr.ID)
		fmt.Printf("%s: %s\n", usr.UserName, um.Msg)
		if _, ok := Users[userID]; ok {
			//fmt.Println("Пользователь есть в мапе")
			Users[userID].LastMsgTime = time.Now().UTC()
			// передаём пользователю пришедшее сообщение
			Users[userID].Player.ChanIn <- um
			// если же нету - создаём пользователя и в БД и в мапе игровой сессии
		} else {
			//fmt.Println("Пользователя нет в мапе")
			// Ищем такого пользователя в БД
			tu, err := getTUserByID(userID)
			if err != nil {
				fmt.Printf("Ошибка поиска пользователя с id %d: %s", userID, err)
			}

			// Если не нашли пользователя - создаём в БД
			if tu.UserID == 0 {
				fmt.Println("Пользователя нет в БД")
				tu = TUser{
					UserID:       int64(usr.ID),
					UserName:     usr.UserName,
					FirstName:    usr.FirstName,
					LastName:     usr.LastName,
					Lang:         usr.LanguageCode,
					CreationDate: time.Now().UTC()}
				err = createTUser(tu)
				if err != nil {
					fmt.Printf("Ошибка создания пользователя с id %d: %s", userID, err)
					continue
				}
			}

			//fmt.Println(tu)
			// Создаём пользователя в мапе игровой сессии
			Users[userID] = &User{
				TUser:       tu,
				ChatID:      um.ChatID,
				LastMsgTime: time.Now().UTC(),
				Player:      createPlayerByID(userID)}

			//fmt.Println(Users[userID].Player.ChunkID)
			// Запускаем горутину пользователя
			go Users[userID].Player.Start()
			// передаём в горутину пользователя пришедшее сообщение
			Users[userID].Player.ChanIn <- um
		}
		//log.Printf("[%s] %s", msg.From.UserName, msg.Text)
	}
}
