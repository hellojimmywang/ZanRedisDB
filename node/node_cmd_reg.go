package node

func (nd *KVNode) registerHandler() {
	// for kv
	nd.router.Register(false, "get", wrapReadCommandK(nd.getCommand))
	nd.router.Register(false, "mget", wrapReadCommandKK(nd.mgetCommand))
	nd.router.Register(false, "exists", wrapReadCommandK(nd.existsCommand))
	nd.router.Register(true, "set", wrapWriteCommandKV(nd, nd.setCommand))
	nd.router.Register(true, "setnx", wrapWriteCommandKV(nd, nd.setnxCommand))
	nd.router.Register(true, "mset", wrapWriteCommandKVKV(nd, nd.msetCommand))
	nd.router.Register(true, "incr", wrapWriteCommandK(nd, nd.incrCommand))
	nd.router.Register(true, "del", wrapWriteCommandKK(nd, nd.delCommand))
	nd.router.Register(false, "plget", nd.plgetCommand)
	nd.router.Register(true, "plset", nd.plsetCommand)
	// for hash
	nd.router.Register(false, "hget", wrapReadCommandKSubkey(nd.hgetCommand))
	nd.router.Register(false, "hgetall", wrapReadCommandK(nd.hgetallCommand))
	nd.router.Register(false, "hkeys", wrapReadCommandK(nd.hkeysCommand))
	nd.router.Register(false, "hexists", wrapReadCommandKSubkey(nd.hexistsCommand))
	nd.router.Register(false, "hmget", wrapReadCommandKSubkeySubkey(nd.hmgetCommand))
	nd.router.Register(false, "hlen", wrapReadCommandK(nd.hlenCommand))
	nd.router.Register(true, "hset", wrapWriteCommandKSubkeyV(nd, nd.hsetCommand))
	nd.router.Register(true, "hsetnx", wrapWriteCommandKSubkeyV(nd, nd.hsetnxCommand))
	nd.router.Register(true, "hmset", wrapWriteCommandKSubkeyVSubkeyV(nd, nd.hmsetCommand))
	nd.router.Register(true, "hdel", wrapWriteCommandKSubkeySubkey(nd, nd.hdelCommand))
	nd.router.Register(true, "hincrby", wrapWriteCommandKSubkeyV(nd, nd.hincrbyCommand))
	nd.router.Register(true, "hclear", wrapWriteCommandK(nd, nd.hclearCommand))
	// for json
	nd.router.Register(false, "json.get", wrapReadCommandKSubkeySubkey(nd.jsonGetCommand))
	nd.router.Register(false, "json.keyexists", wrapReadCommandK(nd.jsonKeyExistsCommand))
	// get the same path from several json keys
	nd.router.Register(false, "json.mkget", nd.jsonmkGetCommand)
	nd.router.Register(false, "json.type", wrapReadCommandKAnySubkey(nd.jsonTypeCommand))
	nd.router.Register(false, "json.arrlen", wrapReadCommandKAnySubkey(nd.jsonArrayLenCommand))
	nd.router.Register(false, "json.objkeys", wrapReadCommandKAnySubkey(nd.jsonObjKeysCommand))
	nd.router.Register(false, "json.objlen", wrapReadCommandKAnySubkey(nd.jsonObjLenCommand))
	nd.router.Register(true, "json.set", wrapWriteCommandKSubkeyV(nd, nd.jsonSetCommand))
	nd.router.Register(true, "json.del", wrapWriteCommandKSubkeySubkey(nd, nd.jsonDelCommand))
	nd.router.Register(true, "json.arrappend", wrapWriteCommandKAnySubkey(nd, nd.jsonArrayAppendCommand, 2))
	nd.router.Register(true, "json.arrpop", wrapWriteCommandKAnySubkey(nd, nd.jsonArrayPopCommand, 0))
	// for list
	nd.router.Register(false, "lindex", wrapReadCommandKSubkey(nd.lindexCommand))
	nd.router.Register(false, "llen", wrapReadCommandK(nd.llenCommand))
	nd.router.Register(false, "lrange", wrapReadCommandKAnySubkey(nd.lrangeCommand))
	nd.router.Register(true, "lfixkey", wrapWriteCommandK(nd, nd.lfixkeyCommand))
	nd.router.Register(true, "lpop", wrapWriteCommandK(nd, nd.lpopCommand))
	nd.router.Register(true, "lpush", wrapWriteCommandKVV(nd, nd.lpushCommand))
	nd.router.Register(true, "lset", nd.lsetCommand)
	nd.router.Register(true, "ltrim", nd.ltrimCommand)
	nd.router.Register(true, "rpop", wrapWriteCommandK(nd, nd.rpopCommand))
	nd.router.Register(true, "rpush", wrapWriteCommandKVV(nd, nd.rpushCommand))
	nd.router.Register(true, "lclear", wrapWriteCommandK(nd, nd.lclearCommand))
	// for zset
	nd.router.Register(false, "zscore", wrapReadCommandKSubkey(nd.zscoreCommand))
	nd.router.Register(false, "zcount", wrapReadCommandKAnySubkey(nd.zcountCommand))
	nd.router.Register(false, "zcard", wrapReadCommandK(nd.zcardCommand))
	nd.router.Register(false, "zlexcount", wrapReadCommandKAnySubkey(nd.zlexcountCommand))
	nd.router.Register(false, "zrange", wrapReadCommandKAnySubkey(nd.zrangeCommand))
	nd.router.Register(false, "zrevrange", wrapReadCommandKAnySubkey(nd.zrevrangeCommand))
	nd.router.Register(false, "zrangebylex", wrapReadCommandKAnySubkey(nd.zrangebylexCommand))
	nd.router.Register(false, "zrangebyscore", wrapReadCommandKAnySubkey(nd.zrangebyscoreCommand))
	nd.router.Register(false, "zrevrangebyscore", wrapReadCommandKAnySubkey(nd.zrevrangebyscoreCommand))
	nd.router.Register(false, "zrank", wrapReadCommandKSubkey(nd.zrankCommand))
	nd.router.Register(false, "zrevrank", wrapReadCommandKSubkey(nd.zrevrankCommand))
	nd.router.Register(true, "zadd", nd.zaddCommand)
	nd.router.Register(true, "zincrby", nd.zincrbyCommand)
	nd.router.Register(true, "zrem", wrapWriteCommandKSubkeySubkey(nd, nd.zremCommand))
	nd.router.Register(true, "zremrangebyrank", nd.zremrangebyrankCommand)
	nd.router.Register(true, "zremrangebyscore", nd.zremrangebyscoreCommand)
	nd.router.Register(true, "zremrangebylex", nd.zremrangebylexCommand)
	nd.router.Register(true, "zclear", wrapWriteCommandK(nd, nd.zclearCommand))
	// for set
	nd.router.Register(false, "scard", wrapReadCommandK(nd.scardCommand))
	nd.router.Register(false, "sismember", wrapReadCommandKSubkey(nd.sismemberCommand))
	nd.router.Register(false, "smembers", wrapReadCommandK(nd.smembersCommand))
	nd.router.Register(true, "sadd", wrapWriteCommandKSubkeySubkey(nd, nd.saddCommand))
	nd.router.Register(true, "srem", wrapWriteCommandKSubkeySubkey(nd, nd.sremCommand))
	nd.router.Register(true, "sclear", wrapWriteCommandK(nd, nd.sclearCommand))
	nd.router.Register(true, "smclear", wrapWriteCommandKK(nd, nd.smclearCommand))
	// for ttl
	nd.router.Register(false, "ttl", wrapReadCommandK(nd.ttlCommand))
	nd.router.Register(false, "httl", wrapReadCommandK(nd.httlCommand))
	nd.router.Register(false, "lttl", wrapReadCommandK(nd.lttlCommand))
	nd.router.Register(false, "sttl", wrapReadCommandK(nd.sttlCommand))
	nd.router.Register(false, "zttl", wrapReadCommandK(nd.zttlCommand))

	nd.router.Register(true, "setex", wrapWriteCommandKVV(nd, nd.setexCommand))
	nd.router.Register(true, "expire", wrapWriteCommandKV(nd, nd.expireCommand))
	nd.router.Register(true, "hexpire", wrapWriteCommandKV(nd, nd.hashExpireCommand))
	nd.router.Register(true, "lexpire", wrapWriteCommandKV(nd, nd.listExpireCommand))
	nd.router.Register(true, "sexpire", wrapWriteCommandKV(nd, nd.setExpireCommand))
	nd.router.Register(true, "zexpire", wrapWriteCommandKV(nd, nd.zsetExpireCommand))

	nd.router.Register(true, "persist", wrapWriteCommandK(nd, nd.persistCommand))
	nd.router.Register(true, "hpersist", wrapWriteCommandK(nd, nd.persistCommand))
	nd.router.Register(true, "lpersist", wrapWriteCommandK(nd, nd.persistCommand))
	nd.router.Register(true, "spersist", wrapWriteCommandK(nd, nd.persistCommand))
	nd.router.Register(true, "zpersist", wrapWriteCommandK(nd, nd.persistCommand))

	// for scan
	nd.router.Register(false, "hscan", wrapReadCommandKAnySubkey(nd.hscanCommand))
	nd.router.Register(false, "sscan", wrapReadCommandKAnySubkey(nd.sscanCommand))
	nd.router.Register(false, "zscan", wrapReadCommandKAnySubkey(nd.zscanCommand))

	// for geohash
	nd.router.Register(true, "geoadd", nd.geoaddCommand)
	nd.router.Register(false, "geohash", wrapReadCommandKAnySubkeyN(nd.geohashCommand, 1))
	nd.router.Register(false, "geodist", wrapReadCommandKAnySubkey(nd.geodistCommand))
	nd.router.Register(false, "geopos", wrapReadCommandKAnySubkeyN(nd.geoposCommand, 1))
	nd.router.Register(false, "georadius", wrapReadCommandKAnySubkeyN(nd.geoRadiusCommand, 4))
	nd.router.Register(false, "georadiusbymember", wrapReadCommandKAnySubkeyN(nd.geoRadiusByMemberCommand, 3))

	//for cross mutil partion
	nd.router.RegisterMerge("scan", wrapMergeCommand(nd.scanCommand))
	nd.router.RegisterMerge("advscan", nd.advanceScanCommand)
	nd.router.RegisterMerge("fullscan", nd.fullScanCommand)
	nd.router.RegisterMerge("hidx.from", nd.hindexSearchCommand)

	// only write command need to be registered as internal
	// kv
	nd.router.RegisterInternal("del", nd.localDelCommand)
	nd.router.RegisterInternal("set", nd.localSetCommand)
	nd.router.RegisterInternal("setnx", nd.localSetnxCommand)
	nd.router.RegisterInternal("mset", nd.localMSetCommand)
	nd.router.RegisterInternal("incr", nd.localIncrCommand)
	nd.router.RegisterInternal("plset", nd.localPlsetCommand)
	// hash
	nd.router.RegisterInternal("hset", nd.localHSetCommand)
	nd.router.RegisterInternal("hsetnx", nd.localHSetNXCommand)
	nd.router.RegisterInternal("hmset", nd.localHMsetCommand)
	nd.router.RegisterInternal("hdel", nd.localHDelCommand)
	nd.router.RegisterInternal("hincrby", nd.localHIncrbyCommand)
	nd.router.RegisterInternal("hclear", nd.localHclearCommand)
	nd.router.RegisterInternal("hmclear", nd.localHMClearCommand)
	// for json
	nd.router.RegisterInternal("json.set", nd.localJSONSetCommand)
	nd.router.RegisterInternal("json.del", nd.localJSONDelCommand)
	nd.router.RegisterInternal("json.arrappend", nd.localJSONArrayAppendCommand)
	nd.router.RegisterInternal("json.arrpop", nd.localJSONArrayPopCommand)
	// list
	nd.router.RegisterInternal("lfixkey", nd.localLfixkeyCommand)
	nd.router.RegisterInternal("lpop", nd.localLpopCommand)
	nd.router.RegisterInternal("lpush", nd.localLpushCommand)
	nd.router.RegisterInternal("lset", nd.localLsetCommand)
	nd.router.RegisterInternal("ltrim", nd.localLtrimCommand)
	nd.router.RegisterInternal("rpop", nd.localRpopCommand)
	nd.router.RegisterInternal("rpush", nd.localRpushCommand)
	nd.router.RegisterInternal("lclear", nd.localLclearCommand)
	nd.router.RegisterInternal("lmclear", nd.localLMClearCommand)
	// zset
	nd.router.RegisterInternal("zadd", nd.localZaddCommand)
	nd.router.RegisterInternal("zincrby", nd.localZincrbyCommand)
	nd.router.RegisterInternal("zrem", nd.localZremCommand)
	nd.router.RegisterInternal("zremrangebyrank", nd.localZremrangebyrankCommand)
	nd.router.RegisterInternal("zremrangebyscore", nd.localZremrangebyscoreCommand)
	nd.router.RegisterInternal("zremrangebylex", nd.localZremrangebylexCommand)
	nd.router.RegisterInternal("zclear", nd.localZclearCommand)
	nd.router.RegisterInternal("zmclear", nd.localZMClearCommand)
	// set
	nd.router.RegisterInternal("sadd", nd.localSadd)
	nd.router.RegisterInternal("srem", nd.localSrem)
	nd.router.RegisterInternal("sclear", nd.localSclear)
	nd.router.RegisterInternal("smclear", nd.localSmclear)
	// expire
	nd.router.RegisterInternal("setex", nd.localSetexCommand)
	nd.router.RegisterInternal("expire", nd.localExpireCommand)
	nd.router.RegisterInternal("lexpire", nd.localListExpireCommand)
	nd.router.RegisterInternal("hexpire", nd.localHashExpireCommand)
	nd.router.RegisterInternal("sexpire", nd.localSetExpireCommand)
	nd.router.RegisterInternal("zexpire", nd.localZSetExpireCommand)
	nd.router.RegisterInternal("persist", nd.localPersistCommand)
	nd.router.RegisterInternal("hpersist", nd.localHashPersistCommand)
	nd.router.RegisterInternal("lpersist", nd.localListPersistCommand)
	nd.router.RegisterInternal("spersist", nd.localSetPersistCommand)
	nd.router.RegisterInternal("zpersist", nd.localZSetPersistCommand)
}
