package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ThanFX/G3/areas"
	uuid "github.com/satori/go.uuid"
)

type Chunk struct {
	ID       uuid.UUID      `json:"id"`
	X        int            `json:"x"`
	Y        int            `json:"y"`
	Terrains []TerrainChunk `json:"terrains"`
	Lakes    []LakeChunk    `json:"lakes"`
	Rivers   []RiverChunk   `json:"rivers"`
}

type TerrainChunk struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

type LakeChunk struct {
	Size int `json:"size"`
}

type RiverChunk struct {
	Size   int    `json:"size"`
	From   string `json:"from"`
	To     string `json:"to"`
	Bridge bool   `json:"bridge"`
}

type AreaMastery struct {
	Name   string
	AreaID uuid.UUID
	Size   int
}

type AreaInfo struct {
	Forest areas.Forest
	Hill   areas.Hill
	Swamp  areas.Swamp
	Meadow areas.Meadow
	Lake   areas.Lake
	Rivers []areas.River
}

var (
	Map              []Chunk
	ChunkMasteryInfo map[uuid.UUID]map[string][]AreaMastery
	ChunkAreasInfo   map[uuid.UUID]map[string]interface{}
	ChunkAreasInfoEx map[uuid.UUID]AreaInfo
)

func ReadMapFromDB() {
	rows, err := db.Query("select chunk from map")
	if err != nil {
		log.Fatalf("Ошибка получения карты из БД: %s", err)
	}
	defer rows.Close()

	var ch string
	var chunk Chunk
	for rows.Next() {
		err = rows.Scan(&ch)
		if err != nil {
			log.Fatal("ошибка получения записи чанка: ", err)
		}
		err = json.Unmarshal([]byte(ch), &chunk)
		if err != nil {
			log.Fatal("ошибка парсинга данных чанка: ", err)
		}
		Map = append(Map, chunk)

		chunk.Terrains = nil
		chunk.Rivers = nil
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	createTerrains()
	fmt.Println("Карта успешно загружена и инициализирована")
}

func getChunkByID(id uuid.UUID) Chunk {
	var c Chunk
	for i := range Map {
		if uuid.Equal(id, Map[i].ID) {
			c = Map[i]
			break
		}
	}
	return c
}

func getChunkByCoord(x, y int) Chunk {
	var c Chunk
	for i := range Map {
		if x == Map[i].X && y == Map[i].Y {
			c = Map[i]
			break
		}
	}
	return c
}

func GetChunkTerrainsInfo(chunkId uuid.UUID) map[string]interface{} {
	return ChunkAreasInfo[chunkId]
}

func GetChunckAreasMastery(chunkId uuid.UUID) map[string][]AreaMastery {
	return ChunkMasteryInfo[chunkId]
}

func createTerrains() {
	ChunkMasteryInfo = make(map[uuid.UUID]map[string][]AreaMastery)
	ChunkAreasInfo = make(map[uuid.UUID]map[string]interface{})
	ChunkAreasInfoEx = make(map[uuid.UUID]AreaInfo)
	for _, m := range Map {
		var ai AreaInfo
		ChunkAreasInfo[m.ID] = make(map[string]interface{})
		for _, t := range m.Terrains {
			switch t.Type {
			case "forest":
				id := areas.CreateForest(m.ID, t.Size)
				ai.Forest = areas.GetForestById(id)
				ChunkAreasInfo[m.ID]["forest"] = ai.Forest
			case "hill":
				id := areas.CreateHill(m.ID, t.Size)
				ai.Hill = areas.GetHillsById(id)
				ChunkAreasInfo[m.ID]["hill"] = ai.Hill
			case "swamp":
				id := areas.CreateSwamp(m.ID, t.Size)
				ai.Swamp = areas.GetSwampById(id)
				ChunkAreasInfo[m.ID]["swamp"] = ai.Swamp
			case "meadow":
				id := areas.CreateMeadow(m.ID, t.Size)
				ai.Meadow = areas.GetMeadowById(id)
				ChunkAreasInfo[m.ID]["meadow"] = ai.Meadow
			case "lake":
				id := areas.CreateLake(m.ID, t.Size)
				ai.Lake = areas.GetLakesById(id)
				ChunkAreasInfo[m.ID]["lakes"] = ai.Lake
			}
		}
		var rs []areas.River
		for _, r := range m.Rivers {
			id := areas.CreateRiver(m.ID, r.Size, r.Bridge)
			rs = append(rs, areas.GetRiversById(id)[0])
		}
		if len(rs) > 0 {
			ai.Rivers = rs
			ChunkAreasInfo[m.ID]["rivers"] = rs
		}

		/*
			-- Озера пока не в отдельной структуре идут
			var ls []areas.Lake
			for _, l := range m.Lakes {
				id := areas.CreateLake(m.ID, l.Size)
				ls = append(ls, areas.GetLakesById(id)[0])
			}
			if len(ls) > 0 {
				ChunkAreasInfo[m.ID]["lakes"] = ls
			}
		*/

		ChunkAreasInfoEx[m.ID] = ai

		ChunkMasteryInfo[m.ID] = make(map[string][]AreaMastery)
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Forest.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Forest.Masterships {
				var am AreaMastery
				am.Name = "forest"
				am.Size = ChunkAreasInfoEx[m.ID].Forest.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Forest.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Hill.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Hill.Masterships {
				var am AreaMastery
				am.Name = "hill"
				am.Size = ChunkAreasInfoEx[m.ID].Hill.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Hill.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Swamp.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Swamp.Masterships {
				var am AreaMastery
				am.Name = "swamp"
				am.Size = ChunkAreasInfoEx[m.ID].Swamp.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Swamp.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Meadow.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Meadow.Masterships {
				var am AreaMastery
				am.Name = "meadow"
				am.Size = ChunkAreasInfoEx[m.ID].Meadow.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Meadow.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if !uuid.Equal(ChunkAreasInfoEx[m.ID].Lake.ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Lake.Masterships {
				var am AreaMastery
				am.Name = "lake"
				am.Size = ChunkAreasInfoEx[m.ID].Lake.Size
				am.AreaID = ChunkAreasInfoEx[m.ID].Lake.ID
				ChunkMasteryInfo[m.ID][v.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][v.Mastership.NameID], am)
			}
		}
		if len(ChunkAreasInfoEx[m.ID].Rivers) > 0 && !uuid.Equal(ChunkAreasInfoEx[m.ID].Rivers[0].ID, uuid.Nil) {
			for _, v := range ChunkAreasInfoEx[m.ID].Rivers {
				for _, vr := range v.Masterships {
					var am AreaMastery
					am.Name = "river"
					am.Size = v.Size
					am.AreaID = v.ID
					ChunkMasteryInfo[m.ID][vr.Mastership.NameID] = append(ChunkMasteryInfo[m.ID][vr.Mastership.NameID], am)

				}

			}

		}
	}
}
