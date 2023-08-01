package database

var connection string

// init akan selalu dipanggil ketika package di import
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
