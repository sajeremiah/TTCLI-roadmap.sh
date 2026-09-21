package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

type TaskStatus string

const (
	StatusDone       TaskStatus = "done"
	StatusInProgress TaskStatus = "in-progress"
)

// Структура всего хранилища json-файла
type Storage struct {
	NextID int    `json:"nextId"` // счетчик для ID, чтобы избежать коллизий
	Tasks  []Task `json:"tasks"`  // массив тасков
}

// Структура таска
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (s *Storage) AddTask(description string) {
	s.Tasks = append(s.Tasks, Task{
		ID:          s.NextID,
		Description: description,
		Status:      StatusInProgress,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	})
	s.NextID++
}
func (s *Storage) UpdateTask(id int, description string) error {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks[i].Description = description
			now := time.Now()
			s.Tasks[i].UpdatedAt = &now
			return nil
		}
	}
	return fmt.Errorf("id not found: %d", id)
}
func (s *Storage) DeleteTask(id int) error {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("id not found: %d", id)
}

// Вывод тасков
func (s *Storage) ListTasks(status string) ([]Task, error) {
	cmpStatus := TaskStatus(status)
	if cmpStatus != StatusDone && cmpStatus != StatusInProgress {
		return []Task{}, fmt.Errorf("invalid list status %s", status)
	}
	var outList []Task
	for _, v := range s.Tasks {
		if v.Status == cmpStatus {
			outList = append(outList, v)
		}
	}
	return outList, nil
}

func (s *Storage) MarkTask(id int, status string) error {
	cmpStatus := TaskStatus(status)
	if cmpStatus != StatusDone && cmpStatus != StatusInProgress {
		return fmt.Errorf("invalid status %s", status)
	}
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks[i].Status = cmpStatus
			now := time.Now()
			s.Tasks[i].UpdatedAt = &now
			return nil
		}
	}
	return fmt.Errorf("id not found: %d", id)
}

// Выгрузка хранилища в оперативную память
func LoadStorage(filename string) (*Storage, error) {
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return &Storage{
				NextID: 1,
				Tasks:  []Task{},
			}, nil
		}
		return nil, err
	}
	defer file.Close()
	var storage Storage
	err = json.NewDecoder(file).Decode(&storage)
	if err != nil {
		return nil, err
	}
	if storage.NextID == 0 {
		return &Storage{
			NextID: 1,
			Tasks:  []Task{},
		}, nil
	}
	return &storage, nil
}

// Загрузка из оперативной памяти в файл
func (s *Storage) SaveStorage(filename string) error {
	// O_WRONLY - только запись, O_TRUNC - обрезать данные до 0 байт, 0666 - все юзеры могут читать и писать, но не запускать
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "   ")

	err = encoder.Encode(s)
	if err != nil {
		return err
	}
	return nil
}
