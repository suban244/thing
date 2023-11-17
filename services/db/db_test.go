package db

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func setup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Could not load env variables")
	}
}

func shutdown() {}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestService(t *testing.T) {
	dbService := service{}
	err := dbService.cleanDatabase()
	if err != nil {
		t.Fatal(err)
	}

	users := [5]string{
		"subant",
		"subant2",
		"subant3",
		"subant4",
		"subant5",
	}

	for i := range users {
		err = dbService.createUser(users[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := range users {
		exists, err := dbService.checkIfUserExists(users[i])
		if err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Fatal("User not in databse")
		}
	}

	toDeleteIndex := 4
	err = dbService.deleteUser(users[toDeleteIndex])
	if err != nil {
		t.Fatal(err)
	}
	exists, err := dbService.checkIfUserExists(users[toDeleteIndex])
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("User is supposed to be deleted")
	}

	assignments := []Assignment{
		{Id: "2023-PY100-1", gradingfile: "Calc0.py", createdBy: users[0]},
		{Id: "2023-PY100-2", gradingfile: "Calc1.py", createdBy: users[0]},
		{Id: "2023-PY100-3", gradingfile: "Calc2.py", createdBy: users[0]},
		{Id: "2023-PY101-1", gradingfile: "Calc2.py", createdBy: users[1]},
		{Id: "2023-PY101-2", gradingfile: "Calc2.py", createdBy: users[1]},
	}

	for _, v := range assignments {
		err := dbService.createAssignment(v.Id, v.gradingfile, v.createdBy)
		if err != nil {
			t.Fatal(err)
		}
	}

	err = dbService.updateScore(assignments[0].Id, 32)
	if err != nil {
	}
	dbService.updateScore(assignments[0].Id, 32)
	dbService.updateScore(assignments[0].Id, 32)
}

// func TestServiceSub(t *testing.T) {
// 	t.Run("init", func(t *testing.T) {
// 		fmt.Println("thing")
// 	})
//
// }
