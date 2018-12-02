package main

import (
	"C"

	"github.com/bluele/mecab-golang"
	"github.com/gin-gonic/gin"
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
			words = append(words, node.Feature())
			//			features := strings.Split(node.Feature(), ",")
			//words = extractSurface(features)
			if node.Next() != nil {
				break
			}
		}
		c.String(200, strings.Join(words[:], "\n"))
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

func extractSurface(feature []string) string {
	if feature[0] == "名詞" {
		return feature[1]
	}
	return ""
}
