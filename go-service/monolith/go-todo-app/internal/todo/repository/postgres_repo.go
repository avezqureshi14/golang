package todo

import (
	"database/sql"
	models "go-todo-app/internal/todo/models"
	appErr "go-todo-app/pkg/errors"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TodoRepo struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewTodoRepo(db *sqlx.DB, log *zap.Logger) *TodoRepo {
	return &TodoRepo{
		db:  db,
		log: log,
	}
}

func (r *TodoRepo) GetAll() ([]models.Todo, error) {
	todos := []models.Todo{}
	query := `SELECT id, title, done FROM todos`

	err := r.db.Select(&todos, query)
	if err != nil {
		r.log.Error("failed to fetch todos",
			zap.Error(err),
		)
		return nil, err
	}

	r.log.Info("fetched todos",
		zap.Int("count", len(todos)),
	)

	return todos, nil
}

func (r *TodoRepo) Get(id int) (models.Todo, error) {
	var todo models.Todo
	query := `SELECT id, title, done FROM todos WHERE id=$1`

	err := r.db.Get(&todo, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			r.log.Warn("todo not found",
				zap.Int("id", id),
			)
			return models.Todo{}, appErr.ErrNotFound
		}

		r.log.Error("failed to fetch todo",
			zap.Int("id", id),
			zap.Error(err),
		)
		return models.Todo{}, err
	}

	r.log.Info("fetched todo",
		zap.Int("id", id),
	)

	return todo, nil
}

func (r *TodoRepo) Create(title string) (models.Todo, error) {
	var todo models.Todo

	query := `
		INSERT INTO todos (title, done)
		VALUES ($1, false)
		RETURNING id, title, done
	`

	err := r.db.Get(&todo, query, title)
	if err != nil {
		r.log.Error("failed to create todo",
			zap.String("title", title),
			zap.Error(err),
		)
		return models.Todo{}, err
	}

	r.log.Info("created todo",
		zap.Int("id", todo.ID),
		zap.String("title", todo.Title),
	)

	return todo, nil
}

func (r *TodoRepo) Update(id int, title string, done bool) (models.Todo, error) {
	var todo models.Todo

	query := `
		UPDATE todos
		SET title=$1, done=$2
		WHERE id=$3
		RETURNING id, title, done
	`

	err := r.db.Get(&todo, query, title, done, id)
	if err != nil {
		if err == sql.ErrNoRows {
			r.log.Warn("todo not found for update",
				zap.Int("id", id),
			)
			return models.Todo{}, appErr.ErrNotFound
		}

		r.log.Error("failed to update todo",
			zap.Int("id", id),
			zap.Error(err),
		)
		return models.Todo{}, err
	}

	r.log.Info("updated todo",
		zap.Int("id", id),
		zap.String("title", title),
		zap.Bool("done", done),
	)

	return todo, nil
}

func (r *TodoRepo) Delete(id int) error {
	query := `DELETE FROM todos WHERE id=$1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error("failed to delete todo",
			zap.Int("id", id),
			zap.Error(err),
		)
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		r.log.Error("failed to get rows affected",
			zap.Int("id", id),
			zap.Error(err),
		)
		return err
	}

	if rows == 0 {
		r.log.Warn("todo not found for delete",
			zap.Int("id", id),
		)
		return appErr.ErrNotFound
	}

	r.log.Info("deleted todo",
		zap.Int("id", id),
	)

	return nil
}