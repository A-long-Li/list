/**
 *@Description
 *@author          liyajun
 *@create          2022-09-09 17:06
 */

package Models

import "list/DAO"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/**
todo 模型的 crud 操作
*/

// CreateTodo  创建待办事件
func CreateTodo(todo *Todo) error {
	return DAO.DB.Create(&todo).Error
}

// GetAllList 获取待办列表
func GetAllList() (list []*Todo, err error) {
	if err = DAO.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return
}

// GetById 使用id获取待办事件 (未启用)
func GetById(id string) (todo *Todo, err error) {
	if err = DAO.DB.Debug().Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateById 更新事件
func UpdateById(todo *Todo) error {
	return DAO.DB.Save(todo).Error
}

// DeleteById 删除
func DeleteById(id string) error {
	return DAO.DB.Debug().Where("id=?", id).Delete(&Todo{}).Error
}
