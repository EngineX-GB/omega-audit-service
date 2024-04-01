package config

func GetDataSourceConfig() string {
	return "root:root@tcp(127.0.0.1:3306)/db_omega?parseTime=true"
}

func GetDriverName() string {
	return "mysql"
}
