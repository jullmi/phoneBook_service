package internal

import "regexp"


func CreateIndex() error {
	Index = make(map[string]int)
	for i, k := range Data {
		key := k.Tel
		Index[key] = i
	}
	return nil
}

func MatchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}