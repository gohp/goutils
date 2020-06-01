package sensitive

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

/**
* @Author: Jam Wong
* @Date: 2020-01-02 20:07
 */

// Filter 敏感词过滤器
type Filter struct {
	trie  *Trie
	noise *regexp.Regexp
}

// New 返回一个敏感词过滤器
func New(path string) *Filter {
	f := &Filter{
		trie:  NewTrie(),
		noise: regexp.MustCompile(`[\|\s&%$@*]+`),
	}
	_ = f.LoadWordDict(path)

	return f
}

// UpdateNoisePattern 更新去噪模式
func (filter *Filter) UpdateNoisePattern(pattern string) {
	filter.noise = regexp.MustCompile(pattern)
}

// LoadWordDict 加载敏感词字典
func (filter *Filter) LoadWordDict(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return filter.Load(f)
}

// Load common method to add words
func (filter *Filter) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		filter.trie.Add(string(line))
	}

	return nil
}

// AddWord
func (filter *Filter) AddWord(words ...string) {
	filter.trie.Add(words...)
}

// DelWord
func (filter *Filter) DelWord(words ...string) {
	filter.trie.Del(words...)
}

// Filter
func (filter *Filter) Filter(text string) string {
	return filter.trie.Filter(text)
}

// Replace
func (filter *Filter) Replace(text string, repl rune) string {
	return filter.trie.Replace(text, repl)
}

// FindIn
func (filter *Filter) FindIn(text string) (bool, string) {
	text = filter.RemoveNoise(text)
	return filter.trie.FindIn(text)
}

// FindAll
func (filter *Filter) FindAll(text string) []string {
	return filter.trie.FindAll(text)
}

// Validate
func (filter *Filter) Validate(text string) (bool, string) {
	text = filter.RemoveNoise(text)
	return filter.trie.Validate(text)
}

// RemoveNoise
func (filter *Filter) RemoveNoise(text string) string {
	return filter.noise.ReplaceAllString(text, "")
}
