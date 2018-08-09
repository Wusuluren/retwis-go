package retwis

import (
	"bytes"
	"encoding/binary"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	redis.Conn
}

func intfc2string(i interface{}) (val string) {
	if i != nil {
		val = string(i.([]uint8))
	}
	return val
}

func intfc2int64(i interface{}) (val int64) {
	if i != nil {
		val = i.(int64)
	}
	return val
}

func intfc2float64(i interface{}) (val float64) {
	if i != nil {
		buf := bytes.NewBuffer([]byte(i.([]uint8)))
		binary.Write(buf, binary.LittleEndian, &val)
	}
	return val
}

func intfc2mapStringString(i interface{}) (val map[string]string) {
	if i != nil {
		val = make(map[string]string)
		items := i.([]interface{})
		for i := 0; i < len(items); i += 2 {
			k := string(items[i].([]uint8))
			v := string(items[i+1].([]uint8))
			val[k] = v
		}
	}
	return val
}

func intfc2stringSlice(i interface{}) (val []string) {
	if i != nil {
		for _, item := range i.([]interface{}) {
			if item != nil {
				val = append(val, string(item.([]uint8)))
			}
		}
	}
	return val
}

// connection command

func (r *Redis) Ping() (string, error) {
	var val string
	reply, err := r.Conn.Do("PING")
	if reply != nil {
		val = reply.(string)
	}
	return val, err
}

func (r *Redis) Auth(passwd string) (int64, error) {
	reply, err := r.Conn.Do("AUTH", passwd)
	if _, ok := reply.(redis.Error); ok {
		return 0, err
	}
	return intfc2int64(reply), err
}

func (r *Redis) Select(db int) error {
	_, err := r.Conn.Do("SELECT", db)
	return err
}

func (r *Redis) Quit() error {
	_, err := r.Conn.Do("QUIT")
	return err
}

func (r *Redis) Echo(msg string) error {
	_, err := r.Conn.Do("ECHO", msg)
	return err
}

// key command

func (r *Redis) Del(key string) error {
	_, err := r.Conn.Do("DEL", key)
	return err
}

func (r *Redis) Dump(key string) (string, error) {
	reply, err := r.Conn.Do("DUMP", key)
	return intfc2string(reply), err
}

func (r *Redis) Exists(key string) (int64, error) {
	reply, err := r.Conn.Do("EXISTS", key)
	return intfc2int64(reply), err
}

func (r *Redis) Expire(key string, seconds int) error {
	_, err := r.Conn.Do("EXPIRE", key, seconds)
	return err
}

func (r *Redis) Expireat(key string, timestamp int64) error {
	_, err := r.Conn.Do("EXPIREAT", key, timestamp)
	return err
}

func (r *Redis) Pexpire(key string, mseconds int64) error {
	_, err := r.Conn.Do("PEXPIRE", key, mseconds)
	return err
}

func (r *Redis) Pexpireat(key string, msecondsTimestamp int64) error {
	_, err := r.Conn.Do("PEXPIREAT", key, msecondsTimestamp)
	return err
}

func (r *Redis) Keys(pattern string) ([]string, error) {
	reply, err := r.Conn.Do("KEYS", pattern)
	return intfc2stringSlice(reply), err
}

func (r *Redis) Move(key string, db int64) error {
	_, err := r.Conn.Do("MOVE", key, db)
	return err
}

func (r *Redis) Persist(key string) error {
	_, err := r.Conn.Do("PERSIST", key)
	return err
}

func (r *Redis) Pttl(key string) (int64, error) {
	reply, err := r.Conn.Do("PTTL", key)
	return intfc2int64(reply), err
}

func (r *Redis) Ttl(key string) (int64, error) {
	reply, err := r.Conn.Do("TTL", key)
	return intfc2int64(reply), err
}

func (r *Redis) RandomKey() (string, error) {
	reply, err := r.Conn.Do("RANDOMKEY")
	return string(reply.([]uint8)), err
}

