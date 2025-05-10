package score

var Score []int

func AddScore(score int) {
	Score = append(Score, score)
}

func GetScore() []int {
	return Score
}
func UpdatedScore(sc int, idx int) {
	Score[idx] = sc
}