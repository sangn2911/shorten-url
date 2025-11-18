# shorten-url

## Scalability
#### Shorten ID Collision  
*How does the service encode a URL?*  
- A shorten ID is generated based on the order number of the URL being encoded.  
- The service uses a character set of a-z (uppercase and lowercase) and numbers 0-9, giving a total of 62 possible characters.  
- The order number (starting from 0) of the URL is converted into a base62 digit and then mapped to the corresponding character.  
  - For example, with an initial shorten ID length of 6:  
    - "aaaaaa", "baaaaa", "caaaaa" are the shorten IDs generated for the 1st (0), 2nd (1), and 3rd (2) URLs.  
    - "GeAi9K" is the shorten ID generated for the 33,884,145,297th (33,884,145,296) URL.  

*How does the service handle collisions?*  
- When the service reaches the maximum order number for shorten IDs of length n, it increases the ID length to n + 1 and starts generating IDs with the new length.  
- This approach ensures that all shorten IDs are unique and avoids collisions.  

---

#### Large Amount of Unused Shorten URLs (Cold Data)  
*Issue:*  
- The service currently does not address cold data (unused URLs), which can result in wasted storage resources.  

*Solution:*  
- Each URL is assigned a TTL (Time To Live) indicating how long the shorten URL is valid.  
- When a URL expires, a scheduled MySQL EVENT can periodically remove it from the database.  
- The IDs of removed URLs are stored in a separate table for reuse.  

---

#### Large Number of Requests for Encoding/Decoding Specific URLs  
*Issue:*  
- High frequency requests to encode/decode the same set of URLs can overload the database.  

*Solution:*  
- Use a cache with TTL (e.g., Redis) for frequently accessed URLs to improve performance and reduce database load.  
  - After encoding a URL, store the URL and its shorten ID in the cache.  
  - After decoding a shorten ID, store the ID and its corresponding URL in the cache.  
  - Subsequent encode/decode requests for these URLs can be served from the cache.  

---

## Potential Attack Vectors & Mitigations

### Supported
*SQL Injection*  
- Description: Database can be manipulated via SQL injected in incoming requests.  
- Mitigations:  
    - Service uses SQL queries with placeholders for values. The MySQL driver automatically converts these queries into prepared statements.  

*Insecure Deserialization*  
- Description: Service failures can occur when receiving invalid JSON requests.  
- Mitigations:  
    - Service validates URLs in requests (checking for empty values, invalid formats, or incorrect domains).  
    - For complex requests, deserialization uses the package "github.com/go-playground/validator/v10" for cleaner and more maintainable code.  

*Sensitive Data Exposure*  
- Description: Leakage of secrets or database credentials.  
- Mitigations:  
    - Service uses environment variables for configuration.  
    - Database users have restricted permissions suitable for their purpose (e.g., MySQL users only allowed SELECT/INSERT on specific tables).  

*Misconfigured CORS*  
- Description: Overly permissive CORS allows any website to call the service's APIs.  
- Mitigations:  
    - Service supports environment variables to configure CORS: `ALLOW_ORIGIN`, `ALLOW_METHODS`, `ALLOW_HEADERS`.  

*Transport Layer Attacks*  
- Description: Man-in-the-middle attacks.  
- Mitigations:  
    - Always use HTTPS. alwaysdata provides HTTPS for the service.  

---

### Unsupported
*Denial of Service (DoS / DDoS)*  
- Description: Attacks attempt to overload service resources.  
- Mitigations:  
    - Apply rate limiting (e.g., Nginx) to control request rates according to service capacity.  
    - Implement circuit breakers to reduce pressure on resources during failures (e.g., MySQL timeouts or pending queries).  
