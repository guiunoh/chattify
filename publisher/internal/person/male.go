package person

//go:generate mockgen -destination=../mock/male_mock.go -package=mock publisher/internal/person Male
//go:generate mockgen -destination=../mock/male2_mock.go -package=mock publisher/internal/person Male2

type Male interface {
	Get(id int64) error
}

type Male2 interface {
	Get(id int64) error
}
