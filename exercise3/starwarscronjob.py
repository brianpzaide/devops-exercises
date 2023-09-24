#!/usr/bin/env python

import datetime
import argparse
import os
from typing import Dict
import yaml
import json
import urllib.request
from urllib.error import HTTPError

input_file = ""
out_dir = ""
base_url = "https://swapi.dev/api"
map_type = {'people': 'people', 'starship': 'starships', 'planet':'planets'}

def fetchAndStoreInfo(requested_item: Dict, dname):
    data = ""
    req_type = map_type[requested_item['type']]

    fetch_url = f"{base_url}/{req_type}/{requested_item['id']}"
    try:
        with urllib.request.urlopen(fetch_url) as response:
            data = response.read()
            data = json.loads(data)
        fname = os.path.join(dname, f"{data['name']}.json")
        data_dict = {}
        for req_attr in requested_item['infoRequest']:
            data_dict[req_attr] = data[req_attr]
        with open(fname, 'w') as f:
            json.dump(data_dict, f)
    except HTTPError as error:
         print(error.status, error.reason)

def main():
    dname = os.path.join(out_dir, str(int(datetime.datetime.now().timestamp())))
    with open(input_file, "r") as stream:
        try:
            request_data = yaml.safe_load(stream)['input']
        except yaml.YAMLError as exc:
            print(exc)
            return
    try:
    	os.makedirs(dname)
    except OSError as err:
        print(err)
        return
    for item in request_data:
        fetchAndStoreInfo(item, dname)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-i", "--inputfile", type=str, required=True)
    parser.add_argument("-o", "--outdir", type=str, required=True)
    args = parser.parse_args()
    input_file = args.inputfile
    out_dir = args.outdir
    main()
