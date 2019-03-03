package observable

import (
	"monita-backend/pkg/webpage"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

// Observable represents observable record data
type Observable struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Periodicity string `json:"periodicity"`
	URL         string `json:"url"`
	Selector    string `json:"selector"`
	LastData    string `json:"lastData"`
	Mute        bool   `json:"mute"`

	Order int `json:"order" gorm:"AUTO_INCREMENT"`

	UserID uint `json:"-"`
}

// Init run initialization code for Observable model
func Init(db *gorm.DB) {
	database = db

	db.AutoMigrate(&Observable{})
}

// GetByID returns Observable by provided id
func GetByID(id uint) *Observable {
	o := Observable{}

	database.First(&o, id)

	if o.ID == 0 {
		return nil
	}

	return &o
}

// DeleteByID deletes Observable by provided id
func DeleteByID(id uint) {
	o := GetByID(id)

	if o != nil {
		database.Delete(o)
	}
}

// CreatePayload provides data for Observable creation
type CreatePayload struct {
	Name string
	// manually, regularly, daily, weekly
	Periodicity string
	// URL with query params or hash
	URL string
	// Selector contains any JQuery selector
	Selector string
	// UserID of user that owns the observable
	UserID uint
}

// Create save new Observable with type/name/url/selector
func Create(p CreatePayload) (*Observable, error) {
	data, err := loadData(p.URL, p.Selector)

	if err != nil {
		return nil, err
	}

	o := Observable{
		Name:        p.Name,
		Periodicity: p.Periodicity,
		URL:         p.URL,
		Selector:    p.Selector,
		LastData:    data,
		UserID:      p.UserID,
	}

	database.Create(&o)

	return &o, err
}

// LoadData gets new data for Observable
func (o *Observable) LoadData() (string, error) {
	return loadData(o.URL, o.Selector)
}

func loadData(url, selector string) (string, error) {
	w, err := webpage.Load(url)

	if err != nil {
		return "", err
	}

	return w.Select(selector), nil
}

// Handle update the LastData of Observable
func (o *Observable) Handle() error {
	data, err := o.LoadData()

	if err != nil {
		return err
	}

	o.LastData = data

	database.Save(o)

	return nil
}

// MuteNotifications disable User email notifications
func (o *Observable) MuteNotifications() {
	o.Mute = true

	database.Save(o)
}

// UnmuteNotifications enable User email notifications
func (o *Observable) UnmuteNotifications() {
	o.Mute = false

	database.Save(o)
}

// Reorder change order of Observables
func (o *Observable) Reorder(order int) {
	observables := []Observable{}

	database.Find(&observables, "id <> ? AND user_id = ? AND order >= ?", o.ID, o.UserID, order)

	if len(observables) > 0 {
		c := order

		for _, o := range observables {
			if o.Order != c {
				break
			}

			o.Order++

			c = o.Order

			database.Save(&o)
		}
	}

	o.Order = order

	database.Save(o)
}
