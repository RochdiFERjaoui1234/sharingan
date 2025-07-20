#!/usr/bin/env ruby
require 'net/http'
require 'uri'
require 'json'

SECURE_HEADERS = [
  'Content-Security-Policy',
  'X-XSS-Protection',
  'Strict-Transport-Security',
  'X-Frame-Options',
  'X-Content-Type-Options'
]

def test_headers(url)
  uri = URI.parse(url)
  resp = Net::HTTP.get_response(uri)
  present = {}
  missing = []
  SECURE_HEADERS.each do |h|
    if resp.key?(h)
      present[h] = resp[h]
    else
      missing << h
    end
  end
  { url: url, status: resp.code.to_i, present: present, missing: missing }
end

if ARGV.empty?
  puts "Usage: header_tester.rb <url>"
  exit 1
end

result = test_headers(ARGV[0])
puts JSON.pretty_generate(result)
exit(result[:missing].empty? ? 0 : 2)