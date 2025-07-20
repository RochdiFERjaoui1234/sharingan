#!/usr/bin/env python3
"""
Deep API vulnerability scanner.
"""
import sys
import json
import requests

def deep_scan(url):
    report = {'url': url, 'checks': []}
    try:
        r = requests.get(url, timeout=5)
        report['checks'].append({'name': 'status_code', 'value': r.status_code})
        # Additional deep checks can be added here
    except Exception as e:
        report['checks'].append({'error': str(e)})
    print(json.dumps(report, indent=2))
    sys.exit(0)

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: deep_scan_api.py <URL>")
        sys.exit(1)
    deep_scan(sys.argv[1])
