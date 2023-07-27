package helper

import "fmt"

var PlayerKeyRedis string = "player_key"

func GenerateRedisKey(Id uint) string {
	return fmt.Sprintf("%s_%d", PlayerKeyRedis, Id)
}
