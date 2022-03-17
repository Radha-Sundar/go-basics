package db

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	fmt.Println(Get().Exec(fmt.Sprintf("CREATE TABLE test_db (user_id serial PRIMARY KEY,username VARCHAR ( 50 ) UNIQUE NOT NULL);")))
}
