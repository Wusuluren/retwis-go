package retwis

import (
	"testing"
)

var r *Redis

func getRedis(t *testing.T) *Redis {
	var err error
	if r == nil {
		r, err = redisLink()
		if err != nil {
			t.Fatal(err)
		}
	} else if pong, err := r.Ping(); pong == "" {
		r, err = redisLink()
		if err != nil {
			t.Fatal(err)
		}
	}
	return r
}

// connection command
func TestRedis_ConnectionCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	_, err = r.Ping()
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Auth("passwd")
	if err != nil {
		t.Log(err)
	}
	err = r.Echo("hello redis")
	if err != nil {
		t.Log(err)
	}
	err = r.Select(1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

// key command
func TestRedis_KeyCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	err = r.Del("key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Dump("key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Exists("key")
	if err != nil {
		t.Fatal(err)
	}
	err = r.Expire("key", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Expireat("key", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Pexpire("key", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Pexpireat("key", 1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Keys("*")
	if err != nil {
		t.Fatal(err)
	}
	err = r.Move("key", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Persist("key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Pttl("key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Ttl("key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.RandomKey()
	if err != nil {
		t.Fatal(err)
	}
	err = r.Rename("key", "new_key")
	if err != nil {
		t.Log(err)
	}
	err = r.RenameNx("key", "new_key")
	if err != nil {
		t.Log(err)
	}
	_, err = r.Type("key")
	if err != nil {
		t.Fatal(err)
	}
}

// string command
func TestRedis_StringCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	err = r.Set("skey", "s_val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Get("skey")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.GetRange("skey", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.GetSet("skey", "s_new_val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.GetBit("skey", 1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.MGet("skey", "not_exist")
	if err != nil {
		t.Fatal(err)
	}
	err = r.SetBit("skey", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.SetEx("skey2", 10, "s_val2")
	if err != nil {
		t.Fatal(err)
	}
	err = r.SetNx("skey", "s_val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.SetRange("skey", 1, "?")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.StrLen("skey")
	if err != nil {
		t.Fatal(err)
	}
	err = r.MSet("skey", "s_val", "not_exist", "not_exist")
	if err != nil {
		t.Fatal(err)
	}
	err = r.MSetNx("skey", "s_val", "not_exist", "not_exist")
	if err != nil {
		t.Fatal(err)
	}
	err = r.PSetEx("skey2", 10, "s_val2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Incr("skey3")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.IncrBy("skey3", 10)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.IncrByFloat("skey3", 20.0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Decr("skey3")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.DecrBy("skey3", 10)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Append("skey", "s_val")
	if err != nil {
		t.Fatal(err)
	}
}

// hash command
func TestRedis_HashCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	err = r.HDel("h_key", "h_field", "h_field2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HExists("h_key", "h_field")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HGet("h_key", "h_field")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HGetAll("h_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HIncrBy("h_key", "h_field_incr", 1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HIncrByFloat("h_key", "h_field_incr", 1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HKeys("h_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HLen("h_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HMGet("h_key", "h_field")
	if err != nil {
		t.Fatal(err)
	}
	err = r.HMSet("h_key", "h_field", "h_val", "h_field2", "h_val2")
	if err != nil {
		t.Fatal(err)
	}
	err = r.HSet("h_key", "h_field", "h_val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.HSetNx("h_key", "h_field", "h_val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.HVals("h_key")
	if err != nil {
		t.Fatal(err)
	}
}

// list command
func TestRedis_ListCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	_, err = r.BLPop(1, "l_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.BRPop(1, "l_key")
	if err != nil {
		t.Fatal(err)
	}
	err = r.BRPopLPush("l_key_src", "l_key_dst", 1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.LIndex("1_key", 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.LInsertBefore("1_key", "l_pivot", "l_val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.LInsertAfter("1_key", "l_pivot", "l_val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.LLen("l_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.LPop("1_key")
	if err != nil {
		t.Fatal(err)
	}
	err = r.LPush("1_key", "l_val", "l_val2")
	if err != nil {
		t.Fatal(err)
	}
	err = r.LPushX("1_key", "l_val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.LRange("1_key", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.LRem("1_key", 1, "l_val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.LSet("1_key", 0, "l_val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.LTrim("1_key", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.RPop("1_key")
	if err != nil {
		t.Fatal(err)
	}
	err = r.RPopLPush("1_key_src", "l_key_dst")
	if err != nil {
		t.Fatal(err)
	}
	err = r.RPush("1_key", "l_val", "l_val2")
	if err != nil {
		t.Fatal(err)
	}
	err = r.RPushX("1_key", "l_val")
	if err != nil {
		t.Fatal(err)
	}
}

// set command
func TestRedis_SetCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	err = r.SAdd("s_key", "s_val", "s_val2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SCard("s_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SDiff("s_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SDiffStore("s_key2", "s_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SInter("s_key", "s_key2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SInterStore("s_key3", "s_key", "s_key2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SIsMember("s_key", "s_val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SMembers("s_key")
	if err != nil {
		t.Fatal(err)
	}
	err = r.SMove("s_key", "s_key2", "s_val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SPop("s_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.SRandMember("s_key", 0)
	if err != nil {
		t.Fatal(err)
	}
	err = r.SRem("s_key", "s_val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.SUnionStore("s_key3", "s_key", "s_key2")
	if err != nil {
		t.Fatal(err)
	}
}

// zset command
func TestRedis_ZSetCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	err = r.ZAdd("z_key", 0, "val", 1, "val2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZCard("z_key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZCount("z_key", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	err = r.ZINCRBY("z_key2", 1, 0)
	if err != nil {
		t.Fatal(err)
	}
	err = r.ZInterStore("z_key3", 2, "z_key", "z_key2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZRange("z_key", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZRangeByScore("z_key", 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZRangeByScoreLimit("z_key", 0, 10, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZRank("z_key", "val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.ZRem("z_key", "val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.ZRemRangeByRank("z_key", 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	err = r.ZRemRangeByScore("z_key", 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZRevRange("z_key", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZRevRangeByScore("z_key", 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZRevRank("z_key", "val")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.ZScore("z_key", "val")
	if err != nil {
		t.Fatal(err)
	}
	err = r.ZUnionStore("z_key3", 2, "z_key", "z_key2")
	if err != nil {
		t.Fatal(err)
	}
}

// pubsub command
func TestRedis_PubSubCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	err = r.PSubscribe("*")
	if err != nil {
		t.Fatal(err)
	}
	err = r.Subscribe("1")
	if err != nil {
		t.Fatal(err)
	}
	err = r.PUnSubscribe("*")
	if err != nil {
		t.Fatal(err)
	}
	err = r.UnSubscribe("1")
	if err != nil {
		t.Fatal(err)
	}
}

// transactions command
func TestRedis_TransactionsCommand(t *testing.T) {
	var err error
	r = getRedis(t)
	defer r.Close()
	err = r.Multi()
	if err != nil {
		t.Fatal(err)
	}
	err = r.Exec()
	if err != nil {
		t.Fatal(err)
	}
	err = r.Multi()
	if err != nil {
		t.Fatal(err)
	}
	err = r.Discard()
	if err != nil {
		t.Fatal(err)
	}
}
