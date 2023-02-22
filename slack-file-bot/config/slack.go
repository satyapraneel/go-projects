package config

type SlackConfig struct {
	Bot       string `mapstructure:"bot"`
	Token     string `mapstructure:"token"`
	ChannelId string `mapstructure:"channel_id"`
}
