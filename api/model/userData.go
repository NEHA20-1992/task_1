package model
import (
	"errors"
	"html"
	"log"
    "strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type UserData struct{
	ID            uint32  `gorm: "column: user_id;primary_key;auto_increment" json:"-"`
	FirstName     string  `gorm: "size:255;not null" json:"firstName"`
	LastName      string  `gorm: "size:255;not null" json:"lastName"`
	Email         string  `gorm: "size:255;not null;unique" json:"email"`
	Age           uint8   `gorm: "not null" json:"age"`
	Password      string  `gorm: "not null" json:"password"`
	ContactNumber string  `gorm: "size:255;not null" json:"contactNumber"`
}

func Hash(password string )([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
 func(u *UserData) BeforeSave()error{
	hashedPassword, err:=Hash(u.Password)
	if err!=nil{
		return err
	}
	u.Password=string(hashedPassword)
	return nil
 }

 func (u *UserData) Prepare() {
	u.ID = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *UserData) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FirstName == "" {
			return errors.New("Required First Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.FirstName == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *UserData) SaveUser(db *gorm.DB) (*UserData, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &UserData{}, err
	}
	return u, nil
}
func (u *UserData) FindAllUsers(db *gorm.DB) (*[]UserData, error) {
	var err error
	users := []UserData{}
	err = db.Debug().Model(&UserData{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]UserData{}, err
	}
	return &users, err
}
func (u *UserData) FindUserByID(db *gorm.DB, uid uint32) (*UserData, error) {
	var err error
	err = db.Debug().Model(UserData{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserData{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserData{}, errors.New("User Not Found")
	}
	return u, err
}
func (u *UserData) UpdateAUser(db *gorm.DB, uid uint32) (*UserData, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&UserData{}).Where("id = ?", uid).Take(&UserData{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"firstName":  u.FirstName,
			"email":     u.Email,
			"age":       u.Age,
			"contactNumber": u.ContactNumber,
			
		},
	)
	if db.Error != nil {
		return &UserData{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&UserData{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserData{}, err
	}
	return u, nil
}
func (u *UserData) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&UserData{}).Where("id = ?", uid).Take(&UserData{}).Delete(&UserData{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

