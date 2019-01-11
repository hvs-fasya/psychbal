package engine

import (
	"github.com/lib/pq"

	"github.com/hvs-fasya/chusha/internal/models"
)

//TabsGet get tabs list
func (db *PgDB) TabsGet(enabled bool) ([]*models.Tab, error) {
	var tabs []*models.Tab
	query := `SELECT t.id, t.title, t.user_type_visible, t.index, t.enabled,
				tp.id, tp.type
			FROM tabs t 
			JOIN tab_types tp ON tp.id=t.tab_type_id`
	if enabled {
		query = query + ` WHERE t.enabled = true`
	}
	query = query + ` ORDER BY t.index`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return tabs, err
	}
	defer rows.Close()
	for rows.Next() {
		tab := new(models.Tab)
		tab.TabType = new(models.TabType)
		if err := rows.Scan(
			&tab.ID,
			&tab.Title,
			pq.Array(&tab.UserTypeVisible),
			&tab.Index,
			&tab.Enabled,
			&tab.TabType.ID,
			&tab.TabType.Type,
		); err != nil {
			return tabs, err
		}
		tabs = append(tabs, tab)
	}
	return tabs, err
}

//TabsSet set tabs state
func (db *PgDB) TabsSet(tabs []*models.Tab) error {
	var q = `UPDATE tabs SET title=$1, index=$2, enabled=$3 WHERE id=$4`
	for _, tab := range tabs {
		_, e := db.Conn.Exec(q, tab.Title, tab.Index, tab.Enabled, tab.ID)
		if e != nil {
			return e
		}
	}
	return nil
}
