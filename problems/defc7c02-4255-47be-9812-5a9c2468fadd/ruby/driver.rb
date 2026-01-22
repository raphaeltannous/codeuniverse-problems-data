require 'json'

def load_testcases(filename)
  data = File.read(filename)
  JSON.parse(data, symbolize_names: true)
rescue => e
  $stderr.puts "Failed to open testcases file."
  exit(1)
end
