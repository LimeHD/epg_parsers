set :stage, :production
set :disallow_pushing, true
server ENV['PRODUCTION_HOST'], user: fetch(:user), port: '22', roles: %w(app).freeze
set :keep_releases, 10
