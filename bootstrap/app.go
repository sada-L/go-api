package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env      *Env
	Database *gorm.DB
}

func App() Application {
	env := NewEnv()
	return Application{
		Env:      env,
		Database: NewDB(env),
	}
}
