package retwis

import (
	"testing"
	"strconv"
)

func TestRedis_Incr(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	ctr, err := r.Incr("test:ctr")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ctr)
}

func TestRedis_Hset(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	err = r.Hset("test:hash", "name", "wave")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedis_Hmset(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	err = r.Hmset("test:hash", "name", "wave", "age", "18")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedis_Hget(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	val, err := r.Hget("test:hash", "name")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
}

func TestRedis_Hgetall(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	vals, err := r.Hgetall("test:hash")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vals)
}

func TestRedis_Hdel(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	err = r.Hdel("test:hash", "name")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedis_Zadd(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	err = r.Zadd("test:set", 1, "first")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedis_Zcard(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	total, err := r.Zcard("test:set")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(total)
}

func TestRedis_Zscore(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	score, err := r.Zscore("test:set", "first")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(score)
}

func TestRedis_Zrem(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	err = r.Zrem("test:set", "first")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedis_Lpush(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	for i := 0; i < 5; i++ {
		err = r.Lpush("test:list", strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestRedis_Lrange(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	vals, err := r.Lrange("test:list", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vals)
}

func TestRedis_Zrange(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	vals, err := r.Zrange("test:set", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vals)
}

func TestRedis_Ltrim(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	err = r.Ltrim("test:list", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedis_Zrevrange(t *testing.T) {
	r, err := redisLink()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	vals, err := r.Zrevrange("test:set", 0, -1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vals)
}