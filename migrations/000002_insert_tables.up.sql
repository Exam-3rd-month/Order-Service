BEGIN;

-- Insert mock data into orders table
INSERT INTO orders (order_id, user_id, kitchen_id, items, total_amount, status, delivery_address, delivery_time, is_done, created_at, updated_at)
VALUES 
('e1d2c3b4-5678-9abc-def0-1234567890ab', 'd1e2f3a4-5678-9abc-def0-1234567890ab', 'a1b2c3d4-5678-9abc-def0-1234567890ab', '[{"dish_id": "f3e2c789-334e-4cbf-8976-50d7b4128a16", "quantity": 2}]', 25.98, 'Preparing', '123 Main St, Cityville', CURRENT_TIMESTAMP + interval '1 hour', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('2f3e4d5c-6789-abcd-ef01-234567890abc', 'd1e2f3a4-5678-9abc-def0-1234567890ab', 'a1b2c3d4-5678-9abc-def0-1234567890ab', '[{"dish_id": "cd50db24-3e29-45df-91a4-50339d2dcf0e", "quantity": 1}]', 8.99, 'Delivered', '456 Elm St, Townsville', CURRENT_TIMESTAMP - interval '2 hours', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('c4d5e6f7-789a-bcde-f012-34567890abcd', 'd4e5f6b7-6789-abcd-ef01-234567890abc', 'b2c3d4e5-6789-abcd-ef01-234567890abc', '[{"dish_id": "f462dfed-15a0-4e08-985e-7e3b1a4ad2a5", "quantity": 1}]', 7.99, 'Preparing', '789 Oak St, Villageville', CURRENT_TIMESTAMP + interval '1.5 hours', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert mock data into reviews table
INSERT INTO reviews (review_id, order_id, user_id, kitchen_id, rating, comment, created_at, updated_at)
VALUES 
('f1d2e3c4-5678-9abc-def0-1234567890ab', 'e1d2c3b4-5678-9abc-def0-1234567890ab', 'd1e2f3a4-5678-9abc-def0-1234567890ab', 'a1b2c3d4-5678-9abc-def0-1234567890ab', 4.5, 'Delicious!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('3f4e5d6c-6789-abcd-ef01-234567890abc', '2f3e4d5c-6789-abcd-ef01-234567890abc', 'd1e2f3a4-5678-9abc-def0-1234567890ab', 'a1b2c3d4-5678-9abc-def0-1234567890ab', 5.0, 'Perfect pizza!', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('5d6e7f8c-789a-bcde-f012-34567890abcd', 'c4d5e6f7-789a-bcde-f012-34567890abcd', 'd4e5f6b7-6789-abcd-ef01-234567890abc', 'b2c3d4e5-6789-abcd-ef01-234567890abc', 3.5, 'Good salad but could use more dressing.', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert mock data into payments table
INSERT INTO payments (payment_id, order_id, amount, status, payment_method, transaction_id, created_at)
VALUES 
('a1b2c3d4-5678-9abc-def0-1234567890ab', 'e1d2c3b4-5678-9abc-def0-1234567890ab', 25.98, 'Completed', 'Credit Card', 'txn_1234567890', CURRENT_TIMESTAMP),
('b2c3d4e5-6789-abcd-ef01-234567890abc', '2f3e4d5c-6789-abcd-ef01-234567890abc', 8.99, 'Completed', 'PayPal', 'txn_0987654321', CURRENT_TIMESTAMP),
('c3d4e5f6-789a-bcde-f012-34567890abcd', 'c4d5e6f7-789a-bcde-f012-34567890abcd', 7.99, 'Pending', 'Credit Card', 'txn_2345678901', CURRENT_TIMESTAMP);

COMMIT;
