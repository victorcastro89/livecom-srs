package firebaseauth

import (
	"context"
	"fmt"
	"log"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseAuth is a wrapper around Firebase Auth
type FirebaseAuth struct {
	Client *auth.Client
}

var (
	instance *FirebaseAuth
	ctx      context.Context
	once     sync.Once
)

func GetInstance() *FirebaseAuth {
    if instance == nil {
        panic("FirebaseAuth must be initialized before getting the instance")
    }
    return instance
}



// Initialize initializes FirebaseAuth using the singleton pattern
func Initialize(serviceAccountKeyFilePath string) (*FirebaseAuth, error) {
	var err error
	once.Do(func() {
		instance, err = initializeFirebase(serviceAccountKeyFilePath)
	})
	return instance, err
}

func initializeFirebase(serviceAccountKeyFilePath string) (*FirebaseAuth, error) {
	ctx = context.Background()
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
		return nil, err
	}

	return &FirebaseAuth{Client: client}, nil
}


// VerifyIDToken verifies the ID token
func (f *FirebaseAuth) VerifyIDToken(idToken string) (*auth.Token, error) {
	token, err := f.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		fmt.Errorf("error verifying ID token: %v\n", err)
		return nil, err
	}

	return token, nil
}