package storage

import "go.uber.org/zap"

type Value struct { // значения, которые могут храниться в мапах
	s string
	d int
	a any
	b bool
}

type Storage struct {
	innerString map[string]string // создание поля для мапы
	logger      *zap.Logger       // Поле для хранения логгера
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction() // создание логгера
	if err != nil {
		return Storage{}, err // если в логгере ошибока, то возвращается storage и ошибка
	}

	logger.Info("created new storage") // логгирование

	return Storage{ // если в логгере нет ошибок, то возвращается storage
		innerString: make(map[string]string), // инициирование мапы
		logger:      logger,                  // присваивание полю из стракта Storage созданного логгера
	}, nil
}
func (r Storage) Set(key string, value string) {
	r.innerString[key] = value // присваивание value значения по ключу
	r.logger.Info("key setted")
}

func (r Storage) Get(key string) *string {
	res, ok := r.innerString[key] // res принимает значение по ключу key, а ok принимает true, если мапа не пустая
	if !ok {                      // проверка на пустую мапу
		return nil
	}
	return &res // возврат ключа по res
}
