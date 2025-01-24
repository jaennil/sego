CREATE TABLE category(
    title VARCHAR(255) NOT NULL PRIMARY KEY
);

CREATE TABLE account (
    name VARCHAR(255) NOT NULL PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    balance DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP NOT NULL,
    currency_code CHAR(3) NOT NULL,
    CHECK ( length(type) > 0 )
);

CREATE TABLE "transaction" (
    transaction_id SERIAL PRIMARY KEY NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP NOT NULL,
    type VARCHAR(255) NOT NULL,
    account VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    FOREIGN KEY (account) REFERENCES account(name),
    FOREIGN KEY (category) REFERENCES category(title),
    CHECK ( length(type) > 0 )
);

INSERT INTO category(title) VALUES('Bills');
INSERT INTO category(title) VALUES('Telephone');
INSERT INTO category(title) VALUES('Electricity');
INSERT INTO category(title) VALUES('Gas');
INSERT INTO category(title) VALUES('Internet');
INSERT INTO category(title) VALUES('Rent');
INSERT INTO category(title) VALUES('Cable TV');
INSERT INTO category(title) VALUES('Water');
INSERT INTO category(title) VALUES('Food');
INSERT INTO category(title) VALUES('Groceries');
INSERT INTO category(title) VALUES('Dining out');
INSERT INTO category(title) VALUES('Leisure');
INSERT INTO category(title) VALUES('Movies');
INSERT INTO category(title) VALUES('Video Rental');
INSERT INTO category(title) VALUES('Magazines');
INSERT INTO category(title) VALUES('Automobile');
INSERT INTO category(title) VALUES('Maintenance');
INSERT INTO category(title) VALUES('Parking');
INSERT INTO category(title) VALUES('Registration');
INSERT INTO category(title) VALUES('Education');
INSERT INTO category(title) VALUES('Books');
INSERT INTO category(title) VALUES('Tuition');
INSERT INTO category(title) VALUES('Homeneeds');
INSERT INTO category(title) VALUES('Clothing');
INSERT INTO category(title) VALUES('Furnishing');
INSERT INTO category(title) VALUES('Others');
INSERT INTO category(title) VALUES('Healthcare');
INSERT INTO category(title) VALUES('Dental');
INSERT INTO category(title) VALUES('Eyecare');
INSERT INTO category(title) VALUES('Physician');
INSERT INTO category(title) VALUES('Prescriptions');
INSERT INTO category(title) VALUES('Insurance');
INSERT INTO category(title) VALUES('Auto');
INSERT INTO category(title) VALUES('Life');
INSERT INTO category(title) VALUES('Home');
INSERT INTO category(title) VALUES('Health');
INSERT INTO category(title) VALUES('Vacation');
INSERT INTO category(title) VALUES('Travel');
INSERT INTO category(title) VALUES('Lodging');
INSERT INTO category(title) VALUES('Sightseeing');
INSERT INTO category(title) VALUES('Taxes');
INSERT INTO category(title) VALUES('Income Tax');
INSERT INTO category(title) VALUES('House Tax');
INSERT INTO category(title) VALUES('Water Tax');
INSERT INTO category(title) VALUES('Miscellaneous');
INSERT INTO category(title) VALUES('Gifts');
INSERT INTO category(title) VALUES('Income');
INSERT INTO category(title) VALUES('Salary');
INSERT INTO category(title) VALUES('Reimbursement/Refunds');
INSERT INTO category(title) VALUES('Investment Income');
INSERT INTO category(title) VALUES('Other Income');
INSERT INTO category(title) VALUES('Other Expenses');
INSERT INTO category(title) VALUES('Transfer');