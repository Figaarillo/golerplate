package entity

import (
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/exeption"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password  string    `json:"password" gorm:"not null" validate:"required,min=12"`
	FirstName string    `json:"firstname" gorm:"not null" validate:"required,alpha"`
	LastName  string    `json:"lastname" gorm:"not null" validate:"required,alpha"`
	Orders    []Order   `json:"orders,omitempty" gorm:"foreignKey:UserID;OnDelete:CASCADE;"`
	Age       int       `json:"age" validate:"gte=0,lte=120"`
	ID        ID        `json:"id"`
}

func NewUser(payload User) (*User, error) {
	user := &User{
		ID:        NewID(),
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Age:       payload.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pass, err := hashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pass

	return user, nil
}

func (c *User) Update(payload User) error {
	utils.AssignIfNotEmpty(&c.FirstName, payload.FirstName)
	utils.AssignIfNotEmpty(&c.LastName, payload.LastName)
	utils.AssignIfNonZero(&c.Age, payload.Age)
	c.UpdatedAt = time.Now()

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *User) Validate() error {
	c.validateEmail()
	c.validatePassword()
	c.validateFirstName()
	c.validateLastName()
	c.validateAge()

	return nil
}

func (c *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
}

func hashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", exeption.ErrorHashingPassword
	}

	return string(hash), nil
}

func (c *User) validateEmail() error {
	if err := utils.EnsureValueIsNotEmpty(c.Email); err != nil {
		return err
	}

	if err := utils.EnsureValueIsAValidEmailFormat(c.Email); err != nil {
		return err
	}

	return nil
}

func (c *User) validatePassword() error {
	if err := utils.EnsureValueIsNotEmpty(c.Password); err != nil {
		return err
	}

	if err := utils.EnsureValueIsValidPasswordComplexity(c.Password); err != nil {
		return err
	}

	return nil
}

func (c *User) validateFirstName() error {
	if err := utils.EnsureValueIsNotEmpty(c.FirstName); err != nil {
		return err
	}

	return nil
}

func (c *User) validateLastName() error {
	if err := utils.EnsureValueIsNotEmpty(c.LastName); err != nil {
		return err
	}

	return nil
}

func (c *User) validateAge() error {
	if err := utils.EnsureValueIsValidAge(c.Age); err != nil {
		return err
	}

	return nil
}
