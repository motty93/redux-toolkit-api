package service

import (
	"app/pkg/db/model"

	"gorm.io/gorm"
)

type Task struct {
	db *gorm.DB
}

func NewServiceTask(db *gorm.DB) *Task {
	return &Task{db: db}
}

// Tasks find all
func (t *Task) Tasks() (*[]model.Task, error) {
	tasks := new([]model.Task)
	if err := t.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// Task find by id task
func (t *Task) Task(id int) (*model.Task, error) {
	task := new(model.Task)
	if err := t.db.First(&task, id).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Create task
func (t *Task) Create(task *model.Task) error {
	if err := t.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

// Update task
func (t *Task) Update(nt *model.Task, id int) (*model.Task, error) {
	var task model.Task
	if err := t.db.First(&task, id).Updates(nt).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

// Delete task
func (t *Task) Delete(id int) error {
	task := new(model.Task)
	if err := t.db.Delete(&task, id).Error; err != nil {
		return err
	}

	return nil
}
