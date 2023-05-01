package connection_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"

	"notifier/internal/connection"
	"notifier/pkg/ulid"
)

var (
	rdb, mock = redismock.NewClientMock()
	repo      = connection.NewRepository(rdb, "test")
	ctx       = context.TODO()

	conn = &connection.Connection{
		ID:       ulid.NewID(),
		CreateAt: time.Now(),
		ExpiryAt: time.Now().Add(time.Minute),
	}
	key             = repo.Key(conn.ID)
	expiration      = conn.ExpiryAt.Sub(conn.CreateAt)
	expectedJson, _ = json.Marshal(conn)
)

func TestRepositorySave(t *testing.T) {
	mock.ExpectSetNX(key, expectedJson, expiration).SetVal(true)
	_ = repo.Save(ctx, conn)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositorySaveDuplicateKey(t *testing.T) {
	mock.ExpectSetNX(key, expectedJson, expiration).SetVal(false)
	err := repo.Save(ctx, conn)
	assert.Error(t, err, connection.ErrDuplicateKey)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositorySaveNotFound(t *testing.T) {
	mock.ExpectGet(key).RedisNil()
	result, err := repo.FetchOne(ctx, conn.ID)
	assert.Error(t, err, connection.ErrNotFound)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryExtendTTL(t *testing.T) {
	mock.ExpectSetNX(key, expectedJson, expiration).SetVal(true)
	err := repo.Save(ctx, conn)
	assert.NoError(t, err)

	mock.ExpectExpire(key, time.Minute).SetVal(true)
	_ = repo.ExtendTTL(ctx, conn.ID, time.Minute)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestRepositoryDelete(t *testing.T) {
	mock.ExpectDel(key).SetVal(1)
	_ = repo.Delete(ctx, conn.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFetchOne(t *testing.T) {
	mock.ExpectGet(key).SetVal(string(expectedJson))
	result, err := repo.FetchOne(ctx, conn.ID)
	assert.NoError(t, err)
	assert.Equal(t, conn.ID, result.ID)
	assert.Equal(t, conn.Host, result.Host)
	assert.NoError(t, mock.ExpectationsWereMet())
}
