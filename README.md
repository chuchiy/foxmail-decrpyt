## Foxmail Saved Password Decrypt/Foxmail保存密码查看器

Open your `Account.stg` file with text editor in foxmail mail account folder. 
Find the hex encypt password line in the file. Just like
```
XXXXPassword=DEADBEEF......
```

Decrypt hex encypt password with foxmail-decrypt

```
> .\foxmail-decrpyt.exe
Usage of foxmail-decrpyt.exe:
  -p string
        hex encrypt password
  -v6
        password from foxmail 6

> .\foxmail-decrpyt.exe -v6 -p 61D05ABC4CFB32CB410C5AFF
smtppass123
```

build binary

```
> go build .\cmd\foxmail-decrpyt\
```