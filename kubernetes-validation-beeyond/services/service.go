package services

// Inits all services (calls their init<Service>() function)
// Initializes the database connection (initDatabase())
func Init() {
	initDatabase()
}
