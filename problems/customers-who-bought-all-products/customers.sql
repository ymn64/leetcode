-- Schema
CREATE TABLE IF NOT EXISTS Customer (
  customer_id INT,
  product_key INT
);

CREATE TABLE Product (
  product_key INT
);

TRUNCATE TABLE Customer;

INSERT INTO Customer (customer_id, product_key)
  VALUES ('1', '5'),
  ('2', '6'),
  ('3', '5'),
  ('3', '6'),
  ('1', '6');

TRUNCATE TABLE Product;

INSERT INTO Product (product_key)
  VALUES ('5'),
  ('6');

SELECT
  customer_id
FROM
  `Customer`
GROUP BY
  `customer_id`
HAVING
  COUNT(DISTINCT `product_key`) = ( -- table may contain duplicates rows
    SELECT
      COUNT(`product_key`)
    FROM
      `Product`);

