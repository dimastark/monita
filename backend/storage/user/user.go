package user

import (
	"errors"
	"log"

	"monita/pkg/mail"
	"monita/storage/observable"

	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var database *gorm.DB

// User represents user data
type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	HashedPassword []byte `json:"-"`
	Email          string `json:"email"`
	Mute           bool   `json:"mute"`

	Observables []observable.Observable `json:"-"`
}

// Init run initialization code for User model
func Init(db *gorm.DB) {
	db.AutoMigrate(&User{})

	db.Model(&User{}).
		AddUniqueIndex("unique_name", "name").
		AddUniqueIndex("unique_email", "email")

	database = db
}

// GetAll returns all users
func GetAll() []User {
	users := []User{}

	database.Find(&users)

	return users
}

// GetByID returns User by provided id
func GetByID(id uint) *User {
	u := User{}

	database.First(&u, id)

	if u.ID == 0 {
		return nil
	}

	return &u
}

// GetByName returns User by provided name
func GetByName(name string) *User {
	u := User{}

	database.Where("name = ?", name).First(&u)

	if u.ID == 0 {
		return nil
	}

	return &u
}

// GetByEmail returns User by provided email
func GetByEmail(email string) *User {
	u := User{}

	database.Where("email = ?", email).First(&u)

	if u.ID == 0 {
		return nil
	}

	return &u
}

// GetByNameOrEmail returns User by provided name and email
func GetByNameOrEmail(name, email string) *User {
	u := User{}

	database.Where("name = ? OR email = ?", name, email).First(&u)

	if u.ID == 0 {
		return nil
	}

	return &u
}

// GetByLogin returns User by provided login
func GetByLogin(login string) *User {
	return GetByNameOrEmail(login, login)
}

// DeleteByName deletes User by provided name
func DeleteByName(name string) {
	if u := GetByName(name); u != nil {
		database.Delete(&u)
	}
}

// Create creates new user with name/password/email
func Create(name, password, email string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	u := User{Name: name, HashedPassword: hashedPassword, Email: email}

	database.Create(&u)

	return &u, nil
}

// Login returns User by login (name or email) and password
func Login(login string, password string) *User {
	u := GetByLogin(login)

	if u == nil {
		return nil
	}

	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(password))

	if err != nil {
		return nil
	}

	return u
}

// ChangePassword updates User password
func (u *User) ChangePassword(newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("password hashing error")
	}

	u.HashedPassword = hashedPassword

	database.Save(u)

	return nil
}

// ChangeEmail updates User email
func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail

	database.Save(u)
}

// MuteNotifications disable User email notifications
func (u *User) MuteNotifications() {
	u.Mute = true

	database.Save(u)
}

// UnmuteNotifications enable User email notifications
func (u *User) UnmuteNotifications() {
	u.Mute = false

	database.Save(u)
}

// GetObservables returns User related Observables
func (u *User) GetObservables() []observable.Observable {
	observables := []observable.Observable{}

	database.
		Order(`"order"`).
		Find(&observables, "user_id = ?", u.ID)

	return observables
}

// GetObservablesWithPeriodicity returns User related Observables with concrete periodicity
func (u *User) GetObservablesWithPeriodicity(periodicity string) []observable.Observable {
	observables := []observable.Observable{}

	database.
		Order(`"order"`).
		Find(&observables, "user_id = ? AND periodicity = ?", u.ID, periodicity)

	return observables
}

// Worker update observables automatically
func Worker() {
	gocron.Every(1).Minute().Do(worker, "regularly")
	gocron.Every(1).Day().At("9:00").Do(worker, "daily")
	gocron.Every(1).Monday().At("9:00").Do(worker, "weekly")

	<-gocron.Start()
}

func worker(periodicity string) {
	for _, u := range GetAll() {
		data := mail.TemplateData{
			ReportType: periodicity,
			Updates:    make([]mail.UpdateData, 0),
		}

		for _, o := range u.GetObservablesWithPeriodicity(periodicity) {
			oldData := o.LastData

			o.Handle()

			newData := o.LastData

			if oldData != newData {
				data.Updates = append(data.Updates, mail.UpdateData{
					ID:      o.ID,
					Name:    o.Name,
					OldData: oldData,
					NewData: newData,
				})
			}
		}

		if len(data.Updates) != 0 {
			err := mail.Send(u.Email, data)

			if err != nil {
				log.Println("error while send:", data)
			}
		}
	}
}
