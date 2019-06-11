// Copyright 2014 The mqrouter Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dbrouter

import (
	"context"
	"fmt"
	"strings"
)

func parseKey(key string) (dbType, dbName string) {
	items := strings.Split(key, "-")
	if len(items) != 2 {
		return "", ""
	}

	return items[0], items[1]
}

func generateKey(dbType, dbName string) string {
	return fmt.Sprintf("%s-%s", dbType, dbName)
}

func Factory(ctx context.Context, key string, configer Configer) (in Instancer, err error) {
	dbType, dbName := parseKey(key)

	config := configer.GetConfig(ctx, dbType, dbName)
	if len(config.DBAddr) == 0 {
		return nil, fmt.Errorf("config.DBAddr err, key: %s", key)
	}

	switch dbType {
	case DB_TYPE_MONGO:
		return NewMongo(config.DBType, config.DBName, config.UserName, config.PassWord, config.DBAddr, config.TimeOut)

	case DB_TYPE_MYSQL:
		fallthrough

	case DB_TYPE_POSTGRES:
		return NewSql(config.DBType, config.DBName, config.DBAddr[0], config.UserName, config.PassWord, config.TimeOut)

	default:
		return nil, fmt.Errorf("dbType err, key: %s", key)
	}

	return nil, fmt.Errorf("dbType err, key: %s", key)
}
