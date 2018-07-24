package generator

import "github.com/SirMetathyst/go-entitas"

type byPriority []*entitas.CP

func (p byPriority) Len() int {
	return len(p)
}
func (p byPriority) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p byPriority) Less(i, j int) bool {
	return p[i].EventPriority() < p[j].EventPriority()
}
