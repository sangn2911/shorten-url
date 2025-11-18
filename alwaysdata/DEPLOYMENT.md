# Deploy on alwaysdata

## Prerequisites
- An account on [alwaysdata](https://admin.alwaysdata.com/)  
- Go binary for the application ShortLink from GitHub  

## Deployment

### MySQL
1. Go to [MySQL on alwaysdata](https://admin.alwaysdata.com/database/?type=mysql)  
    - ![MySQL](./images/mysql-alwaysdata.png)  
2. In the **DATABASES** tab, add a new database named `shortenurl_database` (this name will be used in the application configuration).  
3. In the **USERS** tab, modify the password of the existing user (for example, user `441172`).  
4. Go to [phpMyAdmin](https://phpmyadmin.alwaysdata.com/), and log in with user `441172`.  
5. Select `shortenurl_database`, go to the **SQL** tab:  
    - ![phpMyAdmin](./images/phpMyAdmin-alwaysdata.png)  
6. Copy and execute the script [create_table.sql](./create_table.sql).  
    - **Note:** Some collations may not be supported. You can create a new table using the UI and select a suitable collation.  
        - ![Collations](./images/collation-list.png)  
7. Verify the setup with the query:  
    ```sql
    SELECT * FROM url_encode;
    ```

---

### Application
*Note: For the free plan on alwaysdata, Docker is not supported. To deploy a Go application, we need to build a binary file and upload it to alwaysdata storage.*

#### 1. Build Binary
- Start at the root of the project.  
- Execute:  
    ```bash
    ./alwaysdata/build_binary.sh
    ```  
    This builds a binary file named `shortlink`.  
- Push the binary file to the GitHub branch `deploy_alwaysdata` (this will be downloaded into alwaysdata storage).  
- On GitHub, locate the binary file `shortlink` in branch `deploy_alwaysdata`.  
- Select the file and copy the link address (**binary URL**) from the **Raw** button.  
    - ![Raw Button](./images/raw-button.png)  

#### 2. Download Binary into alwaysdata Storage
*Note: The account name (here `shortenurl`) may vary for different users.*  
- Go to [SSH on alwaysdata](https://admin.alwaysdata.com/ssh/)  
    - ![SSH](./images/ssh-alwaysdata.png)  
- Modify the password of the existing user (e.g., `shortenurl`).  
- Access the terminal via [SSH websites](https://ssh-shortenurl.alwaysdata.net/) using the credentials above.  
- Execute:  
    ```bash
    wget <binary-url>
    ```  
    Replace `<binary-url>` with the link copied in the previous step. Use `ls` to verify the binary is downloaded.  
- Make the binary executable:  
    ```bash
    chmod +x ./shortlink
    ```  
    - ![Terminal](./images/ssh-terminal.png)  

#### 3. Start Go Application on alwaysdata
*Note: For the free plan, the site address must be `shortenurl.alwaysdata.net` (the account name).*  
- Go to [Sites on alwaysdata](https://admin.alwaysdata.com/site/) and click **Add a site**.  
    - ![Sites](./images/sites-alwaysdata.png)  
- Fill in the information:  
    - **Addresses:** `shortenurl.alwaysdata.net`  
    - **Configuration:**  
        - Type: User program  
        - Command: `~/shortlink` (binary in home directory)  
        - Working directory: leave empty  
        - Environment: fill using [alwaysdata.env.sample](./alwaysdata.env.sample)  
          - *Note: For the `PORT` variable, use the port specified below the Command input box.*  
    - ![Sites Config](./images/sites-config.png)  
- Submit the configuration.  
- Verify the application is running properly at:  
    ```
    https://shortenurl.alwaysdata.net/swagger/index.html
    ```  
    - ![Swagger](./images/swagger.png)  
