CREATE TABLE transactions (
    invoice_number VARCHAR(50) PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
	service_id BIGINT UNSIGNED NULL,
    total_amount DECIMAL(15,0) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL DEFAULT 'PAYMENT',
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (service_id) REFERENCES services(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;