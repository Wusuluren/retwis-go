package retwis

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
	"bytes"
	"encoding/binary"
)

type Redis struct {
	redis.Conn
}

func (r *Redis) Incr(field string) (string, error) {
	reply, err := r.Conn.Do("INCR", field)
	if err != nil {
		return "", err
	}
	var value string
	if reply != nil {
		value = strconv.FormatInt(reply.(int64), 10)
	}
	return value, err
}

func (r* Redis) Hget(field, key string) (string, error) {
	reply, err := r.Conn.Do("HGET", field, key)
	if err != nil {
		return "", err
	}
	var value string
	if reply != nil {
		value = string(reply.([]uint8))
	}
	return value, nil
}

func (r* Redis) Hgetall(field string) (map[string]string, error) {
	vals := make(map[string]string)
	reply, err := r.Conn.Do("HGETALL", field)
	if err != nil {
		return vals, err
	}
	if reply != nil {
		replies := reply.([]interface{})
		for i := 0; i < len(replies); i += 2 {
			k := string(replies[i].([]uint8))
			v := string(replies[i+1].([]uint8))
			vals[k] = v
		}
	}
	return vals, nil
}

func (r *Redis) Hset(field string, key, value interface{}) error {
	_, err := r.Conn.Do("HSET", field, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Hdel(field, key string) error {
	_, err := r.Conn.Do("HDEL", field, key)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Hmset(args ...interface{}) error {
	_, err := r.Conn.Do("HMSET", args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Zadd(key string, score int64, value interface{}) error {
	_, err := r.Conn.Do("ZADD", key, score, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Zrem(key string, value interface{}) error {
	_, err := r.Conn.Do("ZREM", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Zcard(key string) (string, error) {
	reply, err := r.Conn.Do("ZCARD", key)
	if err != nil {
		return "0", err
	}
	var value string
	if reply != nil {
		value = strconv.FormatInt(reply.(int64), 10)
	}
	return value, err
}

func (r *Redis) Zscore(key, member string) (float64, error) {
	reply, err := r.Conn.Do("ZSCORE", key, member)
	if err != nil {
		return 0, err
	}
	var value float64
	if reply != nil {
		buf := bytes.NewBuffer(reply.([]uint8))
		binary.Read(buf, binary.LittleEndian, &value)
	}
	return value, err
}

func (r *Redis) Lpush(key, value string) error {
	_, err := r.Conn.Do("LPUSH", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Lrange(key string, start, end int) ([]interface{}, error) {
	vals := make([]interface{}, 0)
	reply, err := r.Conn.Do("LRANGE", key, start, end)
	if err != nil {
		return nil, err
	}
	if reply != nil {
		vals = reply.([]interface{})
	}
	return vals, nil
}

func (r *Redis) Zrange(key string, start, end int) ([]interface{}, error) {
	vals := make([]interface{}, 0)
	reply, err := r.Conn.Do("ZRANGE", key, start, end)
	if err != nil {
		return nil, err
	}
	if reply != nil {
		vals = reply.([]interface{})
	}
	return vals, nil
}

func (r *Redis) Ltrim(key string, start, end int) error {
	_, err := r.Conn.Do("Ltrim", key, start, end)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Zrevrange(key string, start, end int) ([]interface{}, error) {
	vals := make([]interface{}, 0)
	reply, err := r.Conn.Do("ZREVRANGE", key, start, end)
	if err != nil {
		return nil, err
	}
	if reply != nil {
		vals = reply.([]interface{})
	}
	return vals, nil
}