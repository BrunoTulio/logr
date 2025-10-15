package zap

type FnOption func(option *Option)

type Option struct {
	Console struct {
		Enabled   bool
		Level     string
		Formatter string
	}
	File struct {
		Formatter string
		Enabled   bool
		Path      string
		Name      string
		MaxSize   int
		Compress  bool
		MaxAge    int
		Level     string
	}
}

func defaultOption() *Option {
	return &Option{}
}

func WithConsoleLevel(level string) FnOption {
	return func(option *Option) {
		option.Console.Level = level
	}
}

func WithFileLevel(level string) FnOption {
	return func(option *Option) {
		option.File.Level = level
	}
}

func WithConsoleFormatter(formatter string) FnOption {
	return func(option *Option) {
		option.Console.Formatter = formatter
	}
}

func WithFileFormatter(formatter string) FnOption {
	return func(option *Option) {
		option.File.Formatter = formatter
	}
}

func WithConsole(enabled bool) FnOption {
	return func(option *Option) {
		option.Console.Enabled = enabled
	}
}

func WithFile(enabled bool, path, name string) FnOption {
	return func(option *Option) {
		option.File.Enabled = enabled
		option.File.Path = path
		option.File.Name = name
	}
}

func WithFileRotation(maxSize int, maxAge int, compress bool) FnOption {
	return func(option *Option) {
		option.File.MaxSize = maxSize
		option.File.MaxAge = maxAge
		option.File.Compress = compress
	}
}
