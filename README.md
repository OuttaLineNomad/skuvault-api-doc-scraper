# SkuVault API Doc Scraper 

## Overview

Package skuvault-api-doc-scraper scrapes dev.skuvault.com/reference to create a skuvault api package.

After running this, you still have to edit areas where there was not enough data to create clean structs. There will be a lot of []interface{} that you will need to edit. Also, maps may need to be created because the api docs did not specify maps, and only added notes. This tool is great to start with, so you don't have to create the skuvault api package from scratch. 

## Install

```
go get github.com/OuttaLineNomad/skuvault-api-doc-scraper
```
