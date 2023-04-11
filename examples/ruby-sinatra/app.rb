require 'sinatra'
require 'sqlite3'

db = SQLite3::Database.new "mydatabase.db"

get '/' do
  @result = db.execute("SELECT * FROM table_name")
  erb :index
end

__END__

@@index
<html>
  <head>
    <title>Sinatra + SQLite Example</title>
  </head>
  <body>
    <% @result.each do |row| %>
      <div><%= row %></div>
    <% end %>
  </body>
</html>