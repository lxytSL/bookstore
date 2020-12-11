package dao

import (
	"web/bookstore/model"
	"web/bookstore/utils"
)

// AddSession 添加session
func AddSession(ses *model.Session) error {
	sql := "insert into session value(?,?,?)"
	_, err := utils.Db.Exec(sql, ses.SessionID, ses.UserName, ses.UserID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession 删除session
func DeleteSession(sesID string) error {
	sql := "delete from session where session_id = ?"
	_, err := utils.Db.Exec(sql, sesID)
	if err != nil {
		return err
	}
	return nil
}

// GetSession 获取session
func GetSession(uuid string) (*model.Session, error) {
	session := &model.Session{}
	sql := "select * from session where session_id=?"
	err := utils.Db.QueryRow(sql, uuid).Scan(&session.SessionID, &session.UserName, &session.UserID)
	if err != nil {
		return nil, err
	}
	return session, nil

}
