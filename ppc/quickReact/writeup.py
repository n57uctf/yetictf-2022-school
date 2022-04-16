#!/usr/bin/env python3
import socket, re, time, hashlib, sys

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(("localhost", 12345))
sock.settimeout(2)

flagRegexp = re.compile(r"YetiCTF")
plusRegexp = re.compile(r"([\d]+)\+([\d]+)=")
minusRegexp = re.compile(r"([\d]+)\-([\d]+)=")
multiplyRegexp = re.compile(r"([\d]+)\*([\d]+)=")
divideRegexp = re.compile(r"([\d]+)\/([\d]+)=")
md5Regexp = re.compile(r"md5\(([\d]+)\)=")
sha256Regexp = re.compile(r"sha256\(([\d]+)\)=")
sha1Regexp = re.compile(r"sha1\(([\d]+)\)=")
sha512Regexp = re.compile(r"sha512\(([\d]+)\)=")

#sock.recv(4096)

while True:
    try:
        data = sock.recv(4096)
    except TimeoutError as e:
        err = e.args[0]
        if err == 'timed out':
            sleep(1)
            print('recv timed out, retry later')
            continue
        else:
            print(e)
            sys.exit(1)
    except Exception as e:
        # Something else happened, handle error, exit, etc.
        print(e)
        sys.exit(1)
    else:
        if len(data) == 0:
            print('orderly shutdown on server end')
            sys.exit(0)
        else:
            data = str(data.decode())
            flagMatch = flagRegexp.match(data)
            plusMatch = plusRegexp.match(data)
            minusMatch =  minusRegexp.match(data)
            multiplyMatch = multiplyRegexp.match(data)
            divideMatch = divideRegexp.match(data)
            md5Match = md5Regexp.match(data)
            sha256Match = sha256Regexp.match(data)
            sha1Match = sha1Regexp.match(data)
            sha512Match = sha512Regexp.match(data)
            if flagMatch:
                print(data)
                break
            elif plusMatch:
                print(data,str(int(pluMatch[1]) + int(plusMatch[2])))
                sock.send(str.encode(str(int(plusMatch[1]) + int(plusMatch[2])) + "\n"))
            elif minusMatch:
                print(data,str(int(minusMatch[1]) - int(minusMatch[2])))
                sock.send(str.encode(str(int(minusMatch[1]) - int(minusMatch[2])) + "\n"))
            elif multiplyMatch:
                print(data,str(int(multiplyMatch[1]) * int(multiplyMatch[2])))
                sock.send(str.encode(str(int(multiplyMatch[1]) * int(multiplyMatch[2])) + "\n"))
            elif divideMatch:
                print(data,str(int(int(divideMatch[1]) / int(divideMatch[2]))))
                sock.send(str.encode(str(int(int(divideMatch[1]) / int(divideMatch[2]))) + "\n"))
            elif md5Match:
                result = hashlib.md5(md5Match[1].encode()).hexdigest()
                print(data,result)
                sock.send(str.encode(result + "\n"))
            elif sha256Match:
                result = hashlib.sha256(sha256Match[1].encode()).hexdigest()
                print(data,result)
                sock.send(str.encode(result + "\n"))
            elif sha1Match:
                result = hashlib.sha1(sha1Match[1].encode()).hexdigest()
                print(data,result)
                sock.send(str.encode(result + "\n"))
            elif sha512Match:
                result = hashlib.sha512(sha512Match[1].encode()).hexdigest()
                print(data,result)
                sock.send(str.encode(result + "\n"))
