package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongodbHost      = "localhost" // Assuming you're hosting the mongodb instance locally
	databaseName     = "todoDB"
	todoCollection   = "todos"
	notObjectIdError = "Error: not ObjectId"
)

var (
	mgoSession *mgo.Session
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongodbHost)
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

// The withCollection() function takes the name of the collection, along with a function that
// expects the connection object to that collection, and can execute functions on it.
func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}

func SearchTodos(q interface{}, skip int, limit int) (searchResults []Todo, searchErr error) {
	query := func(c *mgo.Collection) error {
		var fn error
		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(&searchResults)
		} else {
			fn = c.Find(q).Skip(skip).Limit(limit).All(&searchResults)
		}
		return fn
	}
	search := func() error {
		return withCollection(todoCollection, query)
	}
	err := search()
	return searchResults, err
}

func RepoFindTodoWithId(id string) (searchResults []Todo, err error) {
	if !bson.IsObjectIdHex(id) {
		return searchResults, fmt.Errorf("Could not find Todo with id of %s", id) // Handles cases of incorrect ids
	}
	searchResults, err = SearchTodos(bson.M{"_id": bson.ObjectIdHex(id)}, 0, -1)
	return searchResults, err
}

func RepoGetAllTodos() (searchResults []Todo, err error) {
	searchResults, err = SearchTodos(nil, 0, -1)
	return searchResults, err
}

func RepoCreateTodo(t Todo) error {
	insert := func() error {
		return withCollection(todoCollection, func(c *mgo.Collection) error {
			err := c.Insert(t)
			return err
		})
	}
	err := insert()
	return err
}

func RepoDestoryTodo(id string) error {
	var err error

	if !bson.IsObjectIdHex(id) {
		err = fmt.Errorf("Could not find Todo with id of %s to delete", id)
		return err
	}
	remove := func(c *mgo.Collection) error {
		fn := c.RemoveId(bson.ObjectIdHex(id))
		return fn
	}
	removeTodo := func() error {
		return withCollection(todoCollection, remove)
	}
	err = removeTodo()
	return err
}
