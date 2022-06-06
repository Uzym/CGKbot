package bot

import (
	"2bot/app/db"
	"errors"
	"fmt"
	"time"
)

func Leaderboard() ([]db.Leader, error) {
	leaders, err := db.GetLeaderboard()
	if err != nil {
		return nil, err
	}
	return leaders, err
}

type ans struct {
	answer      string
	question_id int64
	time        time.Time
}

var userQuestion map[int64]ans

var QuestionTime time.Duration

func garbageCleaner() {
	for {
		time.Sleep(QuestionTime * 5)
		for key, value := range userQuestion {
			if !time.Now().Before(value.time.Add(QuestionTime)) {
				delete(userQuestion, key)
			}
		}
	}
}

func SetNewUser(id int64, username string) error {
	err := db.SetNewUser(id, username)
	return err
}

func Question(id int64) (string, error) {

	question, question_id, answer, err := db.GetQuestion(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", nil
		} else {
			return "", err
		}
	}

	userQuestion[id] = ans{answer, question_id, time.Now()}

	return question, nil
}

func Answer(id int64, userAns string, responseTime time.Time) (bool, error) {

	val, ok := userQuestion[id]

	if !ok {
		return false, errors.New("the question was not found")
	}

	delete(userQuestion, id)

	if !responseTime.Before(val.time.Add(QuestionTime)) {
		return false, nil
	}

	fmt.Println(id, userAns, val.answer)

	if userAns != val.answer {
		return false, nil
	}

	err := db.SetAnswer(id, val.question_id)
	if err != nil {
		return false, err
	}

	return true, nil

}

func StartBot() {
	userQuestion = make(map[int64]ans)
	garbageCleaner()
}