func (r *Redis) Rename(key, newkey string) error {
	_, err := r.Conn.Do("RENAME", key, newkey)
	return err
}

func (r *Redis) RenameNx(key, newkey string) error {
	_, err := r.Conn.Do("RENAMENX", key, newkey)
	return err
}

func (r *Redis) Type(key string) (string, error) {
	var val string
	reply, err := r.Conn.Do("TYPE", key)
	if reply != nil {
		val = reply.(string)
	}
	return val, err
}

// string command

func (r *Redis) Set(key, val string) error {
	_, err := r.Conn.Do("SET", key, val)
	return err
}

func (r *Redis) Get(key string) (string, error) {
	reply, err := r.Conn.Do("GET", key)
	return string(reply.([]uint8)), err
}

func (r *Redis) GetRange(key string, start, end int) (string, error) {
	reply, err := r.Conn.Do("GETRANGE", key, start, end)
	return string(reply.([]uint8)), err
}

func (r *Redis) GetSet(key, val string) (string, error) {
	reply, err := r.Conn.Do("GETSET", key, val)
	return string(reply.([]uint8)), err
}

func (r *Redis) GetBit(key string, offset int) (int64, error) {
	reply, err := r.Conn.Do("GETBIT", key, offset)
	return intfc2int64(reply), err
}

func (r *Redis) MGet(keys ...string) ([]string, error) {
	reply, err := r.Conn.Do("MGET", keys)
	return intfc2stringSlice(reply), err
}

func (r *Redis) SetBit(key string, offset int, val int) error {
	_, err := r.Conn.Do("SETBIT", key, offset, val)
	return err
}

func (r *Redis) SetEx(key string, seconds int64, val string) error {
	_, err := r.Conn.Do("SETEX", key, seconds, val)
	return err
}

func (r *Redis) SetNx(key, val string) error {
	_, err := r.Conn.Do("SETNX", key, val)
	return err
}

func (r *Redis) SetRange(key string, offset int, val string) error {
	_, err := r.Conn.Do("SETRANGE", key, offset, val)
	return err
}

func (r *Redis) StrLen(key string) (int64, error) {
	reply, err := r.Conn.Do("STRLEN", key)
	return intfc2int64(reply), err
}

func (r *Redis) MSet(args ...interface{}) error {
	_, err := r.Conn.Do("MSET", args...)
	return err
}

func (r *Redis) MSetNx(args ...interface{}) error {
	_, err := r.Conn.Do("MSETNX", args...)
	return err
}

func (r *Redis) PSetEx(key string, mseconds int64, val string) error {
	_, err := r.Conn.Do("PSETEX", key, mseconds, val)
	return err
}

func (r *Redis) Incr(key string) (int64, error) {
	reply, err := r.Conn.Do("INCR", key)
	return intfc2int64(reply), err
}

func (r *Redis) IncrBy(key string, increment int) (int64, error) {
	reply, err := r.Conn.Do("INCRBY", key, increment)
	return intfc2int64(reply), err
}

func (r *Redis) IncrByFloat(key string, increment float64) (float64, error) {
	reply, err := r.Conn.Do("INCRBYFLOAT", key, increment)
	return intfc2float64(reply), err
}

func (r *Redis) Decr(key string) (int64, error) {
	reply, err := r.Conn.Do("DECR", key)
	return intfc2int64(reply), err
}

func (r *Redis) DecrBy(key string, increment int) (int64, error) {
	reply, err := r.Conn.Do("DECRBY", key, increment)
	return intfc2int64(reply), err
}

func (r *Redis) Append(key, val string) error {
	_, err := r.Conn.Do("APPEND", key, val)
	return err
}

// hash command

func (r *Redis) HDel(key string, fields ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, fields...)
	_, err := r.Conn.Do("HDEL", input...)
	return err
}

func (r *Redis) HExists(key string, field string) (int64, error) {
	reply, err := r.Conn.Do("HEXISTS", key, field)
	return intfc2int64(reply), err
}

