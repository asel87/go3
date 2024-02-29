
import (
	"assik3/services/contact/internal/domain/group"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	stmt := "INSERT INTO groups (name) VALUES(?)"

	_, err := r.db.Exec(stmt, group.Name())
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *Repository) ReadGroupByID(groupID int) (*group.Group, error) {
	stmt := "SELECT id, name FROM groups WHERE id = ?"

	row := r.db.QueryRow(stmt, groupID)

	var g group.Group

	err := row.Scan(g.ID(), g.Name())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &group.Group{}, fmt.Errorf("no rows found")
		} else {
			return &group.Group{}, err
		}
	}

	return &g, nil
}

func (r *Repository) UpdateGroup(groupID int, group *group.Group) (*group.Group, error) {
	stmt := `UPDATE groups SET name = ? WHERE id = ?`

	_, err := r.db.Exec(stmt, group.Name(), groupID)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r *Repository) DeleteGroup(groupID int) error {
	stmt := "DELETE FROM groups WHERE id = ?"

	_, err := r.db.Exec(stmt, groupID)
	if err != nil {
		return err
	}

	return nil
}
