package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Table represents the structure of a table in the JSON file
type Table struct {
	TableName string `json:"table_name"`
	Structure []struct {
		ColumnName string `json:"column_name"`
		Type       string `json:"type"`
		Nullable   string `json:"nullable"`
		Index      string `json:"index"`
		Unique     string `json:"unique"`
		Extra      string `json:"extra"`
	} `json:"structure"`
}
func CheckAndCreateAllTables() {
	tables, err := readTablesFromJSON("DB/tables.json")
	if err != nil {
		log.Println("Failed to read tables configuration")
		return
	}

	// List of expected table names
	expectedTableNames := make(map[string]struct{})
	for _, table := range tables {
		expectedTableNames[table.TableName] = struct{}{}
	}

	// Check and create tables
	for _, table := range tables {
		if err := createTableIfNotExists(table); err != nil {
			log.Println("Failed to read tables configuration", err.Error())
			return
		}
	}
}
// CheckAndCreateTables checks and creates necessary tables in the database
func CheckAndCreateTable(tableName string) error{
	tables, err := readTablesFromJSON("db/mysqlTables.json")
	if err != nil {
		log.Println("Failed to read tables configuration", tableName)
		return err
	}

	// List of expected table names
	expectedTableNames := make(map[string]struct{})
	for _, table := range tables {
		expectedTableNames[table.TableName] = struct{}{}
	}

	// Check for any tables that are not in the JSON file
	actualTables := getActualTableNames(tableName)
	for _, actualTable := range actualTables {
		if _, exists := expectedTableNames[actualTable]; !exists {
			log.Printf("Table %s is not defined in the JSON configuration\n", actualTable)
			
			return err
		}
	}
	return err
}

// readTablesFromJSON reads the table configuration from a JSON file
func readTablesFromJSON(filePath string) ([]Table, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config struct {
		Tables []Table `json:"tables"`
	}
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return config.Tables, nil
}

// createTableIfNotExists creates a table if it doesn't exist based on the provided Table struct
func createTableIfNotExists(table Table) error {
	query := "CREATE TABLE IF NOT EXISTS " + table.TableName + " ("

	for i, column := range table.Structure {
		query += column.ColumnName + " " + column.Type
		if column.Nullable == "NO" {
			query += " NOT NULL"
		}
		if column.Index != "" {
			query += " " + column.Index
		}
		if column.Unique == "YES" {
			query += " UNIQUE"
		}
		if column.Extra != "" {
			query += " " + column.Extra
		}
		if i < len(table.Structure)-1 {
			query += ", "
		}
	}

	query += ")"

	_, err := MySQLDB.Exec(query)
	if err != nil {
		log.Println("Error creating table:", err)
		return err
	}
	return nil
}

// getActualTableNames retrieves the names of the actual tables in the database
func getActualTableNames(tableName string) []string {
	// This function should implement logic to retrieve the actual table names from the database
	// For example, you can query the information_schema.tables
	// Here, we return a placeholder for demonstration purposes
	return []string{tableName} // Replace with actual implementation
}

// getActualTableStructure retrieves the structure of a specified table from the database
func getActualTableStructure(tableName string) ([]struct {
	ColumnName string
	Type       string
	Nullable   string
	Index      string
	Unique     string
	Extra      string
}, error) {
	query := `
		SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_KEY, 
		       CASE WHEN COLUMN_KEY = 'PRI' THEN 'YES' ELSE 'NO' END AS UNIQUE,
		       EXTRA 
		FROM information_schema.COLUMNS 
		WHERE TABLE_NAME = ? AND TABLE_SCHEMA = DATABASE()`

	rows, err := MySQLDB.Query(query, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []struct {
		ColumnName string
		Type       string
		Nullable   string
		Index      string
		Unique     string
		Extra      string
	}

	for rows.Next() {
		var column struct {
			ColumnName string
			Type       string
			Nullable   string
			Index      string
			Unique     string
			Extra      string
		}
		if err := rows.Scan(&column.ColumnName, &column.Type, &column.Nullable, &column.Index, &column.Unique, &column.Extra); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}