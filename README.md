# Order Processing

Thanks for the assignment, it was quite interesting. 
Few things before you proceed to set expectations
- Some of the code is commented which was created for some purpose but didnt seem useful later. But I kept it because it can be used to further build up the project. So ignore that for now. 
- Missed some of the components like Unit tests due to lack of time. I prioritized to focus on the key aspects first
- Everyone can make mistakes, it might be possible if i had understood the problem in a different spirit. If thats the case, please share the feedback

## 📦 Features
- 'X' no of orders are created simulataneously as soon as server is up.
- 'Y' no of orders are picked up from pending state every 'Z' seconds and put into processing state
- 'P' no of orders are picked up from processing state every 'Q' seconds and put into completed state
- X, Y, Z, P and Q are configurable as per your wishes before running the server
- You can hit the order_metrics api anytime to see the current state which includes total no of orders in the system, no of pending, processing and completed orders and average time taken to complete an order from the time of creation

## 🚀 Setup Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/vkeshav123/parspec_assignment.git
    ```
2. Create a local postgres databse for the project
3. Update the database configuration in the `oms-service-configuration.yml` file in root directory
4. Run the following script to create order table. You can run this script any no of times, if you need to clean the data
    ```bash
    DROP TABLE IF EXISTS "orders";
    DROP TYPE IF EXISTS order_status;

    CREATE TYPE order_status AS ENUM ('pending', 'processing', 'completed');
    CREATE TABLE "orders" (
        id SERIAL PRIMARY KEY,
        product_id INTEGER,
        quantity INTEGER,
        order_amount INTEGER NOT NULL,
        status order_status NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        completed_at TIMESTAMP
    );
    ```
5. Tweak order configuration in `oms-service-configuration.yml` file as per your preference.    For reference
    - `load_creation_batch_size` refers to the no of orders which will be created simulataneously once server is up
    - `processing_batch_size` refers to no of orders which will be picked from `pending` state and will be updated to `processing` state
    - `processing_interval` refers to the time interval in seconds which will do that above step repeatedly
    - `completion_batch_size` refers to no of orders which will be picked from `processing` state and will be updated to `completed` state
    - `completion_interval` refers to the time interval in seconds which will do that above step repeatedly
6. Open into the project directory and run from terminal 
    ```bash 
    go run main.go
    ```
7. Run the metrics api via following curl. Response of API is self explainatory. 
    ```bash
    curl --location 'http://localhost:8080/order_metrics'
    ``` 
8. Keep hitting the server randomly with the metrics API and observe how the average pending state time, processing time and completion time evolves. That will guide us to tweak the configuration intervals to optimize the orders as well. From here, we enter analytics.