/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gsecurity

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type redisStoreImpl struct {
	client *redis.Client
}

func NewRedisStore(r *redis.Client) Store {
	return &redisStoreImpl{client: r}
}

func (r *redisStoreImpl) Set(key string, val interface{}, exTime int64) {
	r.client.SetEx(context.Background(), key, val, time.Duration(exTime)*time.Second)
}

func (r *redisStoreImpl) Get(key string) (interface{}, bool) {
	cmd := r.client.Get(context.Background(), key)
	i, err := cmd.Int64()
	if err != nil {
		if err != redis.Nil {
			log.Println(err)
		}
		return i, false
	}

	return i, true
}

func (r *redisStoreImpl) GetObj(key string) (interface{}, bool) {
	cmd := r.client.Get(context.Background(), key)
	session := new(Session)
	if err := cmd.Scan(session); err != nil {
		if err != redis.Nil {
			log.Println(err)
		}
		return nil, false
	}

	return session, true
}

func (r *redisStoreImpl) Update(key string, val interface{}) {
	ctx := context.Background()
	exTime, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Println(err)
		}
		return
	}
	r.client.Set(ctx, key, val, exTime)
}

func (r *redisStoreImpl) Delete(key string) {
	r.client.Del(context.Background(), key)
}

func (r *redisStoreImpl) UpdateObjTimeout(key string, exTime int64) {
	if exTime == NeverExpire {
		exTime = 0
	}
	r.client.Expire(context.Background(), key, time.Duration(exTime))
}

func (r *redisStoreImpl) GetExTime(key string) int64 {
	cmd := r.client.TTL(context.Background(), key)
	duration, err := cmd.Result()
	if err != nil {
		if err != redis.Nil {
			log.Println(err)
		}
		return NotValueExist
	}

	if duration == -1 {
		return NeverExpire
	}

	return time.Now().Add(duration).Unix()
}
