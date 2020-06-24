package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const dbGetFullMap = `SELECT mc.id, mc.x, mc.y, json_agg(json_build_object('type', mt."type", 'name', mt."name", 'size', mct."size"))
						FROM map_chunk mc
						join map_chunk_terrains mct on mct.chunk_id = mc.id
						join map_terrain mt on mt.id = mct.terrain_id
						group by 1, 2, 3;`
const dbGetAllRivews = `select chunk_id, "from", "to", "size", is_bridge from map_chunk_rivers;`

// ChunkTerrain - структура данных о местности в чанке
type ChunkTerrain struct {
	Type string
	Name string
	Size int
}

// MapRiver - структура данных о реке в чанке
type MapRiver struct {
	From     string
	To       string
	Size     int
	IsBridge bool
}

// MapChunk - структура позиционирования чанка на глобальной карте
type MapChunk struct {
	ID       int
	X        int
	Y        int
	Terrains []ChunkTerrain
	Rivers   []MapRiver
}

// Mastership - структура данных о профессии
type Mastership struct {
	Type string
	Name string
}

// Map - массив всех существующих чанков
var (
	Map            map[int]MapChunk        = make(map[int]MapChunk)
	TerrainMastery map[string][]Mastership = map[string][]Mastership{
		"forest": {
			{Type: "hunting", Name: "Охота"},
			{Type: "food_gathering", Name: "Собирательство грибов и ягод"},
		},
		"meadow": {
			{Type: "hunting", Name: "Охота"},
			{Type: "food_gathering", Name: "Собирательство грибов и ягод"},
		},
		"hill": {
			{Type: "hunting", Name: "Охота"},
			{Type: "food_gathering", Name: "Собирательство грибов и ягод"},
		},
		"swamp": {
			{Type: "hunting", Name: "Охота"},
			{Type: "food_gathering", Name: "Собирательство грибов и ягод"},
		},
		"lake": {
			{Type: "fishing", Name: "Рыбная ловля"},
			{Type: "food_gathering", Name: "Собирательство грибов и ягод"},
		},
		"river": {
			{Type: "hunting", Name: "Охота"},
			{Type: "fishing", Name: "Рыбная ловля"},
			{Type: "food_gathering", Name: "Собирательство грибов и ягод"},
		},
	}
)

// ReadMapFromDB - получаем все существующие чанки и сохраняем их в переменной Map
func ReadMapFromDB() {
	var mc MapChunk
	rows, err := db.Query(dbGetFullMap)
	if err != nil {
		log.Fatalf("Ошибка получения карты из БД: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		err = rows.Scan(&mc.ID, &mc.X, &mc.Y, &s)
		if err != nil {
			log.Fatal("ошибка получения чанка: ", err)
		}

		var cts []ChunkTerrain
		err = json.Unmarshal([]byte(s), &cts)
		if err != nil {
			log.Fatal("Ошибка парсинга местностей чанка: ", err)
		}
		mc.Terrains = cts
		Map[mc.ID] = mc
	}

	rivers := getRiversFormDB()
	for i, v := range rivers {
		c := Map[i]
		c.Rivers = v
		Map[i] = c
	}
	//fmt.Println(Map)
}

func getRiversFormDB() map[int][]MapRiver {
	var r map[int][]MapRiver = make(map[int][]MapRiver)
	rows, err := db.Query(dbGetAllRivews)
	if err != nil {
		log.Fatalf("Ошибка получения рек из БД: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var mr MapRiver
		err = rows.Scan(&id, &mr.From, &mr.To, &mr.Size, &mr.IsBridge)
		if err != nil {
			log.Fatal("Ошибка парсинга данных о реке: ", err)
		}
		r[id] = append(r[id], mr)
	}
	return r
}

func getNeighborChunkID(id int) ([3][3]int, error) {
	var nc [3][3]int
	c := Map[id]
	if c.ID == 0 {
		return nc, fmt.Errorf("Не найден чанк с ID = %d", id)
	}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			nc[1-j][i+1] = getChunkByCoord(c.X+i, c.Y+j).ID
		}
	}

	return nc, nil
}

func getChunkByCoord(x, y int) MapChunk {
	var c MapChunk
	for i := range Map {
		if x == Map[i].X && y == Map[i].Y {
			c = Map[i]
			break
		}
	}
	return c
}
