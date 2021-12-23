package infrastructure

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	mecab "github.com/bluele/mecab-golang"
	"github.com/psykhi/wordclouds"
)

func generateWordcloud() (image.Image, error) {
	msgs, err := getTraqDailyMsgs()
	if err != nil {
		return nil, err
	}

	wordMap := parseToNode(strings.Join(msgs, ","))

	wc := wordclouds.NewWordcloud(
		wordMap,
		wordclouds.FontFile("assets/fonts/rounded-l-mplus-2c-medium.ttf"),
		wordclouds.Height(2048),
		wordclouds.Width(2048),
		wordclouds.Colors([]color.Color{
			color.RGBA{0x1b, 0x1b, 0x1b, 0xff},
			color.RGBA{0x48, 0x48, 0x4B, 0xff},
			color.RGBA{0x59, 0x3a, 0xee, 0xff},
			color.RGBA{0x65, 0xCD, 0xFA, 0xff},
			color.RGBA{0x70, 0xD6, 0xBF, 0xff},
		}),
	)

	fmt.Println(wordMap)

	return wc.Draw(), nil
}

func parseToNode(text string) map[string]int {
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}

	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(text)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	f, _ := os.Create("s.log")
	defer f.Close()

	node := tg.ParseToNode(lt)
	wordMap := make(map[string]int)
	for {
		f.WriteString(node.Surface()+"\n")
		f.WriteString(node.Feature()+"\n")
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" && features[1] != "サ変接続" && len(node.Surface()) > 1 {
			wordMap[node.Surface()]++
		}
		if node.Next() != nil {
			break
		}
	}

	return wordMap
}
