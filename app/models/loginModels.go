package models

import (
	scheme "github.com/meongbego/bgin/app/moduls/migration"
	moduls "github.com/meongbego/bgin/app/moduls/package"
)

// GetAllLogin Get All data from table loginscheme
func GetAllLogin(b *[]scheme.LoginScheme) (err error) {
	if err = moduls.Conn.Find(b).Error; err != nil {
		return err
	}
	return nil
}

// AddNewLogin Save Data to table loginscheme
func AddNewLogin(b *scheme.LoginScheme) (err error) {
	if err = moduls.Conn.Create(b).Error; err != nil {
		return err
	}
	return nil
}

// GetOneLogin Get One data from table loginscheme
func GetOneLogin(b *scheme.LoginScheme, id string) (err error) {
	if err := moduls.Conn.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func GetByUsername(b *scheme.LoginScheme, username string) (err error) {
	if err := moduls.Conn.Where("username = ?", username).First(b).Error; err != nil {
		return err
	}
	return nil
}

func EditLogin(b *scheme.LoginScheme, id string) (err error) {
	moduls.Conn.Save(b)
	return nil
}

func DeleteLogin(b *scheme.LoginScheme, id string) (err error) {
	moduls.Conn.Where("id = ?", id).Delete(b)
	return nil
}
