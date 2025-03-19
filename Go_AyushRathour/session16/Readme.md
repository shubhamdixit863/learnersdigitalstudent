### Concurrent Order Processing System



## Overview
You need to design a concurrent order processing system for an e-commerce platform that efficiently
handles multiple simultaneous orders while ensuring data consistency using goroutines, channels, and mutexes.


## Requirements

1️⃣ Concurrent Order Placement (Goroutines & Channels)

• Customers place orders simultaneously.

• Orders should be processed as they arrive without batching.

• Each order should be assigned a unique orderID.

• Use a channel (orderQueue) to pass orders for processing.

2️⃣ Inventory Management (Mutex)

• Orders reduce stock availability.

• Use a mutex (sync.Mutex) to prevent race conditions when updating inventory.

• If an order is placed but insufficient stock is available, it should be rejected.


3️⃣ Order Status Updates

• Each order can have one of these statuses:

• “Processing”

• “Shipped”

• “Completed”

• Status updates should be handled via channels and logged.



4️⃣ Graceful Shutdown Handling

• Ensure ongoing orders complete before shutting down.

• No new orders should be accepted after shutdown.


Constraints

• The system should handle at least 100 concurrent orders without data inconsistency.

• each order is processed independently.

• Avoid deadlocks and race conditions.

• The order processing time should vary (some take longer than others).