package actions

import "fmt"

func ActionHelp(args Args) {
	fmt.Println(`
Usage:
  >>> [command] [options]

Commands:
  add             Add new credentials.
  edit            Edit existing credentials.
  delete          Delete credentials by service name or username.
  retrieve        Retrieve credentials by service name or username.
  search          Search credentials using partial matches.
  generate        Generate a random secure password.

Options:
  --service       Specify the service name (e.g., Gmail, Facebook).
  --username      Specify the username for the service.
  --password      Provide the password for the service (required for 'add').
  --notes         Add optional notes (e.g., security questions, recovery info).
  --length        Specify the length of the generated password (for 'generate').
  --complexity    Set complexity for password generation: 
                  (e.g., lowercase, uppercase, numbers, symbols).

Examples:
  >>> add --service "Gmail" --username "user@example.com" --password "securePass123"
  >>> edit --service "Gmail" --username "newUser@example.com"
  >>> delete --service "Gmail"
  >>> retrieve --service "Gmail"
  >>> search --service "mail"
  >>> generate --length 16 --complexity "uppercase,numbers,symbols"

  `)
}
