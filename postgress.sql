

CREATE TABLE IF NOT EXISTS categories (
  id bigint  NOT NULL,
  name varchar(191)  NOT NULL,
  slug varchar(191)  NOT NULL,
  status int NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
);



CREATE TABLE IF NOT EXISTS customers (
id bigint  NOT NULL,
  name varchar(191)  NOT NULL,
  mobile varchar(191)  NOT NULL,
  address varchar(191)  NOT NULL,
  email varchar(150)  NOT NULL,
  details text  NOT NULL,
  previous_balance text  NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
);


CREATE TABLE IF NOT EXISTS invoices (
id bigint  NOT NULL,
  customer_id bigint  NOT NULL,
  total varchar(191)  NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) ;


CREATE TABLE IF NOT EXISTS migrations (
id int  NOT NULL,
  migration varchar(191)  NOT NULL,
  batch int NOT NULL
) ;



CREATE TABLE IF NOT EXISTS products (
  id bigint  NOT NULL,
  name varchar(191)  NOT NULL,
  serial_number int NOT NULL,
  model varchar(191)  NOT NULL,
  category_id int NOT NULL,
  sales_price varchar(191)  NOT NULL,
  unit_id int NOT NULL,
  image varchar(191)  NOT NULL,
  tax_id varchar(191)  NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) ;


CREATE TABLE IF NOT EXISTS product_suppliers (
id bigint  NOT NULL,
  product_id int NOT NULL,
  supplier_id int NOT NULL,
  price int NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) ;



CREATE TABLE IF NOT EXISTS purchases (
  id int NOT NULL,
  supplier_id bigint  DEFAULT NULL,
  date date DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP 
);


CREATE TABLE IF NOT EXISTS sales (
id bigint  NOT NULL,
  invoice_id bigint  NOT NULL,
  product_id bigint  NOT NULL,
  qty int NOT NULL,
  price int NOT NULL,
  dis int NOT NULL,
  amount int NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) 


CREATE TABLE IF NOT EXISTS suppliers (
  id bigint  NOT NULL,
  name varchar(191)  NOT NULL,
  mobile varchar(191)  NOT NULL,
  address varchar(191)  NOT NULL,
  details text  NOT NULL,
  previous_balance text  NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) ;


CREATE TABLE IF NOT EXISTS taxes (
id bigint  NOT NULL,
  name varchar(191)  NOT NULL,
  slug varchar(191)  NOT NULL,
  status int NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) ;


CREATE TABLE IF NOT EXISTS units (
id bigint  NOT NULL,
  name varchar(191)  NOT NULL,
  slug varchar(191)  NOT NULL,
  status int NOT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) ;


CREATE TABLE IF NOT EXISTS users (
  id bigint  NOT NULL,
  f_name varchar(191)  NOT NULL,
  l_name varchar(191)  NOT NULL,
  email varchar(191)  NOT NULL,
  image varchar(191)  NOT NULL,
  email_verified_at timestamp NULL DEFAULT NULL,
  password varchar(191)  NOT NULL,
  remember_token varchar(100)  DEFAULT NULL,
  created_at timestamp NULL DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL
) ;




--
-- Indexes for table categories
--
ALTER TABLE categories
 ADD PRIMARY KEY (id);

--
-- Indexes for table customers
--
ALTER TABLE customers
 ADD PRIMARY KEY (id);

--
-- Indexes for table invoices
--
ALTER TABLE invoices
 ADD PRIMARY KEY (id)


--
-- Indexes for table products
--
ALTER TABLE products
 ADD PRIMARY KEY (id);

--
-- Indexes for table product_suppliers
--
ALTER TABLE product_suppliers
 ADD PRIMARY KEY (id);

--
-- Indexes for table purchases
--
ALTER TABLE purchases
 ADD PRIMARY KEY (id);

--
-- Indexes for table sales
--
ALTER TABLE sales
 ADD PRIMARY KEY (id);

--
-- Indexes for table suppliers
--
ALTER TABLE suppliers
 ADD PRIMARY KEY (id);

--
-- Indexes for table taxes
--
ALTER TABLE taxes
 ADD PRIMARY KEY (id);

--
-- Indexes for table units
--
ALTER TABLE units
 ADD PRIMARY KEY (id);

--
-- Indexes for table users
--
ALTER TABLE users
 ADD PRIMARY KEY (id);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table categories
--
ALTER TABLE categories
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;
--
-- AUTO_INCREMENT for table customers
--
ALTER TABLE customers
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table invoices
--
ALTER TABLE invoices
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=14;
--
-- AUTO_INCREMENT for table migrations
--
ALTER TABLE migrations
MODIFY id int  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=13;
--
-- AUTO_INCREMENT for table products
--
ALTER TABLE products
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=10;
--
-- AUTO_INCREMENT for table product_suppliers
--
ALTER TABLE product_suppliers
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=15;
--
-- AUTO_INCREMENT for table purchases
--
ALTER TABLE purchases
MODIFY id int NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table sales
--
ALTER TABLE sales
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=17;
--
-- AUTO_INCREMENT for table suppliers
--
ALTER TABLE suppliers
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table taxes
--
ALTER TABLE taxes
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT for table units
--
ALTER TABLE units
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table users
--
ALTER TABLE users
MODIFY id bigint  NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- Constraints for dumped tables
--

--
-- Constraints for table invoices
--
ALTER TABLE invoices
ADD CONSTRAINT invoices_customer_id_foreign FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE;

--
-- Constraints for table purchases
--
ALTER TABLE purchases
ADD CONSTRAINT fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES suppliers (id);

--
-- Constraints for table sales
--
ALTER TABLE sales
ADD CONSTRAINT sales_invoice_id_foreign FOREIGN KEY (invoice_id) REFERENCES invoices (id) ON DELETE CASCADE,
ADD CONSTRAINT sales_product_id_foreign FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE;

