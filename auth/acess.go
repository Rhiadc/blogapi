package auth

import (
	"strconv"

	"github.com/rhiadc/blogapi/database/redis"
)

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func FetchAuth(authD AccessDetails) (uint64, error) {
	userid, err := redis.Client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := redis.Client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