func (r *Redis) HGet(key string, field string) (string, error) {
	reply, err := r.Conn.Do("HGET", key, field)
	return intfc2string(reply), err
}

func (r *Redis) HGetAll(key string) (map[string]string, error) {
	reply, err := r.Conn.Do("HGETALL", key)
	return intfc2mapStringString(reply), err
}

func (r *Redis) HIncrBy(key, field string, increment int) (int64, error) {
	reply, err := r.Conn.Do("HINCRBY", key, field, increment)
	return intfc2int64(reply), err
}

func (r *Redis) HIncrByFloat(key, field string, increment float64) (float64, error) {
	reply, err := r.Conn.Do("HINCRBYFLOAT", key, field, increment)
	return intfc2float64(reply), err
}

func (r *Redis) HKeys(key string) ([]string, error) {
	reply, err := r.Conn.Do("HKEYS", key)
	return intfc2stringSlice(reply), err
}

func (r *Redis) HLen(key string) (int64, error) {
	reply, err := r.Conn.Do("HLEN", key)
	return intfc2int64(reply), err
}

func (r *Redis) HMGet(key string, fields ...interface{}) ([]string, error) {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, fields...)
	reply, err := r.Conn.Do("HMGET", input...)
	return intfc2stringSlice(reply), err
}

func (r *Redis) HMSet(key string, args ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, args...)
	_, err := r.Conn.Do("HMSET", input...)
	return err
}

func (r *Redis) HSet(key string, filed, val interface{}) error {
	_, err := r.Conn.Do("HSET", key, filed, val)
	return err
}

func (r *Redis) HSetNx(key string, filed, val interface{}) error {
	_, err := r.Conn.Do("HSETNX", key, filed, val)
	return err
}

func (r *Redis) HVals(key string) ([]string, error) {
	reply, err := r.Conn.Do("HVALS", key)
	return intfc2stringSlice(reply), err
}

// list command

func (r *Redis) BLPop(timeout int, keys ...interface{}) (string, error) {
	input := make([]interface{}, 0)
	input = append(input, keys...)
	input = append(input, timeout)
	reply, err := r.Conn.Do("BLPOP", input...)
	return intfc2string(reply), err
}

func (r *Redis) BRPop(timeout int, keys ...interface{}) (string, error) {
	input := make([]interface{}, 0)
	input = append(input, keys...)
	input = append(input, timeout)
	reply, err := r.Conn.Do("BRPOP", input...)
	return intfc2string(reply), err
}

func (r *Redis) BRPopLPush(src, dst string, timeout int) error {
	_, err := r.Conn.Do("BRPOPLPUSH", src, dst, timeout)
	return err
}

func (r *Redis) LIndex(key string, index int) (string, error) {
	reply, err := r.Conn.Do("LINDEX", key, index)
	return intfc2string(reply), err
}

func (r *Redis) LInsertBefore(key string, pivot, val string) error {
	_, err := r.Conn.Do("LINSERT", key, "BEFORE", pivot, val)
	return err
}

func (r *Redis) LInsertAfter(key string, pivot, val string) error {
	_, err := r.Conn.Do("LINSERT", key, "AFTER", pivot, val)
	return err
}

func (r *Redis) LLen(key string) (int64, error) {
	reply, err := r.Conn.Do("LLEN", key)
	return intfc2int64(reply), err
}

func (r *Redis) LPop(key string) (string, error) {
	reply, err := r.Conn.Do("LPOP", key)
	return intfc2string(reply), err
}

func (r *Redis) LPush(key string, vals ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, vals...)
	_, err := r.Conn.Do("LPUSH", input...)
	return err
}

func (r *Redis) LPushX(key string, val interface{}) error {
	_, err := r.Conn.Do("LPUSHX", key, val)
	return err
}

func (r *Redis) LRange(key string, start, stop int) ([]string, error) {
	reply, err := r.Conn.Do("LRANGE", key, start, stop)
	return intfc2stringSlice(reply), err
}

