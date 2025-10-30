package zerolog

type FnOption func(option *Option)

type Option struct {
	Level     string // DEBUG/INFO/WARN/ERROR
	Formatter string // TEXT/JSON

	Console struct {
		Enabled    bool
		ApplyColor bool
	}
	File struct {
		Enabled  bool
		Path     string
		Name     string
		MaxSize  int
		Compress bool
		MaxAge   int
	}
}

func defaultOption() *Option {
	return &Option{}
}

func WithLevel(level string) FnOption {
	return func(option *Option) {
		option.Level = level
	}
}

func WithFormatter(formatter string) FnOption {
	return func(option *Option) {
		option.Formatter = formatter
	}
}

func WithConsole(enabled bool) FnOption {
	return func(option *Option) {
		option.Console.Enabled = enabled
	}
}

func WithConsoleApplyColor(applyColor bool) FnOption {
	return func(option *Option) {
		option.Console.ApplyColor = applyColor
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
