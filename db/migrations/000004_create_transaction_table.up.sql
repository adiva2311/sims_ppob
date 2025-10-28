CREATE TABLE transactions (
    invoice_number SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    service_id INT NOT NULL,
    total_amount DECIMAL(15,0) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL DEFAULT 'PAYMENT',
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (service_id) REFERENCES services(id)
)Engine=InnoDB DEFAULT CHARSET=utf8mb4;