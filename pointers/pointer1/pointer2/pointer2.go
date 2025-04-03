package main

type Database struct {
	user string
}

type Server struct {
	db *Database // uintprt -> 8 bytes long
}

func (s *Server) GetUserFromDB() string {
	// golang is going to deref the db pointer
	// lookup memory address of the pointer
	// if not initialized program will throw the error
	if s.db == nil {
		panic("initialize the pointer")
	}
	return s.db.user
}

func main() {
	s := &Server{
		db: &Database{},
	}
	s.GetUserFromDB()
}
