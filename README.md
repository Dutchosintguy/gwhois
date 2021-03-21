# Google Whois
![Twitter Follow](https://img.shields.io/twitter/follow/chaosd0c?style=social)

A go implementation of the Python based [xeuledoc](https://github.com/Malfrats/xeuledoc) project.

## Look up ownership information on any public Google Document.

This simple golang based CLI allows you to look up ownership information for any public Google document.

## Usage

```
Google Whois 0.0.1 by Doctor Chaos

Usage of ./gwhois: <document URL>
  -filename string
        The filename the API response should be saved as (default "api_response.json")
  -save
        Save the API response to disk
```

## Example

```
./gwhois https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit#gid=0

Document ID: 1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms
Document created at: 2011-05-12 18:29:28.159 +0000 UTC
Document last modified at: 2011-05-12 18:29:28.673 +0000 UTC
Document title: "Example Spreadsheet"
Document Owner: A Googler (gdocsteam@gmail.com) (100630705629414352418)
Last modified by: A Googler
Publicly Editable: false
```

## Installation

```
git clone https://github.com/1cbf94bc-bc47-42b9-9197-244437fad1e6/gwhois
cd gwhois
go build
./gwhois
```
