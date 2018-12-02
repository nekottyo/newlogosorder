package main

import (
	"C"

	"github.com/bluele/mecab-golang"
	"github.com/gin-gonic/gin"

	"unicode/utf8"
)
import "strings"

func main() {
	r := gin.Default()

	m := generateMeCab()
	defer m.Destroy()

	r.GET("/", func(c *gin.Context) {
		words := []string{}
		text := c.Query("text")
		node := parseToNode(text, m)

		for {
			features := strings.Split(node.Feature(), ",")
			word := features[6]
			if extractSurface(features) {
				words = append(words, censorship(word))
			} else {
				words = append(words, word)
			}

			if node.Next() != nil {
				break
			}
		}
		c.String(200, strings.Join(words[:], ""))
	})
	r.Run(":8080")
}

func generateMeCab() *mecab.MeCab {
	m, err := mecab.New("-d /usr/local/lib/mecab/dic/mecab-ipadic-neologd")
	if err != nil {
		panic(err)
	}
	return m
}

func parseToNode(text string, m *mecab.MeCab) *mecab.Node {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	lt, err := m.NewLattice(text)
	if err != nil {
		panic(err)
	}
	node := tg.ParseToNode(lt)
	return node
}

func extractSurface(feature []string) bool {
	if feature[0] == "名詞" {
		return true
	}
	return false
}

func censorship(word string) string {
	text := []string{}
	for i := 0; i < utf8.RuneCountInString(word); i++ {
		text = append(text, "█")
	}
	return strings.Join(text[:], "")
}
