package models

import (
	scheme "github.com/meongbego/bgin/app/moduls/migration"
	packages "github.com/meongbego/bgin/app/moduls/package"
)

func GetAllLogin(b *[]scheme.LoginScheme) (err error) {
	if err = packages.Conn.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewLogin(b *scheme.LoginScheme) (err error) {
	if err = packages.Conn.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneLogin(b *scheme.LoginScheme, id string) (err error) {
	if err := packages.Conn.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func GetByUsername(b *scheme.LoginScheme, username string) (err error) {
	if err := packages.Conn.Where("username = ?", username).First(b).Error; err != nil {
		return err
	}
	return nil
}

func EditLogin(b *scheme.LoginScheme, id string) (err error) {
	packages.Conn.Save(b)
	return nil
}

func DeleteLogin(b *scheme.LoginScheme, id string) (err error) {
	packages.Conn.Where("id = ?", id).Delete(b)
	return nil
}
