# CRX Package Format
CRX files are ZIP files with a special header and the .crx file extension.

## Package header
The header contains the author's public key and the extension's signature. The signature is generated from the ZIP file using SHA-1 with the author's private key. The header requires a little-endian byte ordering with 4-byte alignment. The following table describes the fields of the .crx header in order  

|Field|Type|Length|Value|Description|
|---|---|---|---|---|
|magic number|char[]|32 bits|`Cr24`|Chrome requires this constant at the beginning of every .crx package.|
|version|unsigned int|32 bits|`2`|The version of the `*.crx` file format used (currently 2).|
|public key length|unsigned int|32 bits|pubkey.length|The length of the RSA public key in bytes.|
|signature length|unsigned int|32 bits|sig.length|The length of the signature in bytes.|
|public key|byte[]|pubkey.length|pubkey.contents|The contents of the author's RSA public key, formatted as an X509 SubjectPublicKeyInfo block.|
|signature|byte[]|sig.length|sig.contents|The signature of the ZIP content using the author's private key. The signature is created using the RSA algorithm with the SHA-1 hash function.|

## Extension contents
The extension's ZIP file is appended to the `*.crx` package after the header. This should be the same ZIP file that the signature in the header was generated from.

## Example
The following is an example hex dump from the beginning of a `.crx` file.

```hex
43 72 32 34   # "Cr24" -- the magic number
02 00 00 00   # 2 -- the crx format version number
A2 00 00 00   # 162 -- length of public key in bytes
80 00 00 00   # 128 -- length of signature in bytes
...........   # the contents of the public key
...........   # the contents of the signature
...........   # the contents of the zip file
```
