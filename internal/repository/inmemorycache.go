package repository

import (
	"sync"

	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
)

type InMemoryCache struct {
	sync.RWMutex
	users []domain.User
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		users: []domain.User{},
	}
}

func (i *InMemoryCache) Add(user domain.User) error {
	i.Lock()
	defer i.Unlock()

	for _, u := range i.users {
		if u.Mail == user.Mail {
			return domain.ErrUserExist
		}
	}
	i.users = append(i.users, user)
	return nil
}

func (i *InMemoryCache) Update(user domain.User) error {
	i.Lock()
	defer i.Unlock()

	for idx, u := range i.users {
		if u.Mail == user.Mail {
			i.users[idx] = user
			return nil
		}
	}
	return domain.ErrNoUser
}

func (i *InMemoryCache) Get(mail string) (domain.User, error) {
	i.Lock()
	defer i.Unlock()

	for _, u := range i.users {
		if u.Mail == mail {
			return u, nil
		}
	}
	return domain.User{}, domain.ErrNoUser
}

func (i *InMemoryCache) CheckVerifyRegString(checkMail string) (*domain.User, error) {
	i.Lock()
	defer i.Unlock()

	for _, u := range i.users {
		if u.VerifyString == checkMail {
			if u.Verified {
				return nil, domain.ErrUserVerified
			}
			return &u, nil
		}
	}
	return nil, domain.ErrUserVerifyStringNotMatch
}

func (in *InMemoryCache) LengthCache() (float64, error) {
	lenght := len(in.users)
	return float64(lenght), nil
}

func (in *InMemoryCache) LengthUnverifiedUsers() (float64, error) {
	var len float64
	for _, u := range in.users {
		if !u.Verified {
			len++
		}
	}
	return len, nil
}
