package profile

import (
	"nexus-wallet/internal/app_enum"
)

type ProfileInterface interface {
	GetId() int64
	GetProfile() Profile
	SetProfile(Profile)
	GetTableName() string
}

type Profile struct {
	Id       int64                `primary:"true" must_generate:"true" db:"id"`
	Language app_enum.Language    `db:"language"`
	Type     app_enum.ProfileType `db:"type"`
	UserId   int64                `db:"user_id"`
}

func (p *Profile) GetTableName() string {
	return "profiles"
}

type TelegramProfile struct {
	Id          int64   `primary:"true" must_generate:"true" db:"id"`
	FirstName   string  `db:"first_name"`
	LastName    string  `db:"last_name"`
	Username    string  `db:"username"`
	ProfileId   int64   `db:"profile_id"`
	ImageSource string  `db:"image_source"`
	Profile     Profile `db:"-"`
}

func (p *TelegramProfile) GetId() int64 {
	return p.Id
}

func (p *TelegramProfile) GetProfile() Profile {
	return p.Profile
}

func (p *TelegramProfile) SetProfile(profile Profile) {
	p.Profile = profile
	p.ProfileId = profile.Id
}

func (p *TelegramProfile) GetTableName() string {
	return "telegram_user_profiles"
}
