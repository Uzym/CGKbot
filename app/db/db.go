package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DSN = ""

func SetNewUser(chat_id int64, username string) error {
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("select exists(select 1 from users where chat_id=%d)", chat_id)
	row := db.QueryRow(query)
	var exists bool
	err = row.Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		query = fmt.Sprintf("INSERT INTO users (chat_id, username) VALUES (%d, '%s');", chat_id, username)
		_, err = db.Exec(query)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

	return nil
}

func SetAnswer(user_id int64, question_id int64) error {
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("Insert INTO userstoquestions (user_id, question_id) VALUES(%d, %d);", user_id, question_id)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

type Leader struct {
	username string
	cnt      int64
}

func (l *Leader) GetUsername() string {
	return l.username
}

func (l *Leader) GetCnt() int64 {
	return l.cnt
}

func GetLeaderboard() ([]Leader, error) {
	var leaders []Leader

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return leaders, err
	}
	defer db.Close()

	query := "SELECT users.username, usercnt.cnt FROM users JOIN (SELECT user_id, COUNT(*) as cnt FROM userstoquestions GROUP BY user_id) as usercnt ON usercnt.user_id = users.chat_id ORDER BY cnt DESC;"

	rows, err := db.Query(query)
	if err != nil {
		return leaders, err
	}
	defer rows.Close()

	for rows.Next() {
		var lead Leader
		if err := rows.Scan(&lead.username, &lead.cnt); err != nil {
			return leaders, err
		}
		leaders = append(leaders, lead)
	}

	return leaders, nil
}

func GetQuestion(chat_id int64) (question string, question_id int64, answer string, err error) {

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return "", 0, "", err
	}
	defer db.Close()

	query := fmt.Sprintf("select question, id, answer from questions where not(id in (select question_id from userstoquestions where user_id = %d)) order by random() limit 1", chat_id)

	row := db.QueryRow(query)

	if row == nil {
		return "", 0, "", nil
	}

	err = row.Scan(&question, &question_id, &answer)
	if err != nil {
		return "", 0, "", err
	}

	return question, question_id, answer, nil
}

func GetNumberOfUsers() (int64, error) {

	var count int64

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(*) FROM users;")

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
