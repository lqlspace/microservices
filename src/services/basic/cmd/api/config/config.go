package config

type (

	Config struct {
		Host string
		Port int64
		Mysql struct{
			DataSource string
		}
	}
)
