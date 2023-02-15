package core

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"rest-api-simple-auth/configs"
	hlp "rest-api-simple-auth/helpers"
	"time"

	"github.com/joho/godotenv"
)

type App struct {
	Name    string
	Version string
	Debug   bool
	Port    string
	Configs *configs.Config
}

const version = "1.0.0"

func init() {
	fmt.Println("Init func from package app")
}

func (app *App) Setup() error {

	// Get Working Directory
	path, err := os.Getwd()
	hlp.ErrorLogFatalHandler(err)

	// Check existing .env file
	err = app.checkDotEnv(path)
	hlp.ErrorReturnHandler(err)

	// Read & Load .env
	err = godotenv.Load(path + "/.env")
	hlp.ErrorReturnHandler(err)

	/* =======================================
	* Setup App struct Values
	======================================= */
	app.Name = "Simple Lumen with Go"
	app.Version = version
	app.Port = hlp.Env("SERVER_PORT", "PORT")

	// Fill the App Configs with Config struct values
	app.Configs = &configs.Config{
		RootPath: path,
		Database: hlp.Env("DB_DRIVER", "DATABASE"),
	}
	fmt.Println(hlp.Env("DEBUG", "APP_ENV"))
	// fmt.Println("Root Path is :" + app.Configs.RootPath)
	return nil
}

// ListenAndServe to starts the http server
func (app *App) ListenAndServe() {

	srv := &http.Server{
		// Addr: fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Addr: fmt.Sprintf(":%s", app.Port),
		// ErrorLog:     c.ErrorLog,
		// Handler:      c.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	fmt.Printf("Listening on port %s", os.Getenv("SERVER_PORT"))
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func (app *App) checkDotEnv(path string) error {
	envFile := fmt.Sprintf("%s/.env", path)
	var _, envErr = os.Stat(envFile)

	if os.IsNotExist(envErr) {
		err := app.copyDotEnvExample(path)
		hlp.ErrorReturnHandler(err)
	}

	return nil
}

func (app *App) copyDotEnvExample(path string) error {

	exmpl := fmt.Sprintf("%s/.env.example", path)
	var _, exmplErr = os.Stat(exmpl)
	hlp.ErrorLogFatalHandler(exmplErr)

	// Open & Read the .env.example file
	fIn, err := os.Open(exmpl)
	hlp.ErrorLogFatalHandler(err)

	// cara 1 : Close (*os.File) direcly
	defer fIn.Close()

	// Create empty .env file
	fOut, err := os.Create(fmt.Sprintf("%s/.env", path))
	hlp.ErrorLogFatalHandler(err)

	// Cara 2 : Close (*os.File) by clousure / anonimous function
	defer func(fOut *os.File) {
		_ = fOut.Close()
	}(fOut)

	// Copied content of .env.example file to .env
	copied, errCopy := io.Copy(fOut, fIn)
	hlp.ErrorLogFatalHandler(errCopy)

	// fmt.Println(&copied)
	// fmt.Printf("%T", file)
	if copied != 0 {
		fmt.Println("Copied env.example file to .env Success")
	}
	return nil
}
