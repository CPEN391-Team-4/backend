package environment

import "os"

// Holds environment variables used by server
type Env struct {
	DbUri               string
	Db                  string
	Imagestore          string
	Videostore          string
	ServerAddress       string
	FaceSubscriptionKey string
	FaceEndpoint        string
	FirebaseKeyfile     string
}

// ReadEnv Reads all environment variables used by server into an Env
func (e *Env) ReadEnv() {
	e.DbUri = os.Getenv("DB_URI")
	e.Db = os.Getenv("DB")
	e.Imagestore = os.Getenv("IMAGESTORE")
	e.Videostore = os.Getenv("VIDEOSTORE")
	e.ServerAddress = os.Getenv("SERVER_ADDRESS")
	e.FaceSubscriptionKey = os.Getenv("FACE_SUBSCRIPTION_KEY")
	e.FaceEndpoint = os.Getenv("FACE_ENDPOINT")
	e.FirebaseKeyfile = os.Getenv("FIREBASE_KEYFILE")
}