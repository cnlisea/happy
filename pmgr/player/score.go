package player

import "container/list"

func (p *Player) Score(index int) int32 {
	if p.score == nil || p.score.Len() == 0 || index > p.score.Len() {
		return 0
	}

	var (
		score int32
		e     *list.Element
		i     int
	)
	for e = p.score.Front(); e != nil; e = e.Next() {
		i++
		if index == 0 || index == i {
			score += e.Value.(int32)
		}
	}

	return score
}

func (p *Player) AddScore(score int32) {
	if p.score == nil {
		p.score = list.New()
	}
	p.score.PushBack(score)
}
