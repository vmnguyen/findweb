# findweb
Go tool to find http/http services from list of IPs or domains

# Usage
Usage of findweb:
  -c int
        Concurrentcy (default 50)
  -f string
        Filepath (default "hosts.txt")
  -m string
        Input type: Domain or IP? (default "domain")

# Example
```findweb -m domain -c 10 -f ./hosts.txt```
