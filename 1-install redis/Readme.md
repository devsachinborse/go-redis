## To install Redis on Windows Subsystem for Linux 🐧 (WSL), follow these steps:

#### Step 1: Update your package lists
First, ensure your package lists are up to date:
```
sudo apt update
```

#### Step 2: Install Redis
Next, install Redis using the following command:

```
sudo apt install redis-server
```

#### Step 3: Verify the Installation
After the installation is complete, verify that Redis is installed by checking its version:

```
redis-server --version
```
#### Step 4: Start Redis Server
Start the Redis server using the following command:
```
sudo service redis-server start
```
#### Step 5: Test Redis
You can test if Redis is working by using the Redis CLI:

```
redis-cli
````
Once inside the Redis CLI, you can try setting and getting a key-value pair:

```
SET hello world
GET hello
```
You should see:

```
"world"
````
#### Step 6: Enable Redis to Start on Boot (Optional)
If you want Redis to start automatically when your WSL instance starts, you can enable the Redis service:


```
sudo systemctl enable redis-server
```

Redis should now be installed and running on your WSL environment!

