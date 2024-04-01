package util

func GetInsertQuery() string {
	return "INSERT into tbl_audit (source, source_url, resource_url, " +
		"system_file_path, " +
		"strategy, " +
		"invoker, " +
		"channel, " +
		"title, " +
		"timestamp) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetEntriesQuery(sourceUrl string) string {
	return "SELECT source, source_url, resource_url, " +
		"system_file_path, " +
		"strategy, " +
		"invoker, " +
		"channel, " +
		"title, " +
		"timestamp " +
		"FROM tbl_audit where source_url = '" + sourceUrl + "'"
}

func GetLatestEntriesQuery() string {
	return "SELECT source, source_url, resource_url, " +
		"system_file_path, " +
		"strategy, " +
		"invoker, " +
		"channel, " +
		"title, " +
		"timestamp " +
		"FROM tbl_audit LIMIT 10"
}