func (r *Redis) LRem(key string, count int, val interface{}) error {
	_, err := r.Conn.Do("LREM", key, count, val)
	return err
}

func (r *Redis) LSet(key string, index int, val interface{}) error {
	_, err := r.Conn.Do("LSET", key, index, val)
	return err
}

func (r *Redis) LTrim(key string, start, stop int) error {
	_, err := r.Conn.Do("LTRIM", key, start, stop)
	return err
}

func (r *Redis) RPop(key string) (string, error) {
	reply, err := r.Conn.Do("RPOP", key)
	return intfc2string(reply), err
}

func (r *Redis) RPopLPush(src, dst string) error {
	_, err := r.Conn.Do("RPOPLPUSH", src, dst)
	return err
}

func (r *Redis) RPush(key string, vals ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, vals...)
	_, err := r.Conn.Do("RPUSH", input...)
	return err
}

func (r *Redis) RPushX(key string, val interface{}) error {
	_, err := r.Conn.Do("RPUSHX", key, val)
	return err
}

// set command

func (r *Redis) SAdd(key string, members ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, members...)
	_, err := r.Conn.Do("SADD", input...)
	return err
}

func (r *Redis) SCard(key string) (int64, error) {
	reply, err := r.Conn.Do("SCARD", key)
	return intfc2int64(reply), err
}

func (r *Redis) SDiff(keys ...interface{}) ([]string, error) {
	reply, err := r.Conn.Do("SDIFF", keys...)
	return intfc2stringSlice(reply), err
}

func (r *Redis) SDiffStore(dst string, keys ...interface{}) (int64, error) {
	input := make([]interface{}, 0)
	input = append(input, dst)
	input = append(input, keys...)
	reply, err := r.Conn.Do("SDIFFSTORE", input...)
	return intfc2int64(reply), err
}

func (r *Redis) SInter(keys ...interface{}) ([]string, error) {
	reply, err := r.Conn.Do("SINTER", keys...)
	return intfc2stringSlice(reply), err
}

func (r *Redis) SInterStore(dst string, keys ...interface{}) (int64, error) {
	input := make([]interface{}, 0)
	input = append(input, dst)
	input = append(input, keys...)
	reply, err := r.Conn.Do("SINTERSTORE", input...)
	return intfc2int64(reply), err
}

func (r *Redis) SIsMember(key, member string) (int64, error) {
	reply, err := r.Conn.Do("SISMEMBER", key, member)
	return intfc2int64(reply), err
}

func (r *Redis) SMembers(key string) ([]string, error) {
	reply, err := r.Conn.Do("SMEMBERS", key)
	return intfc2stringSlice(reply), err
}

func (r *Redis) SMove(src, dst, member string) error {
	_, err := r.Conn.Do("SMOVE", src, dst, member)
	return err
}

func (r *Redis) SPop(key string) (string, error) {
	reply, err := r.Conn.Do("SPOP", key)
	return intfc2string(reply), err
}

func (r *Redis) SRandMember(key string, count int) ([]string, error) {
	var err error
	var reply interface{}
	if count > 0 {
		reply, err = r.Conn.Do("SRANDMEMBER", key, count)
	} else {
		reply, err = r.Conn.Do("SRANDMEMBER", key)
	}
	return intfc2stringSlice(reply), err
}

func (r *Redis) SRem(key string, members ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, members...)
	_, err := r.Conn.Do("SREM", input...)
	return err
}

func (r *Redis) SUnionStore(dst string, keys ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, dst)
	input = append(input, keys...)
	_, err := r.Conn.Do("SUNIONSTORE", input...)
	return err
}

// zset command

func (r *Redis) ZAdd(key string, args ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, args...)
	_, err := r.Conn.Do("ZADD", input...)
	return err
}

func (r *Redis) ZCard(key string) (int64, error) {
	reply, err := r.Conn.Do("ZCARD", key)
	return intfc2int64(reply), err
}

