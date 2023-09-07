package echoswagger

type Config struct {
	FILE_PATH string `default:"swagger.json"`
	URL_PATH  string `default:"docs"`
}
