package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

//errors
var errNotFound = errors.New("Not found")
var errWordExists = errors.New("That word already exists")
var errCantUpdate = errors.New("Can not update non-existing word")

//methods // retrun string and error 
//Search for word
func (d Dictionary) Search(word string) (string, error){
	//map have two value// exist is boolean
	value, exists := d[word]	
	if exists{
		//nil (error is nil no error)//we have to return two values
		return value, nil
	}
	//return two values, first one is nothing cause no value
	return "",errNotFound
}

//add a word to the dictionary
func (d Dictionary) Add(word, def string) error{
	//if there is word already that would throw error
	//we don't need definition
	_, err := d.Search(word)
	if err == errNotFound{
		d[word] = def
	}else if err == nil {
		return errWordExists
	}
	return nil
}

func(d Dictionary) Update(word, def string)error{
	//search the word exist first
	_, err := d.Search(word)
	switch err{
	//if there is no word
	case errNotFound:
		return errCantUpdate
	//word exist so we can update word
	case nil:
		d[word] = def
	}
	return nil
}

func(d Dictionary) Delete (word string){
//delete function is built in Go
	delete(d, word)
}
