package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ipreferwater/graphql-theory/config"
	"github.com/ipreferwater/graphql-theory/model"
	log "github.com/sirupsen/logrus"
)

var (
	PostRepository PostRepositoryInterface
	allPosts       []model.Post
)

type MysqlPostRepository struct {
	client *sql.DB
}

func (m MysqlPostRepository) GetPosts() ([]model.Post, error) {
	q, err := m.client.Query("select * from posts  ORDER BY id DESC")

	if err != nil {
		log.Errorf("getPost query %s", err)
		return nil, err
	}
	var res []model.Post

	for q.Next() {
		var id int
		var title, txt string
		var latitude, longitude float64
		if err := q.Scan(&id, &title, &txt, &latitude, &longitude); err != nil {
			log.Errorf("getPost scan %s", err)
			return nil, err
		}
		scannedPost := model.Post{
			ID:        id,
			Title:     title,
			Text:      &txt,
			Latitude:  latitude,
			Longitude: longitude,
		}
		res = append(res, scannedPost)
	}

	return res, nil
}

func (n MysqlPostRepository) CreatePosts(newPosts []*model.InputPost) error {
	sqlStr := "INSERT INTO posts(title,txt,latitude,longitude) VALUES "
	vals := []interface{}{}

	for _, post := range newPosts {
		sqlStr += "(?, ?, ?,?),"
		vals = append(vals, post.Title, post.Text, post.Latitude, post.Longitude)
	}

	//trim the last ,
	sqlStr = strings.TrimRight(sqlStr, ",")
	
	//prepare the statement
	stmt, err := n.client.Prepare(sqlStr)

	if err != nil {
		return err
	}

	//format all vals at once
	res, err := stmt.Exec(vals...)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	return err
}

func (n MysqlPostRepository) UpdatePosts(postsToUpdate []model.Post) error {
	return nil
}
func (n MysqlPostRepository) DeletePosts(idsPostToDelete []int) error {
	return nil
}

func (m MysqlPostRepository) GetUserIdByUsernamePassword(userName string, password string) (int, error) {

	var id int
	m.client.QueryRow("SELECT id FROM users where username = ? AND password = ?", userName, password).Scan(&id)
	return id, nil
}

func InitMysqlPostRepository() {
	c := config.Mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("ping failed: ", err.Error)
	}

	log.Info("mysql connected")

	// See "Important settings" section.
	//TODO
	db.SetConnMaxLifetime(time.Hour * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	PostRepository = &MysqlPostRepository{client: db}
}
