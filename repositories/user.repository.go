package repositories

import "go-test-api/models"

type UserRepository interface {
	Create(user models.User) error
	CheckUserID(id int) (bool, error)
	CheckUsername(username string) (bool, error)
	List() ([]models.User, error)
	Get(id int) (models.User, error)
	Update(user models.User) error
	Delete(id int) error
}

type userRepository struct {
	data map[int]models.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		data: make(map[int]models.User),
	}
}

func (r *userRepository) Create(user models.User) error {
	r.data[user.ID] = user
	return nil
}

func (r *userRepository) CheckUserID(id int) (bool, error) {
	_, ok := r.data[id]
	return ok, nil
}

func (r *userRepository) CheckUsername(username string) (bool, error) {
	for _, user := range r.data {
		if user.Username == username {
			return true, nil
		}
	}
	return false, nil
}

func (r *userRepository) List() ([]models.User, error) {
	users := make([]models.User, 0, len(r.data))
	for _, user := range r.data {
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Get(id int) (models.User, error) {
	user, ok := r.data[id]
	if !ok {
		return models.User{}, nil
	}
	return user, nil
}

func (r *userRepository) Update(user models.User) error {
	r.data[user.ID] = user
	return nil
}

func (r *userRepository) Delete(id int) error {
	delete(r.data, id)
	return nil
}
