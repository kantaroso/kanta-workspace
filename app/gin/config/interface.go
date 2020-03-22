package config

// config系はベストプラクティスがなさそうなので生のgoファイルでやる

// Config 各種config取得
type Config interface {
	GetDatabase() Database
}
