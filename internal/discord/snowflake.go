package discord

import (
	"strconv"

	tempest "github.com/Amatsagu/Tempest"
)

type Snowflake struct {
	tempest.Snowflake
}

func (s *Snowflake) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}

	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}

	s.Snowflake = tempest.Snowflake(i)
	return nil
}
