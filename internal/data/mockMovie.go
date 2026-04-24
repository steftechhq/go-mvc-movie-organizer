package data

import "errors"

type MockMovieModel struct{}

func (m MockMovieModel) Insert(movie *Movie) error {
	// Mock the action...

	return errors.New("123")
}

func (m MockMovieModel) Get(id int64) (*Movie, error) {
	// Mock the action...
	return nil, errors.New("123")

}

func (m MockMovieModel) Update(movie *Movie) error {
	// Mock the action...
	return errors.New("123")

}

func (m MockMovieModel) Delete(id int64) error {
	// Mock the action...
	return errors.New("123")

}
