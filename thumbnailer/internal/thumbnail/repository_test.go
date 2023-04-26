package thumbnail_test

import (
	"context"
	"testing"
	"thumbnailer/internal/thumbnail"
	"thumbnailer/pkg/ulid"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

var (
	expected = thumbnail.Thumbnail{
		ID:     ulid.NewID(),
		Type:   "png",
		Data:   []byte("test data"),
		Expiry: time.Now().Add(1 * time.Hour),
	}
)

func TestThumbnailRepositoryRedisSave(t *testing.T) {
	c := context.TODO()
	rdb, mock := redismock.NewClientMock()
	repo := thumbnail.NewRepository(rdb)

	expectedJSON, err := json.Marshal(expected)
	assert.NoError(t, err)
	expirationTime := time.Until(expected.Expiry)

	mock.ExpectSetNX(expected.ID.String(), expectedJSON, expirationTime).SetVal(true)

	err = repo.Save(c, &expected)
	assert.NoError(t, err)
}

func TestThumbnailRepositoryRedisFindByID(t *testing.T) {
	c := context.TODO()
	rdb, mock := redismock.NewClientMock()
	repo := thumbnail.NewRepository(rdb)

	expectedJSON, err := json.Marshal(expected)
	assert.NoError(t, err)
	mock.ExpectGet(expected.ID.String()).SetVal(string(expectedJSON))

	res, err := repo.FindByID(c, expected.ID)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expected.ID, res.ID)
	assert.Equal(t, expected.Type, res.Type)
	assert.Equal(t, expected.Data, res.Data)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestThumbnailRepositoryRedisFindByID_Nil(t *testing.T) {
	c := context.TODO()
	rdb, mock := redismock.NewClientMock()
	repo := thumbnail.NewRepository(rdb)

	mock.ExpectGet(expected.ID.String()).RedisNil()

	res, err := repo.FindByID(c, expected.ID)
	assert.ErrorIs(t, err, thumbnail.ErrThumbnailNotFound)
	assert.Nil(t, res)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
