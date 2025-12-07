-- Qreate Database Schema
-- PostgreSQL

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(10) NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin')),
    email_verified BOOLEAN DEFAULT false NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uk_users_email UNIQUE (email)
);

-- QR Codes table
CREATE TABLE IF NOT EXISTS qr_codes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name VARCHAR(255),
    short_code VARCHAR(20) NOT NULL,
    destination_url TEXT NOT NULL,
    image_url TEXT NOT NULL,
    is_active BOOLEAN DEFAULT true NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_qr_codes_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT uk_qr_codes_short_code UNIQUE (short_code)
);

CREATE INDEX IF NOT EXISTS idx_qr_codes_short_code ON qr_codes(short_code);
CREATE INDEX IF NOT EXISTS idx_qr_codes_user_id ON qr_codes(user_id);

-- QR Scans table
CREATE TABLE IF NOT EXISTS qr_scans (
    id BIGSERIAL PRIMARY KEY,
    qr_code_id INTEGER NOT NULL,
    scanned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip_address INET,
    user_agent TEXT,
    country_code VARCHAR(2),
    city_name VARCHAR(100),
    device_type VARCHAR(50),
    referer TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_qr_scans_qr_code_id FOREIGN KEY (qr_code_id) REFERENCES qr_codes(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_qr_scans_qr_code_id ON qr_scans(qr_code_id);
CREATE INDEX IF NOT EXISTS idx_qr_scans_scanned_at ON qr_scans(scanned_at);