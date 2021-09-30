package dialects

const Columns = `
	SELECT
		C.TABLE_NAME as table_name,
		C.COLUMN_NAME as column_name,
		C.COLUMN_DEFAULT as default_value,
		C.IS_NULLABLE as is_nullable,
		C.COLUMN_TYPE as data_type,
		C.COLUMN_KEY as column_key,
		C.CHARACTER_MAXIMUM_LENGTH as max_length,
		C.EXTRA as extra
	FROM
		INFORMATION_SCHEMA.COLUMNS AS C
	LEFT JOIN
		INFORMATION_SCHEMA.TABLES AS T ON C.TABLE_NAME = T.TABLE_NAME
		AND C.TABLE_SCHEMA = T.TABLE_SCHEMA
	WHERE
		T.TABLE_TYPE = 'BASE TABLE' AND
		C.TABLE_SCHEMA = ?;
`
