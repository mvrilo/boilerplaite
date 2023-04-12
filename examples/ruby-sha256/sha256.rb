require 'digest'
while input = gets
  puts Digest::SHA256.hexdigest input.chomp
end