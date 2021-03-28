package db

import (
	"database/sql"
	"fmt"
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
			ID:    id,
			Title: title,
			Txt:   &txt,
			Gps: model.Gps{
				X: latitude,
				Y: longitude,
			},
		}
		res = append(res, scannedPost)
	}

	return res, nil
}

func (n MysqlPostRepository) CreatePosts(newPosts []*model.InputPost) error {

	for _, newPost := range newPosts {
		id := len(allPosts)
		allPosts = append(allPosts, model.Post{
			ID:    id,
			Title: newPost.Title,
			Txt:   newPost.Txt,
			Gps: model.Gps{
				X: newPost.Gps.X,
				Y: newPost.Gps.Y,
			},
		})
	}
	return nil
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
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",c.User, c.Password, c.Host, c.Port, c.Database)
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
