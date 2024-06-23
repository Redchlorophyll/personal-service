package firebase

import (
	"context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func NewFirebase(config FirebaseServiceConfig) FirebaseServiceProvider {
	service := &FirebaseService{
		Env:  config.Env,
		apps: make(map[string]*firebase.App),
	}

	service.initializeFirebase()

	return service
}

func (firebaseService FirebaseService) initializeFirebase() {
	for key, firebaseEnv := range firebaseService.Env.Firebase {
		opt := option.WithCredentialsFile(firebaseEnv.Secretpath)
		app, err := firebase.NewApp(context.Background(), &firebase.Config{
			StorageBucket: firebaseEnv.StorageBucket,
		}, opt)
		if err != nil {
			panic(err)
		}

		firebaseService.apps[key] = app
	}
}
