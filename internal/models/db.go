package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) CreateMenu(newMenu CreateMenu) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into menu (name, type, memo, fileString, create_at, update_at) 
		values (?, ?, ?, ?, ?, ?)
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		newMenu.Name,
		newMenu.Type,
		newMenu.Memo,
		newMenu.FileString,
		newMenu.CreatedAt,
		newMenu.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) GetAllMenu() ([]*Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT m.id, m.name, m.type, m.memo, m.fileString, m.create_at, m.update_at, m.rating,
		CASE
			WHEN MAX(om.close_at) > CURRENT_DATE THEN 1
			ELSE 0
		END AS is_open
		FROM
			menu m
		LEFT JOIN
			open_menu om
		ON
			m.id = om.menu_id
		GROUP BY
			m.id,
			m.name,    
			m.type,
			m.memo,
			m.fileString,
			m.create_at,
			m.update_at,
			m.rating,
			m.vote_count;
		`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []*Menu
	for rows.Next() {
		var menu Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Type,
			&menu.Memo,
			&menu.FileString,
			&menu.CreatedAt,
			&menu.UpdatedAt,
			&menu.Rating,
			&menu.IsOpen,
		)
		if err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}

	return menus, nil
}

func (m *DBModel) GetOpenMenu() ([]*OpenedMenu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT
		om.id,
		om.menu_id,
		m.name,
		m.type,
		m.memo,
		m.fileString,
		om.create_user,
		om.create_dept,
		om.open_at,
		om.close_at,
		IFNULL(o.order_count, 0) AS order_count,
		IFNULL(o.order_total_price, 0) AS order_total_price
	FROM
		open_menu om
	LEFT JOIN
		menu m ON m.id = om.menu_id
	LEFT JOIN (
		SELECT
			open_menu_id,
			COUNT(*) AS order_count,
			SUM(price) AS order_total_price
		FROM
			orders
		GROUP BY
			open_menu_id
	) o ON om.id = o.open_menu_id
	WHERE
		om.close_at > CURRENT_DATE;
			`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []*OpenedMenu
	for rows.Next() {
		var menu OpenedMenu
		err := rows.Scan(
			&menu.ID,
			&menu.MenuID,
			&menu.Name,
			&menu.Type,
			&menu.Memo,
			&menu.FileString,
			&menu.CreateUser,
			&menu.CreateDept,
			&menu.OpenAt,
			&menu.CloseAt,
			&menu.OrderCount,
			&menu.OrderTotalPrice,
		)
		if err != nil {
			return nil, err
		}
		menus = append(menus, &menu)
	}

	return menus, nil
}

func (m *DBModel) OpenMenu(openMenu OpenMenu) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into open_menu (menu_id, create_user, create_dept, open_at, close_at) 
	values (?, ?, ?, CURRENT_DATE, ?)`

	_, err := m.DB.ExecContext(ctx, stmt,
		openMenu.ID,
		openMenu.CreateUser,
		openMenu.CreateDept,
		openMenu.CloseAt,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) UpdateMenu(menu Menu) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `update menu set name = ?, type = ?, memo = ?, fileString = ?, update_at = ? where id = ?`
	_, err := m.DB.ExecContext(ctx, stmt,
		menu.Name,
		menu.Type,
		menu.Memo,
		menu.FileString,
		menu.UpdatedAt,
		menu.ID,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) UpdateMenuRating(id int, score float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	selectStmt := `select id, COALESCE(rating, 0), COALESCE(vote_count,0) from menu where id = ?`
	row := m.DB.QueryRowContext(ctx, selectStmt, id)

	var rating Rating

	err := row.Scan(
		&rating.ID,
		&rating.Rating,
		&rating.VoteCount,
	)
	if err != nil {
		return err
	}

	newVoteCount := rating.VoteCount + 1
	newRating := (rating.Rating*float64(rating.VoteCount) + score) / float64(newVoteCount)

	updateStmt := `update menu set rating = ?, vote_count = ? where id = ?
	`
	_, err = m.DB.ExecContext(ctx, updateStmt, newRating, newVoteCount, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) AddOrder(order Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into orders (open_menu_id, name, type, item, sugar, ice, user_memo, user_name, update_at, price, count)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		order.OpenMenuID,
		order.Name,
		order.Type,
		order.Item,
		order.Sugar,
		order.Ice,
		order.UserMemo,
		order.User,
		order.UpdateAt,
		order.Price,
		order.Count,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) AllOrder(openMenuId int) ([]*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select o.id, o.open_menu_id, o.name, o.type, o.item, o.sugar, o.ice, o.user_memo, o.user_name, o.update_at, o.price, o.count from orders o
		where o.open_menu_id = ?
	`
	rows, err := m.DB.QueryContext(ctx, query, openMenuId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.OpenMenuID,
			&order.Name,
			&order.Type,
			&order.Item,
			&order.Sugar,
			&order.Ice,
			&order.UserMemo,
			&order.User,
			&order.UpdateAt,
			&order.Price,
			&order.Count,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (m *DBModel) UpdateOrder(order Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update orders set item = ?, sugar = ?, ice = ?, price = ?, user_memo = ?, update_at = ?, count = ? where id = ? and open_menu_id = ?
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		order.Item,
		order.Sugar,
		order.Ice,
		order.Price,
		order.UserMemo,
		order.UpdateAt,
		order.Count,
		order.ID,
		order.OpenMenuID,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *DBModel) DeleteOrder(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from orders where id = ?
	`
	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
