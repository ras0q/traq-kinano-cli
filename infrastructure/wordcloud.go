package infrastructure

import (
	"image"
	"image/color"
	"strings"

	mecab "github.com/bluele/mecab-golang"
	"github.com/psykhi/wordclouds"
)

func generateWordcloud() (image.Image, error) {
	msgs, err := getTraqDailyMsgs()
	if err != nil {
		return nil, err
	}

	wordMap, err := parseToNode(msgs)
	if err != nil {
		return nil, err
	}

	wc := wordclouds.NewWordcloud(
		wordMap,
		wordclouds.FontFile("assets/fonts/rounded-l-mplus-2c-medium.ttf"),
		wordclouds.Height(2048),
		wordclouds.Width(2048),
		wordclouds.FontMaxSize(256),
		wordclouds.FontMinSize(32),
		wordclouds.Colors([]color.Color{
			color.RGBA{0x1b, 0x1b, 0x1b, 0xff},
			color.RGBA{0x48, 0x48, 0x4B, 0xff},
			color.RGBA{0x59, 0x3a, 0xee, 0xff},
			color.RGBA{0x65, 0xCD, 0xFA, 0xff},
			color.RGBA{0x70, 0xD6, 0xBF, 0xff},
		}),
	)

	return wc.Draw(), nil
}

func parseToNode(msgs []string) (map[string]int, error) {
	m, err := mecab.New("-Owakati")
	if err != nil {
		return nil, err
	}

	tg, err := m.NewTagger()
	if err != nil {
		return nil, err
	}
	defer tg.Destroy()

	wordMap := make(map[string]int)
	for _, msg := range msgs {
		lt, err := m.NewLattice(msg)
		if err != nil {
			return nil, err
		}

		node := tg.ParseToNode(lt)

		wm := make(map[string]struct{})
		for {
			fea := strings.Split(node.Feature(), ",")
			sur := node.Surface()
			if fea[0] == "名詞" && fea[1] == "一般" && len(sur) > 1 {
				if _, ok := wm[sur]; !ok {
					wm[sur] = struct{}{}
				}
			}
			if node.Next() != nil {
				break
			}
		}

		for w := range wm {
			wordMap[w]++
		}

		lt.Destroy()
	}

	return wordMap, nil
}
