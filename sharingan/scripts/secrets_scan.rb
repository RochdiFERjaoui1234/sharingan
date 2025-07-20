#!/usr/bin/env ruby
require 'json'

PATTERNS = {
  aws:    /AKIA[0-9A-Z]{16}/,
  slack:  /xox[baprs]-[0-9a-zA-Z]{10,}/,
  github: /ghp_[0-9A-Za-z]{36}/
}

def scan(path)
  findings = []
  Dir.glob("#{path}/**/*").each do |f|
    next unless File.file?(f)
    content = File.read(f)
    PATTERNS.each do |name, regex|
      content.scan(regex).each { |m| findings << { file: f, type: name, secret: m } }
    end
  end
  findings
end

if ARGV.empty?
  puts "Usage: secrets_scan.rb <path>"
  exit 1
end

results = scan(ARGV[0])
puts JSON.pretty_generate(results)
exit(results.empty? ? 0 : 2)