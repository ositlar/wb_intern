package model

type DatabaseData struct {
	Id   string
	Info map[string]interface{}
}

// Задел на верификацию
func (o *DatabaseData) Validate() bool {
	keys := make([]string, 0, len(o.Info))
	for k := range o.Info {
		keys = append(keys, k)
	}
	for i := range keys {
		if o.Info[keys[i]] == nil {
			return false
		}
	}
	if o.Info["delivery"] == nil || o.Info["delivery"] == "" {
		return false
	}
	return true
}
