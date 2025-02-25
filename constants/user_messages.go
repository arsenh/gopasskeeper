package constants

import . "gopasskeeper/colors"

// user messages
const (
	UserPromptMsg   = "⚡ "
	LoggerPrefixMsg = "🔥 gopasskeeper: "

	AppBannerMsg = Bold + Blue + "***********************************************************" + Reset + "\n" +
		Bold + Green + "🚀  Welcome to GoPassKeeper - Offline Password Manager  🔒" + Reset + "\n" +
		Bold + Blue + "***********************************************************" + Reset + "\n" +
		Yellow + "🔑 Securely store and manage your passwords offline!" + Reset

	StrongPasswordMsg = `
	❌ Your password is not strong enough!

	To ensure your account remains secure, your password must meet the following requirements:

	✅ At least 8 characters long
	✅ Includes at least one lowercase letter (a-z)
	✅ Includes at least one uppercase letter (A-Z)
	✅ Includes at least one digit (0-9)
	✅ Includes at least one special character (!@#$%^&*)

	🔒 A strong password helps protect your data from unauthorized access. 
	Please create a stronger password that meets all security requirements.
	`
	FirstTimeMasterPasswordSetupMsg = "The password manager requires a master password for future access to your credentials."
	EnterMasterPasswordMsg          = "Enter Master Password: "
	ConfirmMasterPasswordMsg        = "Confirm Master Password: "
	UseExistCommandInsteadMsg       = "\nuse '%s' to exit instead.\n"

	PasswordMatchMessageMsg = "✅ Passwords match. Master key is stored in %s password file.\n" +
		"Password file is used to store credentials and changing it may corrupt the data.\n" +
		Red + "Do not edit this file. Otherwise, the data will be lost." + Reset

	InvalidPasswordMsg = "Invalid password. Please try again."
	PasswordCorrectMsg = "✅ Passwords correct. Master key loaded from %s\n"

	PasswordFileEmptyCorruptedMsg = Red + "The password file in %s is empty or corrupted." + Reset +
		"There is no way to download the credentials. To set up a new password file, delete it and restart the program."

	ApplicationDirNotExistCreateNew = "application directory not exist, created new one in %s\n"

	PasswordFileCreatedMsg = "password file is created %s\n"

	HelpMsg = `
Usage:
>>> [command] [options]

📌 Commands:
  🆕  add         ➜ Add new credential.
  📝  edit        ➜ Edit existing credential.
  ❌  delete      ➜ Delete credential by service name or username.
  🔍  get         ➜ Get credential by service name.
  🔢  generate    ➜ Generate a random secure password.
  🚪  quit        ➜ Exit the program.

⚙️ Options:
  🏷️  --service     ➜ Specify the service name (e.g., Gmail, Facebook).
  👤  --username    ➜ Specify the username for the service.
  🔑  --password    ➜ Provide the password for the service (**required for 'add'**).
  📝  --note        ➜ Add optional note (e.g., security questions, recovery info).
  📏  --length      ➜ Specify the length of the generated password (for 'generate').
  🔣  --complexity  ➜ Set complexity for password generation:
                        (e.g., **lowercase, uppercase, numbers, symbols**).

📌 Examples:
  👉  🔹 add --service=gmail --username=user@example.com --password=securePass123 --note="Dog's name"
  👉  🔹 edit --service=gmail --username=newUser@example.com --password=securePass56789 --note="Cat's name"
  👉  🔹 delete --service=gmail
  👉  🔹 get --service=gmail
  👉  🔹 generate --length=16 --complexity={uppercase,numbers,symbols}
      
      - By default, only lowercase letters are used.
      - You can **add options** such as **uppercase letters, numbers, and symbols**.

⚠️ Note:
  --note is optional, but **quotes are required** to preserve the entire string.
  If --complexity is not specified, the password will be generated using lowercase letters only.
	`
)
