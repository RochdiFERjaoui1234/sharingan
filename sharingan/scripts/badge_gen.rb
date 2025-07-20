#!/usr/bin/env ruby
require 'erb'

status = ARGV[0] || 'unknown'
color = case status.downcase
when 'passed' then 'brightgreen'
when 'failed' then 'red'
else 'lightgrey'
end

template = ERB.new <<~SVG
<svg xmlns="http://www.w3.org/2000/svg" width="120" height="20">
  <rect width="80" height="20" fill="#555"/>
  <rect x="80" width="40" height="20" fill="<%= color %>"/>
  <text x="40" y="14" fill="#fff" text-anchor="middle" font-family="Verdana" font-size="11">Security</text>
  <text x="100" y="14" fill="#fff" text-anchor="middle" font-family="Verdana" font-size="11"><%= status.upcase %></text>
</svg>
SVG

puts template.result(binding)