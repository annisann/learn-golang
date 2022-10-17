# Assignment 02: Build REST API

## Database
Tables: `items`, `orders`
1. items
    - item_id
    - item_code
    - description
    - quantity
    - order_id

2. orders
    - order_id
    - customer_name
    - ordered_at

## Endpoints
1. Create Order
- Path: `/order`
- Method: `POST`
2. Get Orders
- Path: `/orders`
- Method: `GET`
3. Update Order
- Path: `/orders/:orderID`
- Method: `PUT`
4. Delete Order
- Path: `/orders/:orderID`
- Method: `DELETE`
