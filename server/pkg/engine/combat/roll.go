package combat

import "math/rand"

func Roll(max int, min int) int {
	return rand.Intn(max-min) + min
}

func RollNPickMHighest(n int, m int, max int, min int) {

}

func MaxHit(dmgBonus int, level int, modifier float32) int {

}

func Accuracy(attackBonus int, attackerLevel, defenseBonus int, defenderLevel int, modifier float32) float32 {

}

func Success(chance float32) bool {
	return chance > rand.Float32()
}

func HitSucceeded(attacker Attacker, defender Defender) bool {
	return true
}

func Damage() int {

}
