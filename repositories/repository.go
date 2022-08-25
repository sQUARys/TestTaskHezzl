package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/sQUARys/TestTaskHezzl/proto"
	"golang.org/x/net/context"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myUser"
	password = "myPassword"
	dbname   = "myDb"

	connectionStringFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	dbInsertFormat = `INSERT INTO user_table ( "id", "name") VALUES ( %d ,'%s');`
	dbDeleteFormat = `DELETE FROM user_table WHERE id = %d;`

	clickHouseInsertFormat = `INSERT INTO log_table ( "id", "body") VALUES ( %d ,'%s');`
)

type Repository struct {
	DbStruct *sql.DB
}

func New() *Repository {
	connectionString := fmt.Sprintf(connectionStringFormat, host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	repo := Repository{
		DbStruct: db,
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &repo
}

func NewClickH() *Repository {
	db, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?username=&compress=true&debug=true")
	checkErr(err)
	checkErr(db.Ping())

	return &Repository{
		DbStruct: db,
	}
}

func (repo *Repository) AddLog(id int, log string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbInsertRequest := fmt.Sprintf(clickHouseInsertFormat, id, log)

	_, err := repo.DbStruct.ExecContext(
		ctx,
		dbInsertRequest,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) AddUser(user *pb.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbInsertRequest := fmt.Sprintf(dbInsertFormat, user.Id, user.Name)

	_, err := repo.DbStruct.ExecContext(
		ctx,
		dbInsertRequest,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) DeleteUser(id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbDeleteRequest := fmt.Sprintf(dbDeleteFormat, id)

	_, err := repo.DbStruct.ExecContext(
		ctx,
		dbDeleteRequest,
	)
	if err != nil {
		return err
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
