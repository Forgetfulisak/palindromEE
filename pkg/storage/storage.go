package storage

import (
	"database/sql"
	"log"
	"palindromee/pkg/userservice"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const DATABASE_DRIVER = "sqlite3"

type Storage struct {
	db *sql.DB
}

type dbUser struct {
	id        uuid.UUID
	firstname string
	lastname  string
}

func (s *Storage) initUserTable() error {
	stmt := `
	create table if not exists users (
		id TEXT not null Primary key, 
		firstname TEXT not null,
		lastname TEXT not null
	);
	`
	_, err := s.db.Exec(stmt)
	return err
}

func New(dbPath string) (*Storage, error) {
	db, err := sql.Open(DATABASE_DRIVER, dbPath)
	if err != nil {
		return nil, err
	}

	store := &Storage{
		db: db,
	}

	store.initUserTable()

	return store, err
}

func (s *Storage) GetUser(userID string) (userservice.User, error) {

	id, err := uuid.Parse(userID)
	if err != nil {
		return userservice.User{}, err
	}

	row := s.db.QueryRow(`
	SELECT id, firstname, lastname
	FROM users
	WHERE id = ?;`, id)

	user := dbUser{}
	err = row.Scan(&user.id, &user.firstname, &user.lastname)
	if err != nil {
		return userservice.User{}, err
	}

	return userservice.User{
		UserID:    user.id.String(),
		FirstName: user.firstname,
		LastName:  user.lastname,
	}, err

}

func (s *Storage) InsertUser(newUser userservice.User) (userservice.User, error) {

	user := dbUser{
		id:        uuid.New(),
		firstname: newUser.FirstName,
		lastname:  newUser.LastName,
	}

	_, err := s.db.Exec(`INSERT into users(id, firstname, lastname) values(?, ?, ?);`,
		user.id, user.firstname, user.lastname)

	newUser.UserID = user.id.String()
	log.Println("New user is inserted. err: ", err)

	return newUser, err
}

func (s *Storage) UpdateUser(newUser userservice.User) (userservice.User, error) {

	id, err := uuid.Parse(newUser.UserID)
	if err != nil {
		return userservice.User{}, err
	}

	_, err = s.db.Exec(`
		UPDATE users 
		SET firstname = ?, lastname = ?
		WHERE id = ?;`,
		newUser.FirstName, newUser.LastName,
		id,
	)

	return newUser, err
}

func (s *Storage) DeleteUser(userID string) (userservice.User, error) {
	user := dbUser{}

	id, err := uuid.Parse(userID)
	if err != nil {
		return userservice.User{}, err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return userservice.User{}, err
	}

	row := tx.QueryRow(`
		SELECT id, firstname, lastname
		FROM users
		WHERE id = ?`,
		id,
	)
	err = row.Scan(&user.id, &user.firstname, &user.lastname)
	if err != nil {
		return userservice.User{}, err
	}

	_, err = tx.Exec(`
		DELETE
		FROM users 
		WHERE id = ?;`,
		id,
	)
	if err != nil {
		return userservice.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return userservice.User{}, err
	}

	return userservice.User{
		UserID:    user.id.String(),
		FirstName: user.firstname,
		LastName:  user.lastname,
	}, err
}
