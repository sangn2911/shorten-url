# API Testing with Postman

## Prerequisites
- Postman version 10.8.0

## Running Tests
1. Ensure the `url_encode` table is empty, as existing data may conflict with the test cases.  
2. Import the [Postman Collection](./ShortenUrl.postman_collection.json).  
3. Open the collection named **ShortenUrl**.  
4. Click **Run** to execute the collection.  
5. Make sure the execution order is **Encode** first, then **Decode**.  
6. Select the [Testcases file](./testcases.json).  
7. Click **Run ShortenUrl** to execute the test cases.
