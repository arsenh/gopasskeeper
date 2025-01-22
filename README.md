# gopasskeeper
Offline Password Manager

### **Functional Specification for an Offline Password Manager**

#### **Objective**
Create an offline password manager that securely stores, retrieves, and manages user credentials without relying on internet access. The application will ensure data privacy through encryption and provide a user-friendly interface via the command line.

---

### **Core Features**

#### **1. User Authentication**
- **Purpose**: Protect access to the stored credentials.
- **Functionality**:
  - Set up a master password during the first use.
  - Require the master password for all subsequent access.
  - Lock the application after a period of inactivity.

---

#### **2. Credential Management**
- **Purpose**: Enable users to securely store and manage credentials.
- **Functionality**:
  - Add new credentials with the following details:
    - Service name.
    - Username.
    - Password.
    - Notes (optional).
  - Edit existing credentials.
  - Delete credentials by service name or username.
  - Retrieve credentials by service name or username.
  - Search credentials using partial matches (e.g., "email" matches "Gmail").
  - Option to generate random secure passwords.

---

#### **3. Data Encryption**
- **Purpose**: Ensure data security and prevent unauthorized access.
- **Functionality**:
  - Encrypt all stored data using AES-256.
  - Hash and salt the master password using a secure algorithm (e.g., bcrypt).
  - Encrypt and decrypt data in memory during runtime.

---

#### **4. Data Storage**
- **Purpose**: Store credentials persistently on the local machine.
- **Functionality**:
  - Save encrypted credentials in a single local file (e.g., `passwords.db`).
  - Ensure the storage file is portable and can be backed up manually.
  - Validate file integrity before decryption to prevent corruption issues.

---

#### **5. User Interface**
- **Purpose**: Provide a simple and intuitive way for users to interact with the application.
- **Functionality**:
  - Command-line interface (CLI) with clear prompts.
  - Menu-based navigation for common tasks:
    - Add credentials.
    - Retrieve credentials.
    - Edit or delete credentials.
    - View a list of stored services.
  - Display help information for all commands.

---

#### **6. Security Features**
- **Purpose**: Protect user data and application access.
- **Functionality**:
  - Automatically lock the application after a configurable period of inactivity.
  - Limit the number of failed master password attempts before locking the user out.
  - Require the master password to decrypt credentials during retrieval.
  - Prevent password recovery (for true offline security).
  - Provide an option to delete all stored credentials if the user forgets the master password.

---

### **Extended Features (Optional)**
- **1. Password Strength Analyzer**
  - Evaluate the strength of passwords entered by the user.
  - Suggest improvements for weak passwords.
- **2. Multi-Profile Support**
  - Allow users to maintain multiple independent credential stores.
- **3. Import/Export**
  - Import credentials from CSV files.
  - Export encrypted backups for safekeeping.
- **4. Clipboard Integration**
  - Copy retrieved passwords to the clipboard temporarily for easy use.
  - Automatically clear the clipboard after a configurable timeout.

---

### **Technical Specification**

#### **1. Programming Language**
- **Choice**: Golang
  - Lightweight and efficient.
  - Provides robust libraries for encryption and file handling.

---

#### **2. Data Format**
- **Storage File**: Encrypted JSON or SQLite database.
- **Encryption**: AES-256 for data encryption and bcrypt for password hashing.

---

#### **3. CLI Design**
- Use a library like [Cobra](https://github.com/spf13/cobra) or [urfave/cli](https://github.com/urfave/cli) for building the CLI.
- Example commands:
  - `passmgr add` – Add a new credential.
  - `passmgr get <service_name>` – Retrieve credentials for a service.
  - `passmgr list` – List all stored services.
  - `passmgr delete <service_name>` – Delete credentials for a service.

---

### **User Scenarios**

#### **Scenario 1: First-Time Setup**
1. User launches the password manager.
2. The system prompts them to create a master password.
3. The master password is hashed and stored locally.

#### **Scenario 2: Adding Credentials**
1. User logs in with their master password.
2. They execute the `add` command to store a new credential.
3. The system encrypts the credential and saves it.

#### **Scenario 3: Retrieving Credentials**
1. User logs in with their master password.
2. They execute the `get` command with the service name.
3. The system decrypts and displays the credential securely.

#### **Scenario 4: Editing Credentials**
1. User logs in and lists stored services using `list`.
2. They edit a credential using the `edit` command.
3. The updated credential is re-encrypted and saved.

---

### **Security Considerations**
1. **No Plaintext Passwords**:
   - Credentials should never be stored or displayed in plaintext.
2. **In-Memory Encryption**:
   - Decrypt data only when needed and clear sensitive data from memory immediately after use.
3. **Master Password Recovery**:
   - Do not provide recovery options for the master password to maintain offline security.
4. **Brute Force Protection**:
   - Introduce a delay after consecutive failed login attempts.

---

### **Deliverables**
1. A CLI-based password manager executable.
2. Documentation with:
   - Installation and usage instructions.
   - Command examples and troubleshooting tips.
3. Sample encrypted data file for demonstration.

This design ensures an offline password manager that is secure, efficient, and user-friendly while addressing common security needs. Let me know if you’d like help with implementation!