func (r *Redis) ZCount(key string, min, max float64) (int64, error) {
	reply, err := r.Conn.Do("ZCOUNT", key, min, max)
	return intfc2int64(reply), err
}

func (r *Redis) ZINCRBY(key string, increment, member float64) error {
	_, err := r.Conn.Do("ZINCRBY", key, increment, member)
	return err
}

func (r *Redis) ZInterStore(dst string, numkeys int, keys ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, dst)
	input = append(input, numkeys)
	input = append(input, keys...)
	_, err := r.Conn.Do("ZINTERSTORE", input...)
	return err
}

func (r *Redis) ZRange(key string, start, stop int64) ([]string, error) {
	reply, err := r.Conn.Do("ZRANGE", key, start, stop)
	return intfc2stringSlice(reply), err
}

func (r *Redis) ZRangeByScore(key string, min, max float64) ([]string, error) {
	reply, err := r.Conn.Do("ZRANGEBYSCORE", key, min, max)
	return intfc2stringSlice(reply), err
}

func (r *Redis) ZRangeByScoreLimit(key string, min, max float64, offset, count int) ([]string, error) {
	reply, err := r.Conn.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count)
	return intfc2stringSlice(reply), err
}

func (r *Redis) ZRank(key, member string) (int64, error) {
	reply, err := r.Conn.Do("ZRANK", key, member)
	return intfc2int64(reply), err
}

func (r *Redis) ZRem(key string, members ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, key)
	input = append(input, members...)
	_, err := r.Conn.Do("ZREM", input...)
	return err
}

func (r *Redis) ZRemRangeByRank(key string, start, stop int64) error {
	_, err := r.Conn.Do("ZREMRANGEBYRANK", key, start, stop)
	return err
}

func (r *Redis) ZRemRangeByScore(key string, min, max float64) error {
	_, err := r.Conn.Do("ZREMRANGEBYSCORE", key, min, max)
	return err
}

func (r *Redis) ZRevRange(key string, start, stop int64) ([]string, error) {
	reply, err := r.Conn.Do("ZREVRANGE", key, start, stop)
	return intfc2stringSlice(reply), err
}

func (r *Redis) ZRevRangeByScore(key string, min, max float64) ([]string, error) {
	reply, err := r.Conn.Do("ZREVRANGEBYSCORE", key, min, max)
	return intfc2stringSlice(reply), err
}

func (r *Redis) ZRevRank(key, member string) (int64, error) {
	reply, err := r.Conn.Do("ZREVRANK", key, member)
	return intfc2int64(reply), err
}

func (r *Redis) ZScore(key, member string) (float64, error) {
	reply, err := r.Conn.Do("ZSCORE", key, member)
	return intfc2float64(reply), err
}

func (r *Redis) ZUnionStore(dst string, numkeys int, keys ...interface{}) error {
	input := make([]interface{}, 0)
	input = append(input, dst)
	input = append(input, numkeys)
	input = append(input, keys...)
	_, err := r.Conn.Do("ZUNIONSTORE", input...)
	return err
}

// pubsub command

func (r *Redis) PSubscribe(patterns ...interface{}) error {
	_, err := r.Conn.Do("PSUBSCRIBE", patterns)
	return err
}

func (r *Redis) PUnSubscribe(patterns ...interface{}) error {
	_, err := r.Conn.Do("PUNSUBSCRIBE", patterns)
	return err
}

func (r *Redis) Subscribe(channels ...interface{}) error {
	_, err := r.Conn.Do("SUBSCRIBE", channels)
	return err
}

func (r *Redis) UnSubscribe(channels ...interface{}) error {
	_, err := r.Conn.Do("UNSUBSCRIBE", channels)
	return err
}

// transactions command

func (r *Redis) Multi() error {
	_, err := r.Conn.Do("MULTI")
	return err
}

func (r *Redis) Exec() error {
	_, err := r.Conn.Do("EXEC")
	return err
}

func (r *Redis) Discard() error {
	_, err := r.Conn.Do("DISCARD")
	return err
}
