Gem::Specification.new do |s|
  s.name        = 'myapp'
  s.version     = '0.0.1'
  s.authors     = ['Your Name']
  s.email       = 'you@example.com'
  s.summary     = 'A Sinatra app with SQLite integration'
  s.description = 'A simple Sinatra app that fetches data from a SQLite database'
  s.homepage    = 'https://example.com'
  s.license     = 'MIT'

  s.files       = Dir['lib/**/*', 'views/**/*', 'public/**/*']
  s.require_path = 'lib'
  s.bindir      = 'bin'
  s.executables = ['myapp']
  s.add_development_dependency 'bundler', '~> 2.0'
end
