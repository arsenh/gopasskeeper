package actions

import "fmt"

func ActionHelp(args Args) {
	fmt.Println(`
Usage:
  >>> [command] [options]

Commands:
  add             Add new credential.
  edit            Edit existing credential.
  delete          Delete credential by service name or username.
  get             Get credential by service name.
  generate        Generate a random secure password.
  quit            Exiting the program

Options:
  --service       Specify the service name (e.g., Gmail, Facebook).
  --username      Specify the username for the service.
  --password      Provide the password for the service (required for 'add').
  --note         Add optional note (e.g., security questions, recovery info).
  --length        Specify the length of the generated password (for 'generate').
  --complexity    Set complexity for password generation: 
                  (e.g., lowercase, uppercase, numbers, symbols).

Examples:
  >>> add  --service=gmail --username=user@example.com --password=securePass123 --note="Dog's name"
  >>> edit --service=gmail --username=newUser@example.com --password=securePass56789 --note="Cat's name"
  >>> delete --service=gmail
  >>> get --service=gmail
  >>> generate --length=16 --complexity={uppercase,numbers,symbols} - by default only lowercase letters, but you can
                                                                      add options such as uppercase letters, numbers, symbols.

  NOTE: --note is optional, but quotes are required to preserve the entire string.
        If the --complexity option is not specified, the password will be generated using lowercase letters only.`)
}
