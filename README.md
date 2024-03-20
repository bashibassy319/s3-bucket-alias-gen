# S3 bucket alias generator for s5cmd 

Configure bash alias for s5cmd to store Ceph S3 access credentials.   

s5cmd is not so handy to manage multiple S3 compatible object stores like Ceph or MinIO.  
s5cmd does not have a config file so we need to pass params each time.  

This tool will help generating an alias for each bucket.  

## How it works
This program will generate a access configuration in `~/.aws/credentials`. 

```
[s3-minio-test-bucket]
aws_access_key_id = AWESOMERANDOMEACCESSKEY
aws_secret_access_key = AWESOMERANDOMSECRETKEY
```

And then, the program will generate an alias in `~/.bashrc`.  
```
alias s3-minio-test-bucket='AWS_PROFILE=s3-minio-test-bucket S3_ENDPOINT_URL=https://myminio.com'
```

Reload the bashrc file if nessesary: 
```
source ~/.bashrc
```

Finally, you can use an alias with s5cmd (or any other alternatives): 
```
s3-minio-test-bucket s5cmd ls 
```

## How to build
```
go build
./s3-bucket-alias-gen
```

## Example 
```
$./s3-bucket-alias-gen
✔ ACCESS KEY: test
✔ SECRET KEY: aaaa
✔ alias name: test-alias
✔ endpoint: https://test.com
```


## Options
```
$./auto-ceph-s3-conf --help
Usage of ./auto-ceph-s3-conf:
  -aws-credential-path string
        Path to AWS credentials file, default path is /home/dev/.aws/credentials (default "/home/dev/.aws/credentials")
  -bashrc-path string
        Path to .bashrc file default path is /home/dev/.bashrc (default "/home/dev/.bashrc")
  -dry-run
        Run in dry-run mode
  -help
        Display help information
```



