# shorten-url


## Scalabilitiy
#### Shorten id collision  
*How service encodes an url?*  
- A shorten id generated according to the number presenting the order of an url encoded
- The service encodes shorten id consisting of a-z in uppercase/lowercase, and number 0-9, in total of 62 character
- The service converts the order number (start from 0) of the current url encoded to a base62 digit, then map with the character  
  - For examples, with starting length of shorten id is 6:  
  - aaaaaa, baaaaa, caaaaa is the shorten id generating when encoding the 1st (0), 2nd (1), 3rd (2) url  
  - GeAi9K is the shorten id generating when encoding the 33884145297th (33884145296) url  

*How service handles collision?*  
- When service reaches the max order number of shorten id with length n. Services will add 1 to length n and start generating short ids with length n + 1. Therefore, service can avoid collision as shorten id are generated uniquely as the order number of urls encoded

#### Large amount of unused shorten urls (cold data)
*Describe about the issue:*  
- The service currently does not address the issue of cold data, which are unused urls. This issue causes waste of resources  

*Describe the solution:*  
- For each url encoded, service will assign a TTL (Time to live) showing how many days the shorten url will expires. When the time expires, I can implement an scheduled EVENT in MySQL to periodically remove the encoded urls
- The shorten ids for these removed urls will be stored in another SQL table, so I can reuse them  

#### Large amount of requests to encode/decode a range of urls
*Describe about the issue:*  
- A range of requests encode/decode some specific urls. This issue causes high pressure on MySQL  

*Describe the solution:*  
- Using cache with TTL (such as Redis) for these specific urls when encoding/decoding occur for better performance and reducing pressure on MySQL
  - For examples:
    - After encoding an url, this url can be stored with its shorten id in cache
    - After decoding an shorten url, the shorten id of this url will be stored with its decoded one in cache
    - The next time user encodes/decodes these urls, service can get them from cache

## Potential attack vectors & mitigations

## Supported
1. SQL Injection  
Description: Database is controlled by SQL injected in incoming requests  
Mitigations:  
    - Service uses SQL queries with placeholder for values, MySQL driver automatically transforms these queries to prepared statements  
2. Insecure Deserialization  
Description: Service failures when receiving invalid JSON request  
Mitigations:  
    - Service has validated the url in request (check empty, invalid format, wrong public domain)  
    - For a more complex request, the deserialization requires to apply package "github.com/go-playground/validator/v10" for better clean and maintainable code  
3. Sensitive Data Exposure  
Description: This attack is the use of leaking secrets, DB credentials  
Mitigations:  
    - Service uses environment variables
    - Service should only use user and password, which have limited permissions suitable with their purposes, to use resources (for example: MySQL user provided can only SELECT, INSERT specific tables) in production
4. Misconfigured CORS  
Description: CORS is too open, any website can call service's APIs  
Mitigations:  
    - Service support env variables to config CORS "ALLOW_ORIGIN", "ALLOW_METHODS", "ALLOW_METHODS"  
5. Transport Layer Attacks  
Description: Man in the middle  
Mitigation:  
    - Always use HTTPS, alwaysdata provides HTTPS for service  

### Unsupported
1. Denial of Service (DoS / DDoS)  
Description: Attack tries to overload servicer resources  
Mitigations:  
    - Applies technologies providing Rate limiting (such as Nginx) to limit the number of requests according to the limitation of service  
    - Implements circuit breaker in case resouces failure (for example: MySQL failures, pending queries) to reduce the pressure on these resources
