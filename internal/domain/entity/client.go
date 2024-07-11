package entity

import (
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/exeption"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
	"golang.org/x/crypto/bcrypt"
)

type Client struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email" gorm:"unique;not null" validate:"required,email,unique"`
	Password  string    `json:"-" gorm:"not null" validate:"required,min=12"`
	FirstName string    `json:"firstname" gorm:"not null" validate:"required,alpha"`
	LastName  string    `json:"lastname" gorm:"not null" validate:"required,alpha"`
	Orders    []Order   `json:"orders,omitempty" gorm:"foreignKey:ClientID;OnDelete:CASCADE;"`
	Age       int       `json:"age" validate:"gte=0,lte=120"`
	ID        ID        `json:"id"`
}

func NewClient(payload Client) (*Client, error) {
	client := &Client{
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
	client.Password = pass

	return client, nil
}

func (c *Client) Update(payload Client) error {
	utils.AssignIfNotEmpty(&c.FirstName, payload.FirstName)
	utils.AssignIfNotEmpty(&c.LastName, payload.LastName)
	utils.AssignIfNonZero(&c.Age, payload.Age)
	c.UpdatedAt = time.Now()

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Validate() error {
	c.validateEmail()
	c.validatePassword()
	c.validateFirstName()
	c.validateLastName()
	c.validateAge()

	return nil
}

func (c *Client) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
}

func hashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", exeption.ErrorHashingPassword
	}

	return string(hash), nil
}

func (c *Client) validateEmail() error {
	if err := utils.EnsureValueIsNotEmpty(c.Email); err != nil {
		return err
	}

	if err := utils.EnsureValueIsAValidEmailFormat(c.Email); err != nil {
		return err
	}

	return nil
}

func (c *Client) validatePassword() error {
	if err := utils.EnsureValueIsNotEmpty(c.Password); err != nil {
		return err
	}

	if err := utils.EnsureValueIsValidPasswordComplexity(c.Password); err != nil {
		return err
	}

	return nil
}

func (c *Client) validateFirstName() error {
	if err := utils.EnsureValueIsNotEmpty(c.FirstName); err != nil {
		return err
	}

	return nil
}

func (c *Client) validateLastName() error {
	if err := utils.EnsureValueIsNotEmpty(c.LastName); err != nil {
		return err
	}

	return nil
}

func (c *Client) validateAge() error {
	if err := utils.EnsureValueIsValidAge(c.Age); err != nil {
		return err
	}

	return nil
}
