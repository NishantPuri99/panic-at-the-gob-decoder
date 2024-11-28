# Panic! at the Gob Decoder

- Namespace: picoctf/18739f24
- ID: panic-at-the-gob-decoder
- Type: custom
- Category: Web Exploitation
- Points: 1
- Templatable: no

## Hints

- You might need this https://goplay.tools/snippet/8IXAiVn0sQj
- If the route doesn't work, then you know I don't love it as much

## Description

*The CVE POC from Walmart*
Go is a pretty safe language. What could *GO* wrong?

Download the server file: {{url_for("main.go")}}

## Learning Objective

- Panic's in GO
- The potential impact of CVE-2024-34156

## Solution Overview

The user submits a gob-encode-nested-string and inputs it in the text area which causes the code to panic.
This leads to the api_key and route for the flag being revealed.

## Details

Connect to the webserver {{link_as("port", "/", "Here")}}

## Attributes
- author: nishant
