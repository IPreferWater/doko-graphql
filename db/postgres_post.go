package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/ipreferwater/doko-graphql/model"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type PostgresRepository struct {
	db *sql.DB
}

func (m PostgresRepository) GetPosts() ([]model.Post, error) {

	q, err := m.db.Query("select * from posts  ORDER BY id DESC")

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

func (n PostgresRepository) CreatePosts(newPosts []*model.InputPost) error {

	sqlStr := "INSERT INTO posts(title,txt,latitude,longitude) VALUES "
	vals := []interface{}{}

	for _, post := range newPosts {
		
		sqlStr += "(?, ?, ? ,? ),"
		vals = append(vals, post.Title, &post.Text, post.Latitude, post.Longitude)
	}

	//trim the last ,
	sqlStr = strings.TrimRight(sqlStr, ",")

	//Replacing ? with $n for postgres
	sqlStr = ReplaceSQL(sqlStr, "?")

	//prepare the statement
	stmt, err := n.db.Prepare(sqlStr)

	if err != nil {
		return err
	}

	//format all vals at once
	_, err = stmt.Exec(vals...)

	return err

	/*rowAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}*/
	//TODO return rowAffected
}

// ReplaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
	   old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
 }

func (n PostgresRepository) UpdatePosts(postsToUpdate []model.Post) error {
	return nil
}
func (n PostgresRepository) DeletePosts(idsPostToDelete []int) error {
	return nil
}

func (m PostgresRepository) GetUserIdByUsernamePassword(userName string, password string) (int, error) {

	var id int
	m.db.QueryRow("SELECT id FROM users where username = $1 AND password = $2", userName, password).Scan(&id)
	return id, nil
}

func InitPostgresRepo() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "postgres", 5432, "ipreferwater", "password", "doko")

	// open database
	client, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Errorf("can't connect to database '%s'", err)
		panic(err)
	}
	if err := client.Ping(); err != nil {
		log.Errorf("can't ping database '%s'", err)
		panic(err)
	}

	log.Info("postgre connected")

	PostRepository = PostgresRepository{db: client}
}
