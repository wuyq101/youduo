package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

// Game 导入的原始数据记录
type Game struct {
	ID              int64
	GameType        string
	GameName        string
	CreatorNickname string
	Blind           string
	MaxPlayerCount  int
	Duration        float64
	TotalHand       int
	PlayerID        int64
	PlayerNickname  string
	ClubID          int64
	ClubName        string
	Buy             int
	Sell            int
	InsuranceBuy    int
	InsuranceSell   int
	InsuranceAmount int
	ClubInsurance   int
	Insurance       int
	Income          int
	FinishTime      time.Time
}

// GameDao 数据读写文件
type GameDao struct {
}

// BatchSave 批量保存到数据库
func (dao GameDao) BatchSave(games []*Game) error {
}

// LoadFromExcel 从excel文件中加载
func (dao GameDao) LoadFromExcel(file string) ([]*Game, error) {
	f, err := xlsx.OpenFile(file)
	if err != nil {
		return nil, err
	}
	result := make([]*Game, 0)
	for _, sheet := range f.Sheets {
		if sheet.Name != "原始数据" {
			continue
		}
		fmt.Printf("len %d\n", len(sheet.Rows))
		for i, row := range sheet.Rows {
			if i == 0 {
				//行头
				fmt.Printf("%+v\n", row)
				continue
			}
			game := dao.convertToGame(row)
			result = append(result, game)
		}
	}
	return result, nil
}

func (dao GameDao) convertToGame(row *xlsx.Row) *Game {
	//	fmt.Printf("%+v\n", row)
	g := &Game{}
	for i, cell := range row.Cells {
		text := cell.String()
		switch i {
		case 0:
			g.GameType = text
		case 1:
			g.GameName = text
		case 2:
			g.CreatorNickname = text
		case 3:
			g.Blind = text
		case 4:
			v, _ := strconv.Atoi(text)
			g.MaxPlayerCount = v
		case 5:
			f, _ := strconv.ParseFloat(text, 64)
			g.Duration = f
		case 6:
			v, _ := strconv.Atoi(text)
			g.TotalHand = v
		case 7:
			v, _ := strconv.ParseInt(text, 10, 64)
			g.PlayerID = v
		case 8:
			g.PlayerNickname = text
		case 9:
			v, _ := strconv.ParseInt(text, 10, 64)
			g.ClubID = v
		case 10:
			g.ClubName = text
		case 11:
			v, _ := strconv.Atoi(text)
			g.Buy = v
		case 12:
			v, _ := strconv.Atoi(text)
			g.Sell = v
		case 13:
			v, _ := strconv.Atoi(text)
			g.InsuranceBuy = v
		case 14:
			v, _ := strconv.Atoi(text)
			g.InsuranceSell = v
		case 15:
			v, _ := strconv.Atoi(text)
			g.InsuranceAmount = v
		case 16:
			v, _ := strconv.Atoi(text)
			g.ClubInsurance = v
		case 17:
			v, _ := strconv.Atoi(text)
			g.Insurance = v
		case 18:
			v, _ := strconv.Atoi(text)
			g.Income = v
		case 19:
			tm, _ := time.ParseInLocation("2006-01-02 15:04:05", text, time.Local)
			g.FinishTime = tm
		}
	}
	return g
}
