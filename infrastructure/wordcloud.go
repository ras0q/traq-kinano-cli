package infrastructure

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/Ras96/traq-kinano-cli/interfaces/external"
	"github.com/Ras96/traq-kinano-cli/util/config"
	mecab "github.com/bluele/mecab-golang"
	"github.com/gofrs/uuid"
	"github.com/psykhi/wordclouds"
)

// wordcloudに含めない単語
var excludeWords = []string{
	"trap",
	"感じ",
}

func isExcludedWord(word string) bool {
	for _, w := range excludeWords {
		if w == word {
			return true
		}
	}

	return false
}

func generateWordcloud() (image.Image, string, error) {
	msgs, err := getTraqDailyMsgs()
	if err != nil {
		return nil, "", err
	}

	wordMap, best, err := parseToNode(msgs)
	if err != nil {
		return nil, "", err
	}

	if len(wordMap) == 0 {
		return nil, "", fmt.Errorf("No wordcloud data")
	}

	wc := wordclouds.NewWordcloud(
		wordMap,
		wordclouds.FontFile("assets/fonts/rounded-l-mplus-2c-medium.ttf"),
		wordclouds.Height(1024),
		wordclouds.Width(1024),
		wordclouds.FontMaxSize(128),
		wordclouds.FontMinSize(8),
		wordclouds.Colors([]color.Color{
			color.RGBA{247, 144, 30, 255},
			color.RGBA{194, 69, 39, 255},
			color.RGBA{38, 103, 118, 255},
			color.RGBA{173, 210, 224, 255},
		}),
	)

	return wc.Draw(), best, nil
}

func parseToNode(msgs []string) (map[string]int, string, error) {
	m, err := mecab.New("-Owakati")
	if err != nil {
		return nil, "", err
	}

	tg, err := m.NewTagger()
	if err != nil {
		return nil, "", err
	}
	defer tg.Destroy()

	wordMap := make(map[string]int)
	for _, msg := range msgs {
		lt, err := m.NewLattice(msg)
		if err != nil {
			return nil, "", err
		}

		node := tg.ParseToNode(lt)

		wm := make(map[string]struct{})
		for {
			fea := strings.Split(node.Feature(), ",")
			sur := strings.ToLower(node.Surface())
			if fea[0] == "名詞" && fea[1] == "一般" && len(sur) > 1 {
				if _, found := wm[sur]; !found && !isExcludedWord(sur) {
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

	var (
		best      string
		bestCount int
	)

	for word, cnt := range wordMap {
		if cnt > bestCount {
			best = word
			bestCount = cnt
		}
	}

	return wordMap, best, nil
}

func PostWordcloudToTraq(q external.TraqAPI) error {
	img, best, err := generateWordcloud()
	if err != nil {
		return fmt.Errorf("Error generating wordcloud: %w", err)
	}

	path, _ := filepath.Abs("./wordcloud.png")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Error creating wordcloud file: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("Error encoding wordcloud: %w", err)
	}

	file.Seek(0, os.SEEK_SET)

	cid := uuid.FromStringOrNil(config.Traq.HomeCh)
	fid, err := q.PostFile(cid, file)
	if err != nil {
		return fmt.Errorf("Error creating wordcloud: %w", err)
	}

	if err := q.PostMessage(
		cid,
		fmt.Sprintf("今日は『%s』の日だったやんね:tada:\nhttps://q.trap.jp/files/%s", best, fid.String()),
		true,
	); err != nil {
		return fmt.Errorf("Error posting wordcloud: %w", err)
	}

	return nil
}
