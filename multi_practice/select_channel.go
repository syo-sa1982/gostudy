package main
import (
	"sync"
	"fmt"
	"log"
)


type SkillMaster struct {
	ID         int
	CategoryID int
	SkillName  string
	Value      int
}
var SkillMasterList = []SkillMaster{
	SkillMaster{ID: 1, CategoryID: 3, SkillName: "言いくるめ", Value: 5},
	SkillMaster{ID: 2, CategoryID: 1, SkillName: "医学", Value: 5},
	SkillMaster{ID: 3, CategoryID: 3, SkillName: "運転", Value: 20},
	SkillMaster{ID: 4, CategoryID: 1, SkillName: "応急手当", Value: 20},
	SkillMaster{ID: 5, CategoryID: 2, SkillName: "オカルト", Value: 5},
	SkillMaster{ID: 6, CategoryID: 2, SkillName: "回避", Value: 0},
	SkillMaster{ID: 7, CategoryID: 3, SkillName: "化学", Value: 1},
	SkillMaster{ID: 8, CategoryID: 3, SkillName: "鍵開け", Value: 1},
	SkillMaster{ID: 9, CategoryID: 3, SkillName: "隠す", Value: 15},
	SkillMaster{ID: 10, CategoryID: 3, SkillName: "隠れる", Value: 10},
	SkillMaster{ID: 11, CategoryID: 3, SkillName: "機械修理", Value: 20},
	SkillMaster{ID: 12, CategoryID: 3, SkillName: "聞き耳", Value: 25},
	SkillMaster{ID: 13, CategoryID: 6, SkillName: "クトゥルフ神話", Value: 0},
	SkillMaster{ID: 14, CategoryID: 3, SkillName: "芸術", Value: 5},
	SkillMaster{ID: 15, CategoryID: 3, SkillName: "経理", Value: 10},
	SkillMaster{ID: 16, CategoryID: 3, SkillName: "考古学", Value: 1},
	SkillMaster{ID: 17, CategoryID: 3, SkillName: "コンピューター", Value: 1},
	SkillMaster{ID: 18, CategoryID: 3, SkillName: "忍び歩き", Value: 10},
	SkillMaster{ID: 19, CategoryID: 3, SkillName: "写真術", Value: 10},
	SkillMaster{ID: 20, CategoryID: 3, SkillName: "重機械操作", Value: 1},
	SkillMaster{ID: 21, CategoryID: 3, SkillName: "乗馬", Value: 5},
	SkillMaster{ID: 22, CategoryID: 3, SkillName: "信用", Value: 15},
	SkillMaster{ID: 23, CategoryID: 4, SkillName: "心理学", Value: 5},
	SkillMaster{ID: 24, CategoryID: 3, SkillName: "人類学", Value: 1},
	SkillMaster{ID: 25, CategoryID: 3, SkillName: "水泳", Value: 25},
	SkillMaster{ID: 26, CategoryID: 3, SkillName: "製作", Value: 5},
	SkillMaster{ID: 27, CategoryID: 1, SkillName: "精神分析", Value: 1},
	SkillMaster{ID: 28, CategoryID: 3, SkillName: "生物学", Value: 1},
	SkillMaster{ID: 29, CategoryID: 3, SkillName: "説得", Value: 15},
	SkillMaster{ID: 30, CategoryID: 3, SkillName: "操縦", Value: 1},
	SkillMaster{ID: 31, CategoryID: 3, SkillName: "地質学", Value: 1},
	SkillMaster{ID: 32, CategoryID: 3, SkillName: "跳躍", Value: 25},
	SkillMaster{ID: 33, CategoryID: 3, SkillName: "追跡", Value: 10},
	SkillMaster{ID: 34, CategoryID: 3, SkillName: "電器修理", Value: 10},
	SkillMaster{ID: 35, CategoryID: 3, SkillName: "電子工学", Value: 1},
	SkillMaster{ID: 36, CategoryID: 3, SkillName: "天文学", Value: 1},
	SkillMaster{ID: 37, CategoryID: 2, SkillName: "投擲", Value: 25},
	SkillMaster{ID: 38, CategoryID: 3, SkillName: "登攀", Value: 40},
	SkillMaster{ID: 39, CategoryID: 3, SkillName: "図書館", Value: 25},
	SkillMaster{ID: 40, CategoryID: 3, SkillName: "ナビゲート", Value: 10},
	SkillMaster{ID: 41, CategoryID: 3, SkillName: "値切り", Value: 5},
	SkillMaster{ID: 42, CategoryID: 3, SkillName: "博物学", Value: 10},
	SkillMaster{ID: 43, CategoryID: 3, SkillName: "物理学", Value: 1},
	SkillMaster{ID: 44, CategoryID: 3, SkillName: "変装", Value: 1},
	SkillMaster{ID: 45, CategoryID: 3, SkillName: "法律", Value: 5},
	SkillMaster{ID: 46, CategoryID: 3, SkillName: "外国語", Value: 1},
	SkillMaster{ID: 47, CategoryID: 3, SkillName: "母国語", Value: 0},
	SkillMaster{ID: 48, CategoryID: 2, SkillName: "格闘", Value: 1},
	SkillMaster{ID: 49, CategoryID: 3, SkillName: "目星", Value: 25},
	SkillMaster{ID: 50, CategoryID: 3, SkillName: "薬学", Value: 1},
	SkillMaster{ID: 51, CategoryID: 3, SkillName: "歴史", Value: 20},
	SkillMaster{ID: 52, CategoryID: 4, SkillName: "拳銃", Value: 20},
	SkillMaster{ID: 53, CategoryID: 4, SkillName: "サブマシンガン", Value: 15},
	SkillMaster{ID: 54, CategoryID: 4, SkillName: "ショットガン", Value: 30},
	SkillMaster{ID: 55, CategoryID: 4, SkillName: "マシンガン", Value: 15},
	SkillMaster{ID: 56, CategoryID: 4, SkillName: "ライフル", Value: 25},
	SkillMaster{ID: 57, CategoryID: 2, SkillName: "キック", Value: 25},
	SkillMaster{ID: 58, CategoryID: 2, SkillName: "組み付き", Value: 25},
	SkillMaster{ID: 59, CategoryID: 2, SkillName: "こぶし", Value: 50},
	SkillMaster{ID: 60, CategoryID: 2, SkillName: "頭突き", Value: 10},
}

//var SkillMap = make(map[int]int)

func worker(SkillList []SkillMaster) (<-chan string) {
	var wg sync.WaitGroup
	receiver := make(chan string)
	go func() {
		for key := range SkillList {
			wg.Add(1)
			i := SkillList[key].ID
			go func(i int) {
				msg := fmt.Sprintf("%d %s done", i, SkillList[key].SkillName)
				receiver <- msg
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(receiver)
	}()
	return receiver
}

func main() {
	receiver := worker(SkillMasterList)
	for {
		receive, ok := <-receiver
		if !ok {
			log.Println("closed")
			return
		}
		log.Println(receive)
	}
}