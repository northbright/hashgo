# hashgo

hashgo is a program written in [go](https://golang.org) which computes hashes of given files.
It's based on [filehashes](https://github.com/northbright/filehashes).

## Usage
```
Usage:
hashgo [-md5 | -sha1 | -sha256 | -sha512] [File]...

  -h	this help
  -md5
    	compute MD5 checksum (default true)
  -sha1
    	compute SHA1 checksum (default true)
  -sha256
    	compute SHA256 checksum
  -sha512
    	compute SHA512 checksum
```

## Examples
* If no hash algorithm is set, it computes MD5 and SHA-1 of file by default

  ```
  ./hashgo ~/Iso/CentOS-8.1.1911-x86_64-dvd1.iso 
  total progress: 100%
  CentOS-8.1.1911-x86_64-dvd1.iso
  SHA-1: 49F028035C52A2FEB89D817E0451EEB4C597B77A
  MD5: 8D0573C5FB5444007936B652D8C6724D
  ```

* Specify SHA-256 hash algorithm

  ```
  ./hashgo -sha256 ~/Iso/CentOS-8.1.1911-x86_64-dvd1.iso
  total progress: 100%
  CentOS-8.1.1911-x86_64-dvd1.iso
  MD5: 8D0573C5FB5444007936B652D8C6724D
  SHA-1: 49F028035C52A2FEB89D817E0451EEB4C597B77A
  SHA-256: 3EE3F4EA1538E026FFF763E2B284A6F20B259D91D1AD5688F5783A67D279423B
  ```

* Compute hashes of multiple files

  ```
  ./hashgo ~/Iso/CentOS-8.1.1911-x86_64-dvd1.iso ~/Iso/CentOS-7-x86_64-Minimal-1810.iso
  total progress: 100%
  CentOS-7-x86_64-Minimal-1810.iso
  MD5: BD43D41E01C2A46B3CB23EB9139DCE4B
  SHA-1: 5833CB3189A61C02ABF50FB8C2DB16CFA669BC3C


  CentOS-8.1.1911-x86_64-dvd1.iso
  MD5: 8D0573C5FB5444007936B652D8C6724D
  SHA-1: 49F028035C52A2FEB89D817E0451EEB4C597B77A
  ```
