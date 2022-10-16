package main

type SocialLogin interface {
	login() string
	register() string
}

type Facebook struct {
	handle string
}

func (f *Facebook) login() string {
	return "Logging In To Facebook"
}

func (f *Facebook) register() string {
	return "Registering On Facebook"
}

//Facebook.login()
//Facebook.register()

type Google struct {
	name string
}

//Google.login()
//Google.register()

type Twitter struct {
	username string
}

//Twitter.login()
//Twitter.register()
