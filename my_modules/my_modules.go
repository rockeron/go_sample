package my_modules

import ("fmt")

func Call(name string) string{
	// Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}