DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    name varchar(50) NOT NULL,
    date_of_birth date NOT NULL,
    city varchar(100) NOT NULL,
    zipcode VARCHAR(10) NOT NULL,
    status SMALLINT NOT NULL
);
CREATE TABLE accounts (
    account_id SERIAL PRIMARY KEY,
    customer_id integer NOT NULL,
    account_type varchar(50),
    opening_date date DEFAULT CURRENT_TIMESTAMP,
    amount numeric(10, 2) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    transaction_type varchar(10) NOT NULL,
    transaction_date date NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id)
);
INSERT INTO customers(
        customer_id,
        name,
        date_of_birth,
        city,
        zipcode,
        status
    )
VALUES (
        2000,
        'Steve',
        '1978-12-15',
        'Delhi',
        '110075',
        1
    ),
    (
        2001,
        'Arlan',
        '1988-05-21',
        'Newburgh NY',
        '12550',
        1
    ),
    (
        2002,
        'Hardley',
        '1988-04-30',
        'Englewood NJ',
        '07631',
        1
    ),
    (
        2003,
        'Ben',
        '1988-01-04',
        'Manchester NH',
        '03102',
        0
    ),
    (
        2004,
        'Nina',
        '1988-04-14',
        'Clarkston MI',
        '48348',
        1
    ),
    (
        2005,
        'Osman',
        '1988-11-08',
        'Hyattsville MD',
        '20782',
        0
    );
INSERT INTO accounts(
        account_id,
        customer_id,
        opening_date,
        account_type,
        amount,
        status
    )
VALUES (
        95470,
        2000,
        '2020-08-22 10:20:06',
        'saving',
        6823.23,
        1
    ),
    (
        95471,
        2002,
        '2020-08-09 10:27:22',
        'checking',
        3342.96,
        1
    ),
    --   (95472,2001,'2020-08-09 10:35:22', 'saving', 7000, 1),
    (
        95473,
        2003,
        '2020-08-09 10:38:22',
        'saving',
        5861.86,
        1
    );