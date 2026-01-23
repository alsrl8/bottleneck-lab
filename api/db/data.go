package db

type SensorLog struct {
	Id        int     `json:"id"`
	SensorId  string  `json:"sensor_id"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"created_at"`
}